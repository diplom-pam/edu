package getquestion

type responseQuestionBody struct {
	Choices map[int]string `json:"choices"`
	Answer  []int          `json:"answer"`
	IsFull  bool           `json:"isFull"`
}

type responseItem struct {
	ID       int                  `json:"id"`
	Question string               `json:"question"`
	Body     responseQuestionBody `json:"body"`
	Content  string               `json:"content"`
}

// @swagger-component domainCreateResponse
type response struct {
	Data responseItem `json:"data"`
}
