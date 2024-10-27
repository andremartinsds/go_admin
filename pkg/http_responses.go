package pkg

import (
	"encoding/json"
	"github.com/andremartinsds/go_admin/internal/errs"
	"net/http"
)

type StandardError struct {
	W          http.ResponseWriter
	Message    string
	StatusCode int
}

func StandardErrorResponse(standardError StandardError) {
	standardError.W.WriteHeader(standardError.StatusCode)
	err := json.NewEncoder(standardError.W).Encode(errs.HttpResponse{ErrorCode: standardError.StatusCode,
		Message: standardError.Message})
	if err != nil {
		standardError.W.WriteHeader(standardError.StatusCode)
	}
	return
}

func DefaultHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-type", "application/json")
}
