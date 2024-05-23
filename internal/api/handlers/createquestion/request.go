package createquestion

type questionBody struct {
	Choices map[int]string `json:"choices"`
	Answer  []int          `json:"answer"`
	IsFull  bool           `json:"isFull"`
}

type request struct {
	Question string       `json:"question"`
	Body     questionBody `json:"body"`
	Content  string       `json:"content"`
}
