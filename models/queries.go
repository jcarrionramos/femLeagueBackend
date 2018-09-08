package models

import (
	"log"

	"github.com/jcarrionramos/femLeagueBackend/structures"
	_ "github.com/mattn/go-sqlite3"
)

// db is a global variable defined in /models/db.go

func SelectAllTeams() (teams []structures.Team, err error) {
	rows, err := db.Query(`SELECT * FROM teams ORDER BY total DESC`)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var win, draw, loss, total, active int
		err = rows.Scan(&name, &win, &draw, &loss, &total, &active)

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

func SelectNextsMatches() (matches []structures.Match, err error) {
	rows, err := db.Query(`SELECT local_name, visit_name, day FROM matches WHERE
		played = 0 ORDER BY day ASC`)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var local, visit string
		var day int
		err = rows.Scan(&local, &visit, &day)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		matches = append(matches, structures.Match{
			Local_name: local,
			Visit_name: visit,
			Day:        day,
		})
	}

	return matches, nil
}

func SelectTopScorers() (players []structures.Player, err error) {
	rows, err := db.Query(`SELECT first_name, last_name, score, dorsal_number,
		team_name FROM players ORDER BY score DESC LIMIT 3`)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var firstName, lastName, dorsalNumber, teamName string
		var score int
		err = rows.Scan(&firstName, &lastName, &score, &dorsalNumber, &teamName)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		players = append(players, structures.Player{
			FirstName:    firstName,
			LastName:     lastName,
			Score:        score,
			DorsalNumber: dorsalNumber,
			TeamName:     teamName,
		})
	}

	return players, nil
}
