package main

type Player struct {
	Rut          int    `json:"rut,omitempty" "db:"rut"`
	FirstName    string `json:"first_name,omitempty" "db:"first_name"`
	LastName     string `json:"last_name,omitempty" "db:"last_name"`
	Email        string `json:"email,omitempty" "db:"email"`
	Phone        string `json:"phone,omitempty" "db:"phone"`
	TeamName     string `json:"team_name,omitempty" "db:"team_name"`
	DorsalNumber string `json:"dorsal_number,omitempty" "db:"dorsal_number"`
	Score        int    `json:"score,omitempty" "db:"score"`
}
