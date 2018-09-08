package structures

type Match struct {
	Local_name  string `json"local_name,omitempty" db"local_name"`
	Visit_name  string `json"visit_name,omitempty" db"visit_name"`
	Day         int    `json"day,omitempty" db"day"`
	Season      string `json:"season,omitempty" db:"season"`
	Local_score int    `json"local_score,omitempty" db"local_score"`
	Visit_score int    `json"visit_score,omitempty" db"visit_score"`
	Referee_rut string `json"referee_rut,omitempty" db"referee_rut"`
	Played      int    `json"played,omitempty" db"played"`
}
