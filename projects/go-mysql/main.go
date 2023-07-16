package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Init sleep until mysql ready
	initWait := 15 * time.Second
	time.Sleep(initWait)
	// Retry connecting to MySQL for up to 30 seconds
	maxRetries := 30
	retryInterval := 1 * time.Second

	var db *sql.DB
	var err error

	for i := 0; i < maxRetries; i++ {
		db, err = sql.Open("mysql", "root:root@tcp(mysql:3306)/mydatabase")
		if err != nil {
			log.Printf("Error connecting to MySQL: %v", err)
			time.Sleep(retryInterval)
			continue
		}

		err = db.Ping()
		if err != nil {
			log.Printf("Error pinging MySQL: %v", err)
			time.Sleep(retryInterval)
			continue
		}
		log.Printf("Successfully connected to MySQL")
		break
	}

	if db == nil {
		log.Fatal("Unable to connect to MySQL")
	}

	defer db.Close()
	log.Printf("GoodBye Steven!")

	// MySQL connection is ready, continue with your application logic
	// ...
}
