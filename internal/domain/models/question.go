package models

type Question struct {
	ID       int          `json:"id" db:"id"`
	Question string       `json:"question" db:"question"`
	Body     QuestionBody `json:"body" db:"body"`
	Content  string       `json:"content" db:"content"`
}
