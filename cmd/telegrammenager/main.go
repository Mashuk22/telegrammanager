package main

import (
	"database/sql"
	"log"

	"github.com/Mashuk22/telegrammanager/db"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:mypassword@localhost:6543/telegrammanager?sslmode=disable"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	db.New(conn)

}
