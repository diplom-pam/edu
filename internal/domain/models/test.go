package models

type Test struct {
	ID        int    `json:"id" db:"id"`
	Title     string `json:"title" db:"title"`
	Questions []int  `json:"questions" db:"questions"`
	//QuestionCount int `json:"question_count" db:"question_count"`
}
