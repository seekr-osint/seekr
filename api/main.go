package api

import (
	"net/http"

	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"log"
)

var DatabaseFile string

func handler(function func(DataBase, *gin.Context), db DataBase) gin.HandlerFunc {
	handlerFunc := func(c *gin.Context) {
		file, _ := ioutil.ReadFile("data.json")
		err := json.Unmarshal(file, &db)
		if err != nil {
			log.Println(err)
		}
		function(db, c)
	}
	return gin.HandlerFunc(handlerFunc)
}

func ServeApi(people DataBase, ip string, databaseFile string) {
	log.Println("running api on" + ip)
	router := gin.Default()
	router.GET("/people", handler(getPeople, people))
	router.GET("/names", handler(getNamesRequest, people))
	router.GET("/names/list", handler(getNamesListRequest, people))
	router.GET("/names/list/len", handler(getNamesListLenRequest, people))
	router.GET("/people/:id", handler(getPersonByIDRequest, people))
	router.GET("/people/:id/addAccounts/:username", handler(addAccounts, people))
	router.GET("/people/:id/getAccounts/:username", handler(getAccountsRequest, people))
	router.POST("/people", handler(postPeople, people))
	router.DELETE("/people/:id", handler(deletePerson, people))
	DatabaseFile = databaseFile
	data, err := ioutil.ReadFile(DatabaseFile)
	if err != nil {
		log.Println("error:", err)
	}
	err = json.Unmarshal(data, &people)
	if err != nil {
		log.Println("error:", err)
	}

	router.Run(ip)
}

func GetStatusCode(url string) int {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
	}
	return resp.StatusCode
}

func getPeople(people DataBase, c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.IndentedJSON(http.StatusOK, people)
}
func getNames(people DataBase) map[string][]person {
	names := map[string][]person{}
	for _, element := range people {
		names[element.Name] = append(names[element.Name], element)
	}
	return names
}
func getNamesList(people DataBase) []string {
	names := []string{}
	for _, element := range people {
		names = append(names, element.Name)
	}
	return names
}

func getNamesRequest(people DataBase, c *gin.Context) {
	names := getNames(people)
	c.IndentedJSON(http.StatusOK, names)
}

func getNamesListRequest(people DataBase, c *gin.Context) {
	names := getNamesList(people)
	c.IndentedJSON(http.StatusOK, names)
}

func getNamesListLenRequest(people DataBase, c *gin.Context) {
	names := len(getNamesList(people))
	c.IndentedJSON(http.StatusOK, names)
}

func deletePerson(people DataBase, c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	if checkPersonExists(people, c.Param("id")) {
		// Add the new person to the slice.
		delete(people, c.Param("id"))
		//c.IndentedJSON(http.StatusCreated, newPerson)
		c.IndentedJSON(http.StatusAccepted, gin.H{"message": "deleted person"})
	}
	jsonBytes, err := json.Marshal(people)
	if err != nil {
		log.Println("error:", err)
	}
	ioutil.WriteFile(DatabaseFile, jsonBytes, 0644)
}

func postPeople(people DataBase, c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	var newPerson person

	if err := c.BindJSON(&newPerson); err != nil {
		return
	}
	if !checkPersonExists(people, newPerson.ID) {
		// Add the new person to the slice.
		people[newPerson.ID] = newPerson
		c.IndentedJSON(http.StatusCreated, newPerson)
	} else {
		log.Println(people[newPerson.ID])
		people[newPerson.ID] = newPerson
		c.IndentedJSON(http.StatusAccepted, gin.H{"message": "overwritten person"})
	}
	SaveJson(people)
}

func getPersonID(people DataBase, id string) (bool, person, int) {
	var selectedPerson person
	var personExists = false
	var index int
	if _, ok := people[id]; ok {
		personExists = true
		selectedPerson = people[id]
	}
	return personExists, selectedPerson, index
}

func checkPersonExists(people DataBase, id string) bool {
	_, ok := people[id]
	return ok
}

func getPersonByIDRequest(people DataBase, c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.IndentedJSON(http.StatusOK, getPersonByID(people, c.Param("id")))
}

func addAccounts(people DataBase, c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	people = getAccounts(people, c.Param("id"), c.Param("username"))
	SaveJson(people)
	c.IndentedJSON(http.StatusOK, getPersonByID(people, c.Param("id")))
}

func getAccounts(people DataBase, id, username string) DataBase {
	person := getPersonByID(people, id)
	person.Accounts = ServicesHandler(DefaultServices, username)
	people[id] = person
	return people
}

func getAccountsRequest(people DataBase, c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.IndentedJSON(http.StatusOK, getPersonByID(getAccounts(people, c.Param("id"), c.Param("username")), c.Param("id")))
}

func getPersonByID(people DataBase, id string) person {
	var personToReturn person
	if checkPersonExists(people, id) {
		personToReturn = people[id]
	}
	return personToReturn
}
