package api

import (
	//"fmt"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

var DatabaseFile string

type SaveJsonFunc func(ApiConfig)
type ApiConfig struct {
	Ip            string       `json:"ip"`
	LogFile       string       `json:"log_file"`
	DataBaseFile  string       `json:"data_base_file"`
	DataBase      DataBase     `json:"data_base"`
	SetCORSHeader bool         `json:"set_CORS_header"`
	SaveJsonFunc  SaveJsonFunc `json:"save_json_func"`
	GinRouter     *gin.Engine  `json:"gin_router"`
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
	config.GinRouter.GET("/deep/github/:username", Handler(GithubInfoDeepRequest, config))                       // deep investigation of github account
	config.GinRouter.GET("/search/google/:query", Handler(GoogleRequest, config))                       // get results from google
	config.GinRouter.GET("/people/:id", Handler(GetPersonByIDRequest, config))                   // return person obj
	config.GinRouter.DELETE("/people/:id", Handler(DeletePerson, config))                        // delete person
	config.GinRouter.GET("/people/:id/delete", Handler(DeletePerson, config))                    // delete person
	config.GinRouter.DELETE("/people/:id/accounts/:account", Handler(DeleteAccount, config))     // delete account
	config.GinRouter.GET("/people/:id/accounts/:account/delete", Handler(DeleteAccount, config)) // delete account
	config.GinRouter.POST("/person", Handler(PostPerson, config))                                // post person
	config.GinRouter.GET("/getAccounts/:username", Handler(GetAccountsRequest, config))          // get accounts
	config.GinRouter.Run(config.Ip)
}


func GithubInfoDeepRequest(config ApiConfig, c *gin.Context) {
	if c.Param("username") != "" {
		c.IndentedJSON(http.StatusOK, GithubInfoDeep(c.Param("username"), true))
	}
}

func GetPersonByID(config ApiConfig, id string) (bool, Person) {
	if _, ok := config.DataBase[id]; ok {
		return true, config.DataBase[id]
	}
	return false, Person{}
}

func GetPersonByIDRequest(config ApiConfig, c *gin.Context) {
	exists, person := GetPersonByID(config, c.Param("id"))
	if exists {
		c.IndentedJSON(http.StatusOK, person)
	} else {
		c.IndentedJSON(http.StatusNotFound, nil)
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

func ParsePerson(newPerson Person) Person {
	newPerson = ReplaceNil(newPerson)
	newPerson = CheckMail(newPerson)
	return newPerson
}

func GetAccounts(config ApiConfig, username string) Accounts {
	return ServicesHandler(DefaultServices, username)
}

func GetAccountsRequest(config ApiConfig, c *gin.Context) {
	c.IndentedJSON(http.StatusOK, GetAccounts(config, c.Param("username")))
}
func ReplaceNil(newPerson Person) Person {
	if newPerson.Pictures == nil {
		newPerson.Pictures = Pictures{}
	}
	if newPerson.Accounts == nil {
		newPerson.Accounts = Accounts{}
	}
	if newPerson.Sources == nil {
		newPerson.Sources = Sources{}
	}
	return newPerson
}

// THIS HAS NO C.PARAM("id")
// THIS IS THE WORST PEICE OF BAD CODE I EVER WROTE
func PostPerson(config ApiConfig, c *gin.Context) { // c.BindJSON is a person not people (POST "localhost:8080/person")
	var newPerson Person

	// exit if the json is invalid
	if err := c.BindJSON(&newPerson); err != nil {
		return
	}
	// newPerson = CheckMail(newPerson) // FIXME
	newPerson = ParsePerson(newPerson)
	// DON'T BE LIKE ME AND USE NEWPERSON.ID !!!
	exsits, _ := GetPersonByID(config, newPerson.ID) // check rather the person Exsts
	if !exsits {
		// Add the new person to the database.
		if config.DataBase != nil {
			config.DataBase[newPerson.ID] = newPerson
		} else {
			config.DataBase = DataBase{}
			config.DataBase[newPerson.ID] = newPerson
		}
		c.IndentedJSON(http.StatusCreated, newPerson)
	} else {
		log.Printf("overwritten person %s", newPerson.ID)
		if config.DataBase != nil {
			config.DataBase[newPerson.ID] = newPerson
		} else {
			config.DataBase = DataBase{}
			config.DataBase[newPerson.ID] = newPerson
		}
		c.IndentedJSON(http.StatusAccepted, gin.H{"message": "overwritten person"})
	}
	config.SaveJsonFunc(config)
}

//func ServeApiOld(people DataBase, ip string, databaseFile string) {
//	log.Println("running api on" + ip)
//	//router := gin.Default()
//	//	router.GET("/people", handler(getPeople, people))
//	//	router.GET("/names", handler(getNamesRequest, people))
//	//	router.GET("/names/list", handler(getNamesListRequest, people))
//	//	router.GET("/github/:username/mail", handler(getGithubEmail, people))
//	//	router.GET("/github/:username/addMail/:id", handler(addGithubEmail, people))
//	//	router.GET("/names/list/len", handler(getNamesListLenRequest, people))
//	//	router.GET("/people/:id", handler(getPersonByIDRequest, people))
//	//	router.GET("/people/:id/accounts", handler(getPersonByIDRequestAccount, people))
//	//	router.GET("/people/:id/addAccounts/:username", handler(addAccounts, people))
//	//	router.POST("/people/:id/addAccount", handler(addAccount, people))
//	//	//router.GET("/people/:id/getAccounts/:username", handler(getAccountsRequest, people))
//	//	router.GET("/getAccounts/:username", handler(getAccountsRequest, people))
//	//	router.GET("/markdown/:id", handler(mdPersonByIDRequest, people))
//	//	router.POST("/people", handler(postPeople, people))
//	//	router.POST("/people/noAccounts", handler(postPeopleNoAccounts, people))
//	//	router.DELETE("/people/:id", handler(deletePerson, people))
//	//	router.GET("/people/:id/delete", handler(deletePerson, people))
//	//	router.POST("/dataJson", handler(writeDataJson, people))
//	//	DatabaseFile = databaseFile
//	//	data, err := ioutil.ReadFile(DatabaseFile)
//	//	if err != nil {
//	//		log.Println("error:", err)
//	//	}
//	//	err = json.Unmarshal(data, &people)
//	//	if err != nil {
//	//		log.Println("error:", err)
//	//	}
//	//
//	//	router.Run(ip)
//}
//
//
//func getPeople(people DataBase, c *gin.Context) {
//	c.IndentedJSON(http.StatusOK, people)
//}
//func getNames(people DataBase) map[string][]Person {
//	names := map[string][]Person{}
//	for _, element := range people {
//		names[element.Name] = append(names[element.Name], element)
//	}
//	return names
//}
//func getNamesList(people DataBase) []string {
//	names := []string{}
//	for _, element := range people {
//		names = append(names, element.Name)
//	}
//	return names
//}
//
//func getNamesRequest(people DataBase, c *gin.Context) {
//	names := getNames(people)
//	c.IndentedJSON(http.StatusOK, names)
//}
//
//func getNamesListRequest(people DataBase, c *gin.Context) {
//	names := getNamesList(people)
//	c.IndentedJSON(http.StatusOK, names)
//}
//
//func getNamesListLenRequest(people DataBase, c *gin.Context) {
//	names := len(getNamesList(people))
//	c.IndentedJSON(http.StatusOK, names)
//}
//
//func deletePerson(people DataBase, c *gin.Context) {
//	c.Header("Access-Control-Allow-Headers", "Content-Type,Content-Length,Server,Date,access-control-allow-methods,access-control-allow-origin")
//	if checkPersonExists(people, c.Param("id")) {
//		// Add the new person to the slice.
//		delete(people, c.Param("id"))
//		//c.IndentedJSON(http.StatusCreated, newPerson)
//		//c.IndentedJSON(http.StatusAccepted, gin.H{"message": "deleted person"})
//	}
//	jsonBytes, err := json.Marshal(people)
//	if err != nil {
//		log.Println("error:", err)
//	}
//	ioutil.WriteFile(DatabaseFile, jsonBytes, 0644)
//}
//
//func postPeople(people DataBase, c *gin.Context) { // c.BindJSON is a person not people
//	var newPerson person
//
//	if err := c.BindJSON(&newPerson); err != nil {
//		return
//	}
//	newPerson = CheckMail(newPerson)
//	if !checkPersonExists(people, newPerson.ID) {
//		// Add the new person to the slice.
//		people[newPerson.ID] = newPerson
//		c.IndentedJSON(http.StatusCreated, newPerson)
//	} else {
//		log.Println(people[newPerson.ID])
//		people[newPerson.ID] = newPerson
//		c.IndentedJSON(http.StatusAccepted, gin.H{"message": "overwritten person"})
//	}
//	SaveJson(people)
//}
//
//func rm(indexes []int, strs []EmailServiceEnum) []EmailServiceEnum { // TODO write test
//	// Create a map to keep track of which indexes to remove
//	remove := make(map[int]bool)
//	for _, i := range indexes {
//		remove[i] = true
//	}
//
//	// Iterate through the input EmailServiceEnums and only append to the result
//	// if the current index is not in the map of indexes to remove
//	result := make([]EmailServiceEnum, 0, len(strs)-len(indexes))
//	for i, s := range strs {
//		if !remove[i] {
//			result = append(result, s)
//		}
//	}
//	return result
//}
//
//
//func postPeopleNoAccounts(people DataBase, c *gin.Context) { // you only get a person noy people
//	var newPerson person
//
//	if err := c.BindJSON(&newPerson); err != nil {
//		return
//	}
//
//	newPerson = CheckMail(newPerson)
//	if !checkPersonExists(people, newPerson.ID) {
//		// Add the new person to the slice.
//		people[newPerson.ID] = newPerson
//		c.IndentedJSON(http.StatusCreated, newPerson)
//	} else {
//		log.Println(people[newPerson.ID])
//		var accounts = people[newPerson.ID].Accounts
//		newPerson.Accounts = accounts
//		people[newPerson.ID] = newPerson
//		c.IndentedJSON(http.StatusAccepted, gin.H{"message": "overwritten person"})
//	}
//	SaveJson(people)
//}
//
//func getPersonID(people DataBase, id string) (bool, person, int) {
//	var selectedPerson person
//	var personExists = false
//	var index int
//	if _, ok := people[id]; ok {
//		personExists = true
//		selectedPerson = people[id]
//	}
//	return personExists, selectedPerson, index
//}
//
//func checkPersonExists(people DataBase, id string) bool {
//	_, ok := people[id]
//	return ok
//}
//
//func getPersonByIDRequest(people DataBase, c *gin.Context) {
//	c.IndentedJSON(http.StatusOK, getPersonByID(people, c.Param("id")))
//}
//
//func getPersonByIDRequestAccount(people DataBase, c *gin.Context) {
//	c.IndentedJSON(http.StatusOK, getPersonByID(people, c.Param("id")).Accounts)
//}
//
//func mdPersonByIDRequest(people DataBase, c *gin.Context) {
//	fmt.Println(GenMD(getPersonByID(people, c.Param("id"))))
//	c.IndentedJSON(http.StatusOK, GenMD(getPersonByID(people, c.Param("id"))))
//}
//
//func addAccounts(people DataBase, c *gin.Context) {
//	people = getAccounts(people, c.Param("id"), c.Param("username"))
//	SaveJson(people)
//	c.IndentedJSON(http.StatusOK, getPersonByID(people, c.Param("id")))
//}
//
//func writeDataJson(people DataBase, c *gin.Context) {
//	if err := c.BindJSON(&people); err != nil {
//		log.Println("bindjson error")
//		return
//	}
//	SaveJson(people)
//}
//
//func addAccount(people DataBase, c *gin.Context) {
//
//	var account Account
//
//	if err := c.BindJSON(&account); err != nil {
//		log.Println("bindjson error")
//		return
//	}
//	exists := false
//	personToAdd := getPersonByID(people, c.Param("id"))
//	if personToAdd.Accounts == nil {
//		personToAdd.Accounts = Accounts{}
//	} else {
//		for _, element := range personToAdd.Accounts {
//			if element.Username == account.Username && element.Service == account.Service {
//				exists = true
//				log.Println("username already exists")
//			}
//		}
//	}
//	if !exists {
//		personToAdd.Accounts = append(personToAdd.Accounts, account)
//		people[c.Param("id")] = personToAdd
//
//		//people[c.Param("id")].Accounts,account)
//		SaveJson(people)
//		c.IndentedJSON(http.StatusOK, getPersonByID(people, c.Param("id")))
//	}
//}
//
//func getAccounts(people DataBase, id, username string) DataBase {
//	person := getPersonByID(people, id)
//	person.Accounts = ServicesHandler(DefaultServices, username)
//	people[id] = person
//	return people
//}
//
//func getAccountsSimple(username string) Accounts {
//	// remove @ if it is the first char
//	if username[0:1] == "@" {
//		log.Println("@ is first")
//		username = username[1:]
//	}
//
//	return ServicesHandler(DefaultServices, username)
//}
//
//func getAccountsRequest(people DataBase, c *gin.Context) {
//	//if c.Param("username") != "" {
//	c.IndentedJSON(http.StatusOK, getAccountsSimple(c.Param("username")))
//	//}
//}
//
//func getGithubEmail(people DataBase, c *gin.Context) {
//	if c.Param("username") != "" {
//		c.IndentedJSON(http.StatusOK, GithubInfoDeep(c.Param("username"), true))
//	}
//}
//
//func addGithubEmail(people DataBase, c *gin.Context) {
//	if c.Param("username") != "" {
//		//people[c.Param("id")].Email = append(getPersonByID(people, c.Param("id")).Email,GithubInfoDeep(c.Param("username"),true)[:]...)
//		if entry, ok := people[c.Param("id")]; ok {
//			entry.Email = append(getPersonByID(people, c.Param("id")).Email, GithubInfoDeep(c.Param("username"), true)[:]...)
//			entry = CheckMail(entry)
//			people[c.Param("id")] = entry
//		}
//		//people[c.Param("id")].Email = getPersonByID(people, c.Param("id")).Email
//		//c.IndentedJSON(http.StatusOK, GithubInfoDeep(c.Param("username"), true))
//		SaveJson(people)
//	}
//}
//
//func getPersonByID(people DataBase, id string) person {
//	var personToReturn person
//	if checkPersonExists(people, id) {
//		personToReturn = people[id]
//	}
//	return personToReturn
//}
