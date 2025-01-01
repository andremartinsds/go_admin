package handlers

import (
	"encoding/json"
	"github.com/andremartinsds/go_admin/internal/auth"
	"github.com/andremartinsds/go_admin/internal/infra/repositories"
	"github.com/andremartinsds/go_admin/pkg"
	"net/http"
)

// LoginHandler handles HTTP requests related to user accounts.
type LoginHandler struct {
	repository repositories.UserContract
}

// LoginHandlerInstancy creates a new instance of LoginHandler with the provided contract.
func LoginHandlerInstance(userRepository repositories.UserContract) *LoginHandler {
	return &LoginHandler{
		repository: userRepository,
	}
}

// Login handles.
func (a *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	pkg.DefaultHeaders(w)

	accountID := r.Header.Get("accountID")
	sellerID := r.Header.Get("sellerID")

	if !pkg.ValidUUID(accountID) || (sellerID != "" && !pkg.ValidUUID(sellerID)) {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "unauthorized", StatusCode: http.StatusUnauthorized})
		return
	}

	JwtAuth, _ := auth.JWT(r)
	var inputDTO struct {
		username string
		password string
	}
	err := json.NewDecoder(r.Body).Decode(&inputDTO)
	if err != nil {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	userEntity, err := a.repository.Login(inputDTO.username, inputDTO.password)
	if err != nil {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	_, tokenString, _ := JwtAuth.Encode(map[string]interface{}{
		"sub": userEntity.ID,
	})

	accessToken := struct {
		AccessToken string `json:"accessToken"`
	}{
		AccessToken: tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}
