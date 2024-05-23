package createquestion

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/diplom-pam/edu/internal/domain/models"
	"github.com/diplom-pam/edu/internal/logger"
	"github.com/diplom-pam/edu/internal/utils"
	"go.uber.org/zap"
)

type store interface {
	CreateQuestion(ctx context.Context, question models.Question) (id int, err error)
}

func Handler(rw http.ResponseWriter, req *http.Request, store store) {
	r := request{}

	errDecode := json.NewDecoder(req.Body).Decode(&r)
	if errDecode != nil {
		logger.Debug("error decoding request", zap.Error(errDecode))
		utils.ErrorResponse(rw, http.StatusBadRequest, "bad request")
		return
	}

	// todo: validate request

	q := models.Question{
		Question: r.Question,
		Body: models.QuestionBody{
			Choices: r.Body.Choices,
			Answer:  r.Body.Answer,
			IsFull:  r.Body.IsFull,
		},
		Content: r.Content,
	}

	questionID, errGetQuestion := store.CreateQuestion(req.Context(), q)
	if errGetQuestion != nil {
		logger.Error("error getting domain", zap.Error(errGetQuestion))
		utils.ErrorResponse(rw, http.StatusInternalServerError, "internal error")
		return
	}

	resp := response{
		Data: responseItem{
			ID:       questionID,
			Question: r.Question,
			Body: responseQuestionBody{
				Choices: r.Body.Choices,
				Answer:  r.Body.Answer,
				IsFull:  r.Body.IsFull,
			},
			Content: r.Content,
		},
	}

	utils.Response(rw, resp)
}
