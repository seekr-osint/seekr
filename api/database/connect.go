package database

import (
	"github.com/glebarez/sqlite"
	"github.com/seekr-osint/seekr/api/person"
	"gorm.io/gorm"
)

// type AuthorizedDB struct{
// 	person.Person
// 	Username string
// }

func Connect(path string) (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&person.Person{})

	return db, nil
}
