package models

import (
	"log"

	structures "github.com/jcarrionramos/femLeague-backend/fem-structures"
	_ "github.com/mattn/go-sqlite3"
)

// db is a global variable defined in /models/db.go

func SelectAllTeams() (teams []structures.Team, err error) {
	rows, err := db.Query("SELECT * FROM teams ORDER BY total DESC")

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var win, draw, loss, total int
		err = rows.Scan(&name, &win, &draw, &loss, &total)

		if err != nil {
			log.Println(err)
			return nil, err

		}

		teams = append(teams, structures.Team{
			Name:  name,
			Win:   win,
			Draw:  draw,
			Loss:  loss,
			Total: total,
		})
	}

	return teams, nil
}
