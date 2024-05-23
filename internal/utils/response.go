package utils

import (
	"encoding/json"
	"net/http"

	"github.com/diplom-pam/edu/internal/logger"
	"go.uber.org/zap"
)

type apiErrorResponseMessage struct {
	Message string `json:"message,omitempty"`
}

func Response(rw http.ResponseWriter, result any) {
	rw.Header().Set("Content-Type", "application/json")
	errEncode := json.NewEncoder(rw).Encode(result)
	if errEncode != nil {
		logger.Error("encode response error", zap.Error(errEncode))
		ErrorResponse(rw, http.StatusInternalServerError, "internal server error")
	}
}

func ErrorResponse(rw http.ResponseWriter, code int, message string) {
	rw.WriteHeader(code)

	resp := apiErrorResponseMessage{
		Message: message,
	}

	errEncode := json.NewEncoder(rw).Encode(resp)
	if errEncode != nil {
		logger.Error("encode response error", zap.Error(errEncode))
	}
}

func Forbidden(rw http.ResponseWriter) {
	ErrorResponse(rw, http.StatusForbidden, "Forbidden")
}

func CustomResponse(rw http.ResponseWriter, code int, headers map[string]string, body []byte) {
	rw.WriteHeader(code)
	for k, v := range headers {
		rw.Header().Set(k, v)
	}
	rw.Write(body)
}
