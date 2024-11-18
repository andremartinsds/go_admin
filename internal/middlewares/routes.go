package middlewares

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func ParamID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if len(id) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("forbidden pet: %s!\n", id)))
			return
		}
		next.ServeHTTP(w, r)
	})
}
