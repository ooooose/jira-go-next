package util

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/ooooose/jira_go/config"
    "log"
)

var DB *gorm.DB

func InitDB(cfg *config.Config) {
    var err error
    DB, err = gorm.Open(postgres.Open(cfg.Database.DSN), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }
}
