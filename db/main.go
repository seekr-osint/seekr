package main

import (
	"encoding/json"
	"fmt"
	"log"

	badger "github.com/dgraph-io/badger/v3"
)

type Person struct {
	Id      int
	Name    string
	Age     int
	Address string
}

func main() {
	opts := badger.DefaultOptions("./badger")
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	people := []Person{
		{Id: 1, Name: "Alice", Age: 30, Address: "123 Main St"},
		{Id: 2, Name: "Bob", Age: 25, Address: "456 Elm St"},
	}

	for _, person := range people {
		err := person.Save(db)
		if err != nil {
			log.Fatal(err)
		}
	}

	person, err := Read(db, 1)
	fmt.Printf("%s,%d,%s", person.Name, person.Age, person.Address)
	if err != nil {
		log.Fatal(err)
	}
}

func Read(db *badger.DB, id int) (Person, error) {
	var person Person
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(fmt.Sprintf("person-%d", id)))
		if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {
			err := json.Unmarshal(val, &person) // Deserialize using JSON
			return err
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return Person{}, err
	}
	return person, nil

}
func (person Person) Save(db *badger.DB) error {
	if person.Id == 0 {
		return fmt.Errorf("missing ID")
	}
	key := []byte(fmt.Sprintf("person-%d", person.Id))
	value, err := json.Marshal(person) // Serialize using JSON
	if err != nil {
		return err
	}
	err = db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, value)
	})
	return err
}
