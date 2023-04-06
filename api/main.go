package api

import (
	//"fmt"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

var DatabaseFile string

func TestApi(dataBase DataBase) {
	var apiConfig = ApiConfig{
		Ip:            "localhost:8080",
		LogFile:       "/tmp/seekr.log",
		DataBaseFile:  "test-data.json",
		DataBase:      dataBase,
		SetCORSHeader: true,
		SaveJsonFunc:  DefaultSaveJson,
		Testing:       true,
	}
	DefaultSaveJson(apiConfig)

	// Start the background Seekrd service
	go Seekrd(DefaultSeekrdServices, 30) // run every 30 minutes

	// Start the API server
	go ServeApi(apiConfig)
}

func DefaultSaveJson(config ApiConfig) {
	log.Println("Saving json to file")
	jsonBytes, err := json.MarshalIndent(config.DataBase, "", "\t")
	CheckAndLog(err, "error saving the database to file", config)
	ioutil.WriteFile(config.DataBaseFile, jsonBytes, 0644)
}

func CheckPersonExists(config ApiConfig, id string) bool {
	_, ok := config.DataBase[id]
	return ok
}

func ServeApi(config ApiConfig) {
	config = LoadJson(config)
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
	runningFile, err := os.Create("/tmp/running")
	if err != nil {
		log.Println(err) // Fix me (breaks tests)
	}
	defer os.Remove("/tmp/running")
	defer runningFile.Close()
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
	config.SaveJsonFunc(config)
}

func DeleteAccount(config ApiConfig, c *gin.Context) {
	if CheckPersonExists(config, c.Param("id")) {
		delete(config.DataBase[c.Param("id")].Accounts, c.Param("account")) // TODO check if stuff nonesense nobody needs
	}
	config.SaveJsonFunc(config)
}

func Handler(function func(ApiConfig, *gin.Context), config ApiConfig) gin.HandlerFunc {
	handlerFunc := func(c *gin.Context) {
		if config.SetCORSHeader {
			c.Header("Access-Control-Allow-Origin", "*")
		}
		config = LoadJson(config)
		function(config, c)
	}
	return gin.HandlerFunc(handlerFunc)
}

func LoadJson(config ApiConfig) ApiConfig {
	if _, err := os.Stat(config.DataBaseFile); errors.Is(err, os.ErrNotExist) {
		log.Printf("creating %s DataBaseFile", config.DataBaseFile)
		err := os.WriteFile(config.DataBaseFile, []byte("{}"), 0755)
		CheckAndLog(err, fmt.Sprintf("error creating DataBaseFile: %s", config.DataBaseFile), config)
	}
	file, err := os.Open(config.DataBaseFile)
	if err != nil {
		log.Println(err)
	}
	CheckAndLog(err, "error opening database file", config)
	defer file.Close()

	err = json.NewDecoder(file).Decode(&config.DataBase)
	CheckAndLog(err, "error decoding", config)

	log.Println("loading json database from file")
	return config
}

func GetDataBase(config ApiConfig, c *gin.Context) {
	config = LoadJson(config)
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
	config.SaveJsonFunc(config)
}
