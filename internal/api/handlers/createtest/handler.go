package createtest

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
	CreateTest(ctx context.Context, test models.Test) (id int, err error)
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

	t := models.Test{
		Questions: r.QuestionIDs,
		Title:     r.Title,
	}

	questionID, errGetQuestion := store.CreateTest(req.Context(), t)
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
