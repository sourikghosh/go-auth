package main

import (
	"log"

	"auth/implementation/auth"
	"auth/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Loading Configuration failed: %s", err.Error())
	}

	conn, err := gorm.Open(mysql.Open(cfg.DB_DSN), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := conn.AutoMigrate(&auth.User{}); err != nil {
		log.Fatal(err.Error())
	}

	log.Println("successfully migrated")
}
