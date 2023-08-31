package db

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/genjidb/genji"
	"github.com/genjidb/genji/document"
	"github.com/genjidb/genji/types"
	"github.com/seekr-osint/seekr/api/names"
	"github.com/seekr-osint/seekr/api/person"

	_ "github.com/genjidb/genji/driver"
	"github.com/seekr-osint/seekr/api/config"
)

const (
	DBName = "PEOPLE"
)

func Init(cfg config.Config, ctx *context.Context) (*genji.DB, error) {
	if cfg.DataBasePath == "" {
		return nil, ErrEmptyDBPath
	}

	db, err := genji.Open(cfg.DataBasePath)
	if err != nil {
		return db, err
	}
	db = db.WithContext(*ctx)
	// defer db.Close()

	fmt.Println(GenerateCreateTableSQL(person.Person{}, DBName))
	err = db.Exec(GenerateCreateTableSQL(person.Person{}, DBName))
	if err != nil {
		return db, err
	}

	return db, nil
}
func Write(p person.Person, db *genji.DB) error {
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	return WriteTX(p, tx)
}

func Delete(id int, db *genji.DB) error {
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	return DeleteTX(id, tx)
}
func DeleteTX(id int, tx *genji.Tx) error {
	_, err := ReadTX(id, tx)
	if err != nil {
		return err
	}

	s := fmt.Sprintf("DELETE FROM %s WHERE id = %d", DBName, id)
	fmt.Println(s)
	err = tx.Exec(s)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func WriteTX(p person.Person, tx *genji.Tx) error {
	p2, err := ReadTX(int(p.ID), tx)
	if err == ErrPersonNotExist {
		err := tx.Exec(fmt.Sprintf(`INSERT INTO %s VALUES ?`, DBName), &p)
		if err != nil {
			tx.Rollback()
			return err
		}
		err = tx.Commit()
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	} else if err == nil {
		if !reflect.DeepEqual(p2, p) {
			set := CompareStructsAndGenerateString(p2, p)
			s := fmt.Sprintf("UPDATE %s SET %s WHERE id = %d", DBName, set, p.ID)
			fmt.Println(s)
			err := tx.Exec(s)
			if err != nil {
				log.Println(err)
				tx.Rollback()
				return err
			}
			err = tx.Commit()
			if err != nil {
				log.Println(err)
				return err
			}
			return nil
		}
	} else if err != nil {
		// weird error to handle FIXME
		return err
	}
	return tx.Rollback()
}

func ReadNamesTX(tx *genji.Tx) (names.Names, error) {
	res, err := tx.Query(fmt.Sprintf("SELECT name,id FROM %s", DBName))
	if err != nil {
		return names.Names{}, err
	}

	defer res.Close()
	data := names.Names{}
	err = res.Iterate(func(d types.Document) error {
		data1 := names.Name{}
		document.StructScan(d, &data1)
		data = append(data, data1)
		fmt.Println("i", data)
		return nil
	})
	if err != nil {
		return names.Names{}, err
	}
	if reflect.DeepEqual(data, names.Names{}) {
		return names.Names{}, ErrPersonNotExist
	}

	return data, nil
}

// NO tx.Rollback !!!!
func ReadTX(id int, tx *genji.Tx) (person.Person, error) {
	res, err := tx.Query(fmt.Sprintf("SELECT * FROM %s WHERE id = ?", DBName), id)
	if err != nil {
		return person.Person{}, err
	}

	defer res.Close()
	data := person.Person{}
	err = res.Iterate(func(d types.Document) error {
		document.StructScan(d, &data)
		return nil
	})
	if err != nil {
		return person.Person{}, err
	}
	if reflect.DeepEqual(data, person.Person{}) {
		return person.Person{}, ErrPersonNotExist
	}

	return data, nil
}

func ReadNames(db *genji.DB) (names.Names, error) {
	tx, err := db.Begin(false)
	defer tx.Rollback()
	if err != nil {
		return names.Names{}, err
	}
	defer tx.Rollback()
	return ReadNamesTX(tx)
}
func Read(id int, db *genji.DB) (person.Person, error) {
	tx, err := db.Begin(false)
	defer tx.Rollback()
	if err != nil {
		return person.Person{}, err
	}
	defer tx.Rollback()
	return ReadTX(id, tx)
}

func CompareStructsAndGenerateString(struct1, struct2 person.Person) string {
	structType := reflect.TypeOf(struct1)
	result := ""

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		fieldName := field.Name

		tag1 := field.Tag.Get("genji")

		val1 := reflect.ValueOf(struct1).FieldByName(fieldName).Interface()
		val2 := reflect.ValueOf(struct2).FieldByName(fieldName).Interface()

		if tag1 == "" {
			tag1 = fieldName
		}
		if tag1 != "" && !reflect.DeepEqual(val1, val2) {
			if result != "" {
				result += ", "
			}

			if field.Type.Kind() == reflect.String {
				result += fmt.Sprintf("%s = '%s'", tag1, val2)
			} else {
				result += fmt.Sprintf("%s = %v", tag1, val2)
			}
		}
	}

	return result
}
