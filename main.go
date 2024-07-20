package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "password"
	DB_NAME     = "postgres-go"
)

func main() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
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
