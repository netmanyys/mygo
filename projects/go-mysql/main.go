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



package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

// Define a struct to represent your data model
type User struct {
    gorm.Model
    Name  string
    Email string
}

func main() {
    // Set up the database connection
    dsn := "username:password@tcp(db:3306)/database_name?parseTime=true"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }

    // Migrate the database schema
    err = db.AutoMigrate(&User{})
    if err != nil {
        log.Fatal(err)
    }

    // Create a new router
    router := mux.NewRouter()

    // Define the API endpoints
    router.HandleFunc("/users", getUsers).Methods("GET")
    router.HandleFunc("/users/{id}", getUser).Methods("GET")
    router.HandleFunc("/users", createUser).Methods("POST")
    router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
    router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

    // Start the HTTP server
    log.Fatal(http.ListenAndServe(":8080", router))
}

// Handler functions for each API endpoint
func getUsers(w http.ResponseWriter, r *http.Request) {
    // Retrieve all users from the database
    var users []User
    db.Find(&users)
    fmt.Fprint(w, users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
    // Retrieve a specific user by ID from the database
    vars := mux.Vars(r)
    id := vars["id"]

    var user User
    db.First(&user, id)
    fmt.Fprint(w, user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
    // Create a new user in the database
    var user User
    json.NewDecoder(r.Body).Decode(&user)

    db.Create(&user)
    fmt.Fprint(w, user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
    // Update an existing user in the database
    vars := mux.Vars(r)
    id := vars["id"]

    var user User
    db.First(&user, id)
    json.NewDecoder(r.Body).Decode(&user)

    db.Save(&user)
    fmt.Fprint(w, user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
    // Delete a user from the database
    vars := mux.Vars(r)
    id := vars["id"]

    var user User
    db.Delete(&user, id)
    fmt.Fprint(w, "User deleted successfully")
}