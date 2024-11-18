package auth

import (
	"github.com/go-chi/jwtauth/v5"
	"net/http"
)

func JWT(r *http.Request) (*jwtauth.JWTAuth, int) {
	return r.Context().Value("TokenAuth").(*jwtauth.JWTAuth), r.Context().Value("JWTExpiresIn").(int)
}
