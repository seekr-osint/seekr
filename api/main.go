package api

import (
	"net/http"

	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

var DatabaseFile string

// person represents data about a record person.
type person struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Age            int8      `json:"age"`
	Birthday       string    `json:"bday"`
	Address        string    `json:"address"`
	Phone          string    `json:"phone"`
	Civilstatus    string    `json:"civilstatus"`
	Kids           string    `json:"kids"`
	Hobbies        string    `json:"hobbies"`
	Email          string    `json:"email"`
	Occupation     string    `json:"occupation"`
	Prevoccupation string    `json:"prevoccupation"`
	Military       string    `json:"military"`
	Club           string    `json:"club"`
	Legal          string    `json:"legal"`
	Political      string    `json:"political"`
	Notes          string    `json:"notes"`
	Accounts       []Account `json:"accounts"`
}

type DataBase map[string]person

func handler(function func(DataBase, *gin.Context), db DataBase) gin.HandlerFunc {
	handlerFunc := func(c *gin.Context) {
		function(db, c)
	}
	return gin.HandlerFunc(handlerFunc)
}

func ServeApi(persons DataBase, ip string, databaseFile string) {
	fmt.Println("running api on" + ip)
	router := gin.Default()
	router.GET("/persons", handler(getPersons, persons))
	router.GET("/names", handler(getNamesRequest, persons))
	router.GET("/names/list", handler(getNamesListRequest, persons))
	router.GET("/names/list/len", handler(getNamesListLenRequest, persons))
	router.GET("/persons/:id", handler(getPersonByIDRequest, persons))
	router.POST("/persons", handler(postPersons, persons))
	router.DELETE("/persons/:id", handler(deletePerson, persons))
	DatabaseFile = databaseFile
	data, err := ioutil.ReadFile(DatabaseFile)
	if err != nil {
		fmt.Print(err)
	}
	err = json.Unmarshal(data, &persons)
	if err != nil {
		fmt.Println("error:", err)
	}

	router.Run(ip)
}

func GetStatusCode(url string) int {
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

func getPersons(persons DataBase, c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.IndentedJSON(http.StatusOK, persons)
}
func getNames(persons DataBase) map[string][]person {
	names := map[string][]person{}
	for _, element := range persons {
		names[element.Name] = append(names[element.Name], element)
	}
	return names
}
func getNamesList(persons DataBase) []string {
	names := []string{}
	for _, element := range persons {
		names = append(names, element.Name)
	}
	return names
}

func getNamesRequest(persons DataBase, c *gin.Context) {
	names := getNames(persons)
	c.IndentedJSON(http.StatusOK, names)
}

func getNamesListRequest(persons DataBase, c *gin.Context) {
	names := getNamesList(persons)
	c.IndentedJSON(http.StatusOK, names)
}

func getNamesListLenRequest(persons DataBase, c *gin.Context) {
	names := len(getNamesList(persons))
	c.IndentedJSON(http.StatusOK, names)
}

func deletePerson(persons DataBase, c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	if checkPersonExists(persons, c.Param("id")) {
		// Add the new person to the slice.
		delete(persons, c.Param("id"))
		//c.IndentedJSON(http.StatusCreated, newPerson)
		c.IndentedJSON(http.StatusAccepted, gin.H{"message": "deleted person"})
	}
	jsonBytes, err := json.Marshal(persons)
	if err != nil {
		fmt.Println("error:", err)
	}
	ioutil.WriteFile(DatabaseFile, jsonBytes, 0644)
}

func postPersons(persons DataBase, c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
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
	ioutil.WriteFile(DatabaseFile, jsonBytes, 0644)
}

func getPersonID(persons DataBase, id string) (bool, person, int) {
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

func getPersonByIDRequest(persons DataBase, c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.IndentedJSON(http.StatusOK, getPersonByID(persons, c.Param("id")))
}

func getPersonByID(persons DataBase, id string) person {
	var personToReturn person
	if checkPersonExists(persons, id) {
		personToReturn = persons[id]
	}
	return personToReturn
}
