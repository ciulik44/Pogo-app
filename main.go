package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Sprawdzenie, czy tabela istnieje i utworzenie jej, jeśli nie istnieje
	createTableQuery := `
    CREATE TABLE IF NOT EXISTS transactions (
        id SERIAL PRIMARY KEY,
        client_id INT NOT NULL,
        transaction_id INT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Error creating table: ", err)
	}

	clientID := 1
	transactionID := 1

	for {
		insertStmt := `INSERT INTO transactions(client_id, transaction_id) VALUES($1, $2)`
		_, err := db.Exec(insertStmt, clientID, transactionID)
		if err != nil {
			log.Println("Error inserting record: ", err)
		} else {
			log.Printf("Inserted record with client_id: %d and transaction_id: %d\n", clientID, transactionID)
			clientID++
			transactionID++
		}

		time.Sleep(2 * time.Minute)
	}
}
