package models

type Topic struct {
	ID             int    `json:"id" db:"id"`
	Title          string `json:"title" db:"title"`
	Body           string `json:"body" db:"body"`
	Content        string `json:"content" db:"content"`
	Tests          []int  `json:"tests" db:"tests"`
	Achievement    string `json:"achievement" db:"achievement"`
	IsTestRequired bool   `json:"is_test_required" db:"is_test_required"`
}
