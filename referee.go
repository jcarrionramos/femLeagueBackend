package main

type Referee struct {
	Rut        int    `json:"rut,omitempty" "db:"rut"`
	FirstName  string `json:"first_name,omitempty" "db:"first_name"`
	LastName   string `json:"last_name,omitempty" "db:"last_name"`
	Email      string `json:"email,omitempty" "db:"email"`
	Phone      string `json:"phone,omitempty" "db:"phone"`
	YellowCard int    `json:"yellow_card,omitempty" "db:"yellow_card"`
	RedCard    int    `json:"red_card,omitempty" "db:"red_card"`
}
