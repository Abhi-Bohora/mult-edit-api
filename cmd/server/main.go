package main

import (
	"fmt"
	"log"

	"github.com/Abhi-Bohora/multi-edit-api/config"
)

func main(){
	config, err := config.LoadConfig()
	if err != nil {
        log.Fatal("load config err:", err)
    }

	fmt.Println("Database Host:", config.Database.Host)
    fmt.Println("Server Port:", config.Server.Port)
	fmt.Println("Database Port:", config.Database.Port)

	connectToDatabase(config)
}

func connectToDatabase(cfg *config.Config) {
    dbHost := cfg.Database.Host
    dbPort := cfg.Database.Port
    dbUser := cfg.Database.User
    // just printing
    fmt.Printf("Connecting to database at %s:%s as user %s\n", dbHost, dbPort, dbUser)
}