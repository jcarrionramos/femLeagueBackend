package models

import (
	"log"

	"github.com/jcarrionramos/femLeagueBackend/structures"
	_ "github.com/mattn/go-sqlite3"
)

// db is a global variable defined in /models/db.go

func SelectActivesTeams() (teams []structures.Team, err error) {
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

func SelectAllPlayers() (players []structures.Player, err error) {
	rows, err := db.Query(`SELECT rut, first_name, last_name, dorsal_number,
	team_name FROM players ORDER BY team_name DESC`)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var firstName, lastName, dorsalNumber, teamName string
		var rut int
		err = rows.Scan(&rut, &firstName, &lastName, &dorsalNumber, &teamName)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		players = append(players, structures.Player{
			Rut:          rut,
			FirstName:    firstName,
			LastName:     lastName,
			DorsalNumber: dorsalNumber,
			TeamName:     teamName,
		})
	}

	return players, nil
}

func SelectAllReferees() (referees []structures.Referee, err error) {
	rows, err := db.Query(`SELECT rut, first_name, last_name, email
		 FROM referees`)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var firstName, lastName, email string
		var rut int
		err = rows.Scan(&rut, &firstName, &lastName, &email)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		referees = append(referees, structures.Referee{
			Rut:       rut,
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
		})
	}

	return referees, nil
}

func InsertTeam(team structures.Team) error {
	sqlstmt := `INSERT INTO teams VALUES ($1,$2,$3,$4,$5,$6)`

	_, err := db.Exec(sqlstmt, team.Name, team.Win, team.Draw,
		team.Loss, team.Total, 0)

	if err != nil {
		log.Println(err)
		return err

	}
	return nil
}

func InsertPlayer(player structures.Player) error {
	sqlstmt := `INSERT INTO players VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`

	_, err := db.Exec(sqlstmt, player.Rut, player.FirstName, player.LastName,
		player.Email, player.Phone, player.TeamName, player.DorsalNumber, 0)

	if err != nil {
		log.Println(err)
		return err

	}
	return nil
}

func InsertReferee(referee structures.Referee) error {
	sqlstmt := `INSERT INTO referees VALUES ($1,$2,$3,$4,$5,$6,$7)`

	_, err := db.Exec(sqlstmt, referee.Rut, referee.FirstName, referee.LastName,
		referee.Email, referee.Phone, 0, 0)

	if err != nil {
		log.Println(err)
		return err

	}
	return nil
}

func InsertMatch(local, visit, season string, day int) error {
	sqlstmt := `INSERT INTO matches VALUES ($1,$2,$3,$4, 0, 0, 0, "null")`

	_, err := db.Exec(sqlstmt, local, visit, season, day)

	if err != nil {
		log.Println(err)
		return err

	}
	return nil
}

func DeleteTeam(name string) error {
	sqlstmt := `DELETE FROM teams WHERE name=$1`

	_, err := db.Exec(sqlstmt, name)

	if err != nil {
		log.Println(err)
		return err

	}
	return nil
}

func DeletePlayer(rut int) error {
	sqlstmt := `DELETE FROM players WHERE rut=$1`

	_, err := db.Exec(sqlstmt, rut)

	if err != nil {
		log.Println(err)
		return err

	}
	return nil
}

func DeleteReferee(rut int) error {
	sqlstmt := `DELETE FROM referees WHERE rut=$1`

	_, err := db.Exec(sqlstmt, rut)

	if err != nil {
		log.Println(err)
		return err

	}
	return nil
}
