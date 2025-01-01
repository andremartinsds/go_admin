package pkg

import (
	"encoding/json"
	"github.com/andremartinsds/go_admin/internal/errs"
	"net/http"
)

type InternalError struct {
	ResponseWriter http.ResponseWriter
	Message        string
	StatusCode     int
}

func ErrorResponse(standardError InternalError) {
	standardError.ResponseWriter.WriteHeader(standardError.StatusCode)
	err := json.NewEncoder(standardError.ResponseWriter).Encode(
		errs.HttpResponse{ErrorCode: standardError.StatusCode, Message: standardError.Message})
	if err != nil {
		standardError.ResponseWriter.WriteHeader(standardError.StatusCode)
	}
	return
}

func DefaultHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-type", "application/json")
}
