package main

type Match struct {
	Local_name  string `json"local_name,omitempty" db"local_name"`
	Visit_name  string `json"visit_name,omitempty" db"visit_name"`
	Local_goal  int    `json"local_goal,omitempty" db"local_goal"`
	Visit_goal  int    `json"visit_goal,omitempty" db"visit_goal"`
	Referee_rut string `json"referee_rut,omitempty" db"referee_rut"`
	Day         int    `json"day,omitempty" db"day"`
	Season      string `json:"season,omitempty" db:"season"`
	Played      int    `json"played,omitempty" db"played"`
}
