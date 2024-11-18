package middlewares

import (
	"context"
	"fmt"
	"github.com/andremartinsds/go_admin/internal/infra/db/connection"
	"github.com/andremartinsds/go_admin/internal/infra/repositories"
	"github.com/andremartinsds/go_admin/pkg"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
)

func AuthPermissions(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, claims, err := jwtauth.FromContext(r.Context())
		if r.URL.Path == "/login" {
			next.ServeHTTP(w, r)
			return
		}
		currentRequestUserId := token.Subject()
		//TODO: find the current claim for route

		//TODO: load user id on request
		userRepository := repositories.UserRepositoryInstancy(connection.DataSource)
		//TODO: find user with claims and role
		user, _ := userRepository.SelectOneById(currentRequestUserId)

		if pkg.IsEmptyUUID(user.ID) {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		//TODO: verify if role has the current claim, if it is true than return next middleware
		//TODO: verify if current claim is on the claim user than return next middleware

		//token, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {

		}
		fmt.Println(claims)
		//fmt.Println(userId)
		fmt.Println(token)
		// create new context from `r` request context, and assign key `"user"`
		// to value of `"123"`
		ctx := context.WithValue(r.Context(), "user", "123")

		// call the next handler in the chain, passing the response writer and
		// the updated request object with the new context value.
		//
		// note: context.Context values are nested, so any previously set
		// values will be accessible as well, and the new `"user"` key
		// will be accessible from this point forward.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
