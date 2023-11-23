package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectToSQLite() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("arknet.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
