package models

type Course struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Tests       []int  `json:"tests" db:"tests"`
	Achievement string `json:"achievement" db:"achievement"`
}
