package database

import (
	"log"

	pgDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(dsn string) (*gorm.DB, error) {
	if dsn == "" {
		dsn = "host=tenant_db user=user password=password dbname=tenantdb port=54323 sslmode=disable"
	}

	db, err := gorm.Open(pgDriver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	return db, err
}
