package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/andremartinsds/go_admin/internal/auth"
	"github.com/andremartinsds/go_admin/internal/infra/repositories"
	"github.com/andremartinsds/go_admin/pkg"
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

	JwtAuth, _ := auth.JWT(r)
	var inputDTO struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&inputDTO)
	if err != nil {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	userEntity, err := a.repository.FindUserByUsername(inputDTO.Username, inputDTO.Password)
	if err != nil {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: err.Error(), StatusCode: http.StatusUnauthorized})
		return
	}
	if !userEntity.ValidatePassword(inputDTO.Password) {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "invalid username or password", StatusCode: http.StatusUnauthorized})
		return
	}

	userID := pkg.UUIDToString(userEntity.ID)
	_, tokenString, _ := JwtAuth.Encode(map[string]interface{}{
		"sub":       userID,
		"accountID": userEntity.AccountID,
		"sellerID":  userEntity.SellerID,
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
