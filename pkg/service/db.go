package service

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/repodevs/gomis/pkg/util"
)

// define table query
const tableCreationQuery = `CREATE TABLE IF NOT EXISTS notes (
    id SERIAL PRIMARY KEY,
    title VARCHAR(45) NULL,
    content TEXT NULL,
    author VARCHAR(45) NULL
);`

// Init Database Connection
func (service *Server) DBInit(dbHost, dbPort, dbUser, dbPass, dbName string) {

	var err error
	log.Println(util.ConnectingToDB)

	conStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
	fmt.Println(conStr)

	// Open connection (only validating the connection arguments
	service.DBCon, err = sql.Open("postgres", conStr)

	if err != nil {
		log.Fatal(err)
	}

	// Make real connection to database
	err = service.DBCon.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(util.DBConnectionSuccess)

	// Init table creation
	if _, err := service.DBCon.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}

}
