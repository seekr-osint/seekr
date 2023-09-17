package database

import "gorm.io/gorm"

type DBStorage struct {
	DB *gorm.DB
}

func (dbs DBStorage) Reset() error {
	// newdb, err := dbs.DB.DB()
	// if err != nil {
	// 	return err
	// }
	return nil
}

//	func (dbs DBStorage) Delete(id string) error {
//		return dbs.DB.Delete()
//	}
func (dbs DBStorage) Close() error {
	newdb, err := dbs.DB.DB()
	if err != nil {
		return err
	}
	return newdb.Close()
}
