package main

import (
	"github.com/repodevs/gomis/pkg/service"
	"github.com/repodevs/gomis/pkg/util"
)

func main() {

	dbHost := util.Getenv("DB_HOST", "127.0.0.1")
	dbPort := util.Getenv("DB_PORT", "5789")
	dbUser := util.Getenv("DB_USER", "gomis")
	dbPass := util.Getenv("DB_PASS", "qweasd123")
	dbName := util.Getenv("DB_NAME", "gomisdb")

	serverPort := 9090

	service := service.Server{}

	// Init services
	service.Init()
	// Connect database
	service.ConnectDB(dbHost, dbPort, dbUser, dbPass, dbName)
	// Start the server
	service.Start(serverPort)
}
