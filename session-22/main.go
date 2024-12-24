package main

import (
	log "github.com/sirupsen/logrus"
	"session-22/db"
	"session-22/logger"
	"session-22/server"
)

func main() {
	if err := logger.Init(); err != nil {
		log.Fatal(err)
	}

	if err := db.InitDB(); err != nil {
		log.Fatal(err)
	}

	if err := server.Init(); err != nil {
		log.Fatal(err)
	}
}
