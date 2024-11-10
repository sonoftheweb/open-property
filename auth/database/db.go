package database

import (
    "log"

    pgDriver "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func ConnectDB(dsn string) (*gorm.DB, error) {
    if dsn == "" {
        dsn = "host=auth_db user=user password=password dbname=authdb port=5432 sslmode=disable"
    }
    
    db, err := gorm.Open(pgDriver.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }
    return db, err
}