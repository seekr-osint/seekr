package api

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/seekr-osint/seekr/api/github"
	"github.com/seekr-osint/seekr/api/server"
	"github.com/seekr-osint/seekr/api/webserver"
)

var DatabaseFile string

func TestApi(dataBase DataBase) {
	apiConfig, err := ApiConfig{
		Server: server.Server{
			Ip:   "localhost",
			Port: uint16(8080),
			ApiServer: server.ApiServer{
				Disable: false,
			},
			WebServer: webserver.Webserver{
				Disable: true,
			},
		},
		LogFile:       "/tmp/seekr.log",
		DataBaseFile:  "test-data",
		DataBase:      dataBase,
		SetCORSHeader: true,
		LoadDBFunc:    DefaultLoadDB,
		SaveDBFunc:    DefaultSaveDB,
		Testing:       true,
	}.Parse()
	if err != nil {
		log.Fatalf("Error parsing test config: %s", err)
	}

	err = apiConfig.SaveDB()
	if err != nil {
		log.Fatalf("Error saving to databse: %s", err)
	}

	// Start the API server
	go ServeApi(apiConfig)
}

func CheckPersonExists(config ApiConfig, id string) bool {
	_, ok := config.DataBase[id]
	return ok
}

func ServeApi(config ApiConfig) {
	config, err := config.LoadDB()
	if err != nil {
		log.Fatalf("Error loading database: %s", err)
	}
	SetupLogger(config)
	config.GinRouter = gin.Default()
	if !config.Server.WebServer.Disable {
		fmt.Printf("Running WebServer on: %s", fmt.Sprintf("%s:%d", config.Server.Ip, config.Server.Port))
		config.SetupWebServer()
	}
	config.ServeTempMail()
	config.GinRouter.GET("/", Handler(GetDataBase, config))                                      // return entire database
	config.GinRouter.GET("/deep/github/:username", Handler(GithubInfoDeepRequest, config))       // deep investigation of github account // FIXME
	config.GinRouter.GET("/search/google/:query", Handler(GoogleRequest, config))                // get results from google
	config.GinRouter.GET("/search/whois/:query", Handler(WhoisRequest, config))                  // get whois of domain
	config.GinRouter.GET("/people/:id", Handler(GetPersonByIDRequest, config))                   // return person obj
	config.GinRouter.GET("/people/:id/markdown", Handler(MarkdownPersonRequest, config))         // return person obj
	config.GinRouter.DELETE("/people/:id", Handler(DeletePerson, config))                        // delete person
	config.GinRouter.GET("/people/:id/delete", Handler(DeletePerson, config))                    // delete person
	config.GinRouter.DELETE("/people/:id/accounts/:account", Handler(DeleteAccount, config))     // delete account
	config.GinRouter.GET("/people/:id/accounts/:account/delete", Handler(DeleteAccount, config)) // delete account
	config.GinRouter.POST("/person", Handler(PostPerson, config))                                // post person
	config.GinRouter.GET("/getAccounts/:username", Handler(GetAccountsRequest, config))          // get accounts
	config, err = config.Parse()
	if err != nil {
		log.Println(err) // Fix me (breaks tests)
	}
	config.SaveDB()

	config.DataBase, err = config.DataBase.Parse(config)
	if err != nil {
		log.Printf("error parsing databse:%s\n", err)
	}
	visited := make(map[reflect.Type]bool)
	fmt.Printf("%s", PrintTypeTreeRec(reflect.TypeOf(ApiConfig{}), visited, 0, 0))
	config.GinRouter.Run(fmt.Sprintf("%s:%d", config.Server.Ip, config.Server.Port))
}

func GithubInfoDeepRequest(config ApiConfig, c *gin.Context) {

	if c.Param("username") != "" {
		deep := github.DeepInvestigation{
			Username:  c.Param("username"),
			Tokens:    []string{},
			ScanForks: false,
		}
		apiEmails := EmailsType{}.Parse()
		emails, rateLimitRate, err := deep.GetEmails()
		log.Printf("RateLimitRate: %d\n", rateLimitRate)
		if err != nil {
			apiErr := err.(APIError)
			c.IndentedJSON(apiErr.Status, gin.H{"message": apiErr.Message})
			return
		} else {
			for _, emailObj := range emails {
				apiEmails[emailObj.Email] = Email{
					Mail:     emailObj.Email,
					Src:      emailObj.CommitUrl,
					Value:    1,
					Services: EmailServices{},
				}
				apiEmails[emailObj.Email].Services["GitHub"] = EmailService{
					Name:     "GitHub",
					Link:     fmt.Sprintf("https://github.com/%s", c.Param("username")),
					Username: c.Param("username"),
					Icon:     "./images/mail/github.svg",
				}
			}

			// concept impl to provide tate limitation info
			//c.IndentedJSON(http.StatusOK, map[string]interface{}{"rate_limit_rate": fmt.Sprintf("%d", rateLimitRate), "emails": emails})
			c.IndentedJSON(http.StatusOK, apiEmails.Parse())
		}
	}
}

func GetPersonByID(config ApiConfig, id string) (bool, Person) {
	if _, ok := config.DataBase[id]; ok {
		return true, config.DataBase[id]
	}
	return false, Person{}
}

func MarkdownPersonRequest(config ApiConfig, c *gin.Context) {
	person, err := config.GetPerson(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, nil)
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"markdown": person.Markdown()})
	}
}

func GetPersonByIDRequest(config ApiConfig, c *gin.Context) {
	person, err := config.GetPerson(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, nil)
	} else {
		c.IndentedJSON(http.StatusOK, person)
	}
}

func DeletePerson(config ApiConfig, c *gin.Context) {
	if CheckPersonExists(config, c.Param("id")) {
		delete(config.DataBase, c.Param("id"))
	}
	err := config.SaveDB()
	if err != nil {
		log.Fatalf("Error saving to databse: %s", err)
	}

}

func DeleteAccount(config ApiConfig, c *gin.Context) {
	if CheckPersonExists(config, c.Param("id")) {
		delete(config.DataBase[c.Param("id")].Accounts, c.Param("account")) // TODO check if stuff nonesense nobody needs
	}
	err := config.SaveDB()
	if err != nil {
		log.Fatalf("Error saving to database: %s", err)
	}
}

func Handler(function func(ApiConfig, *gin.Context), config ApiConfig) gin.HandlerFunc {
	handlerFunc := func(c *gin.Context) {
		if config.SetCORSHeader {
			c.Header("Access-Control-Allow-Origin", "*")
		}
		config, err := config.LoadDB()
		if err != nil {
			log.Fatalf("Error database: %s", err)
		}

		function(config, c)
	}
	return gin.HandlerFunc(handlerFunc)
}

func GetDataBase(config ApiConfig, c *gin.Context) {
	config, err := config.LoadDB()
	if err != nil {
		log.Fatalf("Error database: %s", err)
	}
	c.IndentedJSON(http.StatusOK, config.DataBase)
}

func GoogleRequest(config ApiConfig, c *gin.Context) {
	c.IndentedJSON(http.StatusOK, SearchString(c.Param("query")))
}

func WhoisRequest(config ApiConfig, c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Whois(c.Param("query"), config))
}

func GetAccounts(config ApiConfig, username string) Accounts {
	return ServicesHandler(DefaultServices, username, config)
}

func GetAccountsRequest(config ApiConfig, c *gin.Context) {
	c.IndentedJSON(http.StatusOK, GetAccounts(config, strings.ToLower(c.Param("username"))))
}

// THIS HAS NO C.PARAM("id")
// THIS IS THE WORST PEICE OF BAD CODE I EVER WROTE
func PostPerson(config ApiConfig, c *gin.Context) { // c.BindJSON is a person not people (POST "localhost:8080/person")
	var person Person

	// exit if the json is invalid
	if err := c.BindJSON(&person); err != nil {
		c.IndentedJSON(http.StatusAccepted, gin.H{"message": "invalid person"})
		return
	}
	err := person.Validate()
	if err != nil {
		apiErr := err.(APIError)
		c.IndentedJSON(apiErr.Status, gin.H{"message": apiErr.Message})
		return
	}

	person, _ = person.Parse(config)
	person, _ = person.CheckMail(config)
	// no error handeling doue to no error impl
	//if err := c.BindJSON(&person); err != nil {
	//	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "invalid person"})
	//	return
	//}
	// DON'T BE LIKE ME AND USE NEWPERSON.ID !!!
	exsits, _ := GetPersonByID(config, person.ID) // check rather the person Exsts
	if !exsits {
		// Add the new person to the database.
		if config.DataBase != nil {
			config.DataBase[person.ID] = person
		} else {
			config.DataBase = DataBase{}
			config.DataBase[person.ID] = person
		}
		c.IndentedJSON(http.StatusCreated, person)
	} else {
		log.Printf("overwritten person %s", person.ID)
		if config.DataBase != nil {
			config.DataBase[person.ID] = person
		} else {
			config.DataBase = DataBase{}
			config.DataBase[person.ID] = person
		}
		c.IndentedJSON(http.StatusAccepted, gin.H{"message": "overwritten person"})
	}
	err = config.SaveDB()
	if err != nil {
		log.Fatalf("Error saving to databse: %s", err)
	}
	//fmt.Println(person.Markdown())
}
