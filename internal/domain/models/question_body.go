package models

type QuestionBody struct {
	Choices map[int]string `json:"choices" db:"choices"`
	Answer  []int          `json:"answer" db:"answer"`
	IsFull  bool           `json:"isFull" db:"isFull"`
}
