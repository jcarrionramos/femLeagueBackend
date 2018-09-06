package structures

type Team struct {
	Name  string `json:"name,omitempty" db:"name"`
	Win   int    `json:"win,omitempty" db:"win"`
	Draw  int    `json:"draw,omitempty" db:"draw"`
	Loss  int    `json:"loss,omitempty" db:"loss"`
	Total int    `json:"total,omitempty" db:"total"`
}
