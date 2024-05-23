package createtest

type request struct {
	QuestionIDs []int  `json:"question_ids"`
	Title       string `json:"title"`
}
