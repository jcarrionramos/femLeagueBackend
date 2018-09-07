package structures

type Team struct {
	Name  string `json:"name,omitempty" db:"name"`
	Win   int    `json:"win" db:"win"`
	Draw  int    `json:"draw" db:"draw"`
	Loss  int    `json:"loss" db:"loss"`
	Total int    `json:"total" db:"total"`
}
