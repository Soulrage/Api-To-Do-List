package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitDB() error {
	dsn := "host=db user=Oleg password=Oleg dbname=ToDo port=5432"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func Migrate() error {
	if err := InitDB(); err != nil {
		return err
	}

	tables := []interface{}{
		&Tasks{},
	}
	for _, table := range tables {
		if !db.Migrator().HasTable(table) {
			if err := db.Migrator().CreateTable(table); err != nil {
				log.Printf("Failed to create table: %v", err)
				return err
			}
		}
	}

	return nil
}
