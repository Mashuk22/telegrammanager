package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Mashuk22/telegrammanager/db"
	userpb "github.com/Mashuk22/telegrammanager/pkg/user_service"
	"google.golang.org/protobuf/proto"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://postgres:mypassword@localhost:6543/telegrammanager?sslmode=disable"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	db.New(conn)

	file, err := os.OpenFile("../output.bin", os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}

	result, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	params := &userpb.User{}
	err = proto.Unmarshal(result, params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("result:\n%v\n\n\n", params)
}
