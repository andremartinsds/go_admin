package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/andremartinsds/go_admin/internal/entities"
	"github.com/andremartinsds/go_admin/pkg"

	"github.com/andremartinsds/go_admin/internal/errs"
	"github.com/andremartinsds/go_admin/internal/mappers"
	"github.com/samber/lo"

	"github.com/go-chi/chi/v5"

	"github.com/andremartinsds/go_admin/internal/dto"

	"github.com/andremartinsds/go_admin/internal/infra/repositories"
)

type UserHandler struct {
	userRepository    repositories.UserContract
	accountRepository repositories.AccountContract
	sellerRepository  repositories.SellerContract
}

func UserHandlerInstance(userRepository repositories.UserContract, accountRepository repositories.AccountContract, sellerRepository repositories.SellerContract) *UserHandler {
	return &UserHandler{
		userRepository:    userRepository,
		accountRepository: accountRepository,
		sellerRepository:  sellerRepository,
	}
}

func (a *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var userInputDto dto.UserInputCreateDTO
	err := json.NewDecoder(r.Body).Decode(&userInputDto)
	if err != nil {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "decoder with error", StatusCode: http.StatusBadRequest})
		return
	}

	userInputDto.AccountID = r.Header.Get("accountID")
	userInputDto.SellerID = r.Header.Get("sellerID")

	if !pkg.ValidUUID(userInputDto.AccountID) || !pkg.ValidUUID(userInputDto.SellerID) {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "the account or seller invalid", StatusCode: http.StatusBadRequest})
		return
	}

	accountEntity, _ := a.accountRepository.SelectOneById(userInputDto.AccountID)
	if lo.IsEmpty(accountEntity.Document) {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "account does not found", StatusCode: http.StatusNotFound})
		return
	}
	sellerEntity, _ := a.sellerRepository.SelectOneById(userInputDto.SellerID)
	if lo.IsEmpty(sellerEntity.Document) {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "seller does not found", StatusCode: http.StatusNotFound})
		return
	}

	if !sellerEntity.IsAccountIDEqual(userInputDto.AccountID) {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "account or seller does not valid", StatusCode: http.StatusNotFound})
		return
	}
	err = a.userRepository.UserExists(map[string]string{"email": userInputDto.Email})
	if err != nil {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "user already exists", StatusCode: http.StatusNotFound})
		return
	}
	user, err := entities.CreateUser(userInputDto)
	if err != nil {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}

	err = a.userRepository.Create(user)
	if err != nil {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(mappers.FromUserToUserOutputCreateDTO(user))
}

func (a *UserHandler) SelectUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	w.Header().Set("Content-type", "application/json")

	user, err := a.userRepository.SelectOneById(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(w).Encode(errs.HttpResponse{ErrorCode: http.StatusBadRequest, Message: "user not found"})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mappers.FromUserToUserOutputCreateDTO(user))
}

func (a *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	userId := chi.URLParam(r, "userId")

	userFound, err := a.userRepository.SelectOneById(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var userUpdateDto dto.UserInputUpdateDTO
	err = json.NewDecoder(r.Body).Decode(&userUpdateDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userUpdateDto.ID = userId
	if lo.IsNotEmpty(userFound.Address.ID) {
		userUpdateDto.Address.ID = userFound.Address.ID.String()
	}
}
