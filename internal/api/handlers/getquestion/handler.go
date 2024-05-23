package getquestion

import (
	"context"
	"net/http"
	"strconv"

	"github.com/diplom-pam/edu/internal/domain/models"
	"github.com/diplom-pam/edu/internal/logger"
	"github.com/diplom-pam/edu/internal/utils"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type store interface {
	GetQuestion(ctx context.Context, questionID int) (*models.Question, error)
}

func Handler(rw http.ResponseWriter, req *http.Request, store store) {

	questionID, errGet := strconv.Atoi(chi.URLParam(req, "id"))
	if errGet != nil {
		logger.Debug("error getting question id", zap.Error(errGet))
		utils.ErrorResponse(rw, http.StatusBadRequest, "bad request")
		return
	}

	// todo: validate request

	question, errGetQuestion := store.GetQuestion(req.Context(), questionID)
	if errGetQuestion != nil {
		logger.Error("error getting domain", zap.Error(errGetQuestion))
		utils.ErrorResponse(rw, http.StatusInternalServerError, "internal error")
		return
	}

	resp := response{
		Data: responseItem{
			ID:       question.ID,
			Question: question.Question,
			Body: responseQuestionBody{
				Choices: question.Body.Choices,
				Answer:  question.Body.Answer,
				IsFull:  question.Body.IsFull,
			},
			Content: question.Content,
		},
	}

	utils.Response(rw, resp)
}
