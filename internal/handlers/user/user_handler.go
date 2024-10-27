package handlers

import (
	"encoding/json"
	"github.com/andremartinsds/go_admin/internal/entities"
	"net/http"

	"github.com/andremartinsds/go_admin/internal/errs"
	"github.com/andremartinsds/go_admin/internal/mappers"
	"github.com/samber/lo"

	"github.com/go-chi/chi"

	"github.com/andremartinsds/go_admin/internal/dto"

	"github.com/andremartinsds/go_admin/internal/infra/repositories"
)

type UserHandler struct {
	repository repositories.UserContract
}

func UserHandlerInstance(contract repositories.UserContract) *UserHandler {
	return &UserHandler{
		repository: contract,
	}
}

func (a *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var userInputDto dto.UserInputCreateDTO
	err := json.NewDecoder(r.Body).Decode(&userInputDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(errs.HttpResponse{ErrorCode: http.StatusBadRequest, Message: err.Error()})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}
	err = a.repository.UserExists(map[string]string{"email": userInputDto.Email})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(errs.HttpResponse{ErrorCode: http.StatusBadRequest, Message: "user already exists"})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}
	user, err := entities.CreateUser(userInputDto)
	if err != nil {
		err = json.NewEncoder(w).Encode(errs.HttpResponse{ErrorCode: http.StatusBadRequest, Message: err.Error()})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	err = a.repository.Create(user)
	if err != nil {
		err = json.NewEncoder(w).Encode(errs.HttpResponse{ErrorCode: http.StatusBadRequest, Message: err.Error()})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(mappers.FromUserToUserOutputCreateDTO(user))
}

func (a *UserHandler) SelectUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	w.Header().Set("Content-type", "application/json")

	user, err := a.repository.SelectOneById(userId)
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

	userFound, err := a.repository.SelectOneById(userId)
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
