package main

import (
	"net/http"

	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

var databaseFile string = "./data.json"

// person represents data about a record person.
type person struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Age     int8   `json:"age"`
	Youtube string `json:"youtube"`
  Address string `json:"address"`
  Civilstatus string `json:"civilstatus"`
}

type DataBase map[string]person

// var persons = []person{}
var persons = make(DataBase)

func main() {
	fmt.Println(checkUsername("9glenda"))
	router := gin.Default()
	router.GET("/persons", getPersons)
	router.GET("/names", getNamesRequest)
  router.GET("/names/list", getNamesListRequest)
  router.GET("/names/list/len", getNamesListLenRequest)
	//router.GET("/names/:name", getPersonsByName)
	router.GET("/persons/:id", getPersonByIDRequest)
	router.POST("/persons", postPersons)

	data, err := ioutil.ReadFile(databaseFile)
	if err != nil {
		fmt.Print(err)
	}
	err = json.Unmarshal(data, &persons)
	if err != nil {
		fmt.Println("error:", err)
	}

	router.Run("localhost:8080")
}

func getStatusCode(url string) int {

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	if err != nil {
		fmt.Println(err)
	}

	return resp.StatusCode
}

func checkUsername(username string) []string {
	services := []string{
		"https://github.com/",
	}

	valid := []string{}
	for _, service := range services {
		if getStatusCode(service+username) == 200 {
			valid = append(valid, service+username)
		}
	}
	return valid
}

func getPersons(c *gin.Context) {
  c.Header("Access-Control-Allow-Origin","*")
	c.IndentedJSON(http.StatusOK, persons)
}
func getNames() map[string][]person {
	names := map[string][]person{}
	for _, element := range persons {
		names[element.Name] = append(names[element.Name], element)
	}
	return names
}
func getNamesList() []string {
	names := []string{}
	for _, element := range persons {
		names = append(names,element.Name)
	}
	return names
}



func getNamesRequest(c *gin.Context) {
	names := getNames()
	c.IndentedJSON(http.StatusOK, names)
}

func getNamesListRequest(c *gin.Context) {
	names := getNamesList()
	c.IndentedJSON(http.StatusOK, names)
}

func getNamesListLenRequest(c *gin.Context) {
	names := len(getNamesList())
	c.IndentedJSON(http.StatusOK, names)
}

func postPersons(c *gin.Context) {
	var newPerson person

	if err := c.BindJSON(&newPerson); err != nil {
		return
	}
	if !checkPersonExists(persons, newPerson.ID) {
		// Add the new person to the slice.
		persons[newPerson.ID] = newPerson
		c.IndentedJSON(http.StatusCreated, newPerson)
	} else {
		fmt.Println(persons[newPerson.ID])
		persons[newPerson.ID] = newPerson
		c.IndentedJSON(http.StatusAccepted, gin.H{"message": "overwritten person"})
	}

	jsonBytes, err := json.Marshal(persons)
	if err != nil {
		fmt.Println("error:", err)
	}
	ioutil.WriteFile(databaseFile, jsonBytes, 0644)
}

func getPersonID(id string) (bool, person, int) {
	var selectedPerson person
	var personExists = false
	var index int
	if _, ok := persons[id]; ok {
		personExists = true
		selectedPerson = persons[id]
	}
	return personExists, selectedPerson, index
}

func checkPersonExists(persons DataBase, id string) bool {
	_, ok := persons[id]
	return ok
}

func getPersonByIDRequest(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, getPersonByID(persons, c.Param("id")))
}

func getPersonByID(persons DataBase, id string) person {
	var personToReturn person
	if checkPersonExists(persons, id) {
		personToReturn = persons[id]
	}
	return personToReturn
}

func getPersonsByName(c *gin.Context) {
}
