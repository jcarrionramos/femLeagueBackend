package main

import (
	"log"

	models "github.com/jcarrionramos/femLeague-backend/fem-models"
	server "github.com/jcarrionramos/femLeague-backend/fem-server"
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
