package db

// import (
// 	// "context"
// 	// "fmt"
// 	"context"
// 	"fmt"

// 	// "fmt"
// 	// "log"

// 	// "database/sql"
// 	"github.com/genjidb/genji"
// 	"github.com/genjidb/genji/document"
// 	_ "github.com/genjidb/genji/driver"
// 	"github.com/genjidb/genji/types"
// 	"github.com/seekr-osint/seekr/api/config"
// 	"github.com/seekr-osint/seekr/api/person"
// )

// type User struct {
// 	ID   uint
// 	Name string
// 	Age  uint
// }

// func Init(cfg config.Config) error {
// 	if cfg.DataBasePath == "" {
// 		return ErrEmptyDBPath
// 	}
// 	fmt.Println("hello")
// 	fmt.Println(GenerateCreateTableSQL(person.Person{}, "PEOPLE"))
// 	fmt.Println("hello")

// 	db, err := genji.Open(cfg.DataBasePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer db.Close()

// 	// Create a user table if it doesn't exist
// 	err = db.Exec(`
//         CREATE TABLE IF NOT EXISTS people (
//             id              INT     PRIMARY KEY,
//             name            TEXT    NOT NULL,
// 						age 						INT
//         )
//     `)
// 	if err != nil {
// 		return err
// 	}

// 	db = db.WithContext(context.Background())

// 	u := User{
// 		ID:   2,
// 		Name: "foo",
// 	}
// 	err = db.Exec(`INSERT INTO PEOPLE VALUES ?`, &u)

// 	if err != nil {
// 		return err
// 	}

// 	// // Query some documents
// 	res, err := db.Query("SELECT * FROM PEOPLE WHERE id = ?", 1)
// 	if err != nil {
// 		return err
// 	}

// 	defer res.Close()
// 	user := User{}
// 	err = res.Iterate(func(d types.Document) error {
// 		document.Scan(d, &user.ID, &user.Name, &user.Age)
// 		return nil
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
