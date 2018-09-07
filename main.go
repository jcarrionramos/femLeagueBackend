package main

import (
	"log"

	"github.com/jcarrionramos/femLeagueBackend/models"
	"github.com/jcarrionramos/femLeagueBackend/server"
)

func main() {
	log.SetFlags(log.Lshortfile)
	log.Println("Starting FemLeague API-REST on port 9000")

	server := server.New()
	err := models.InitDB("./models/database.db")

	if err != nil {
		log.Println(err)
		return
	}

	server.Run(":9000")
}
