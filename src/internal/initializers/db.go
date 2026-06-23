package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB(dbURL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}