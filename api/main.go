package api

import (
	//"fmt"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var DatabaseFile string

func TestApi(dataBase DataBase) {
	var apiConfig = ApiConfig{
		Ip:            "localhost:8080",
		LogFile:       "/tmp/seekr.log",
		DataBaseFile:  "test-data",
		DataBase:      dataBase,
		SetCORSHeader: true,
		LoadDBFunc:    DefaultLoadDB,
		SaveDBFunc:    DefaultSaveDB,
		Testing:       true,
	}
	err := apiConfig.SaveDB()
	if err != nil {
		log.Fatalf("Error saving to databse: %s", err)
	}

	// Start the background Seekrd service
	go Seekrd(DefaultSeekrdServices, 30) // run every 30 minutes

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
	config.GinRouter.GET("/", Handler(GetDataBase, config))                                      // return entire database
	config.GinRouter.GET("/deep/github/:username", Handler(GithubInfoDeepRequest, config))       // deep investigation of github account // FIXME
	config.GinRouter.GET("/search/google/:query", Handler(GoogleRequest, config))                // get results from google
	config.GinRouter.GET("/search/whois/:query", Handler(WhoisRequest, config))                  // get whois of domain
	config.GinRouter.GET("/people/:id", Handler(GetPersonByIDRequest, config))                   // return person obj
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

	runningFile, err := os.Create("/tmp/running")
	if err != nil {
		log.Println(err) // Fix me (breaks tests)
	}
	if err != nil {
		log.Fatalf("Error saving to databse: %s", err)
	}
	defer os.Remove("/tmp/running")
	defer runningFile.Close()

	config.DataBase, err = config.DataBase.Parse(config)
	config.GinRouter.Run(config.Ip)
}

func GithubInfoDeepRequest(config ApiConfig, c *gin.Context) {
	if c.Param("username") != "" {
		githubInfo, err := GithubInfoDeep(c.Param("username"), true, config)
		if err != nil {
			c.IndentedJSON(http.StatusForbidden, map[string]string{"fatal": fmt.Sprintf("%s", err)})
		} else {
			c.IndentedJSON(http.StatusOK, githubInfo)
		}
	}
}

func GetPersonByID(config ApiConfig, id string) (bool, Person) {
	if _, ok := config.DataBase[id]; ok {
		return true, config.DataBase[id]
	}
	return false, Person{}
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
	fmt.Println(person.Markdown())
	err = config.SaveDB()
	if err != nil {
		log.Fatalf("Error saving to databse: %s", err)
	}
}
