package handlers

import (
	"encoding/json"
	"github.com/andremartinsds/go_admin/internal/dto"
	"github.com/andremartinsds/go_admin/internal/entities"
	"github.com/andremartinsds/go_admin/internal/infra/repositories"
	"github.com/andremartinsds/go_admin/internal/mappers"
	"github.com/andremartinsds/go_admin/pkg"
	"github.com/go-chi/chi/v5"
	"github.com/samber/lo"
	"net/http"
)

// AccountHandler handles HTTP requests related to user accounts.
type AccountHandler struct {
	repository repositories.AccountContract // Repository for account operations
}

// AccountHandlerInstancy creates a new instance of AccountHandler with the provided contract.
func AccountHandlerInstancy(contract repositories.AccountContract) *AccountHandler {
	return &AccountHandler{
		repository: contract,
	}
}

// CreateAccount handles the creation of a new account.
func (a *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	pkg.DefaultHeaders(w)

	var accountDto dto.AccountInputCreateDTO
	err := json.NewDecoder(r.Body).Decode(&accountDto)
	if err != nil {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}

	// Check if the account already exists
	exists, _ := a.repository.ExistsBy(map[string]string{"document": accountDto.Document})
	if exists {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "account already exists", StatusCode: http.StatusBadRequest})
		return
	}

	// Create a new account entity from the DTO
	account, err := entities.NewAccount(accountDto)
	if err != nil {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}

	// Save the account in the repository
	err = a.repository.Create(account)
	if err != nil {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(mappers.ToAccountOutputDTO(account)) // Respond with the created account
}

// SelectAccount retrieves an account by its ID.
func (a *AccountHandler) SelectAccount(w http.ResponseWriter, r *http.Request) {
	accountId := chi.URLParam(r, "accountID")
	if len(accountId) != pkg.SizeGuid {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "invalid account id", StatusCode: http.StatusBadRequest})
		return
	}
	pkg.DefaultHeaders(w)

	// Fetch the account from the repository
	account, err := a.repository.SelectOneById(accountId)
	if err != nil {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "account not found", StatusCode: http.StatusNotFound})
		return
	}

	accountOutput := mappers.ToAccountOutputDTO(account)
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(accountOutput)
	if err != nil {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "encode error", StatusCode: http.StatusInternalServerError})
	}
}

// UpdateAccount modifies an existing account.
func (a *AccountHandler) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	pkg.DefaultHeaders(w)
	accountId := chi.URLParam(r, "accountID")

	// Check if the account exists
	accountFound, err := a.repository.SelectOneById(accountId)
	if err != nil {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "account not found", StatusCode: http.StatusNotFound})
		return
	}

	var accountUpdateDto dto.AccountInputUpdateDTO
	err = json.NewDecoder(r.Body).Decode(&accountUpdateDto)
	if err != nil {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}

	// Preserve the existing address ID if it exists
	accountUpdateDto.Id = accountId
	if lo.IsNotEmpty(accountFound.Address.ID) {
		accountUpdateDto.Address.ID = accountFound.Address.ID.String()
	}

	account, err := entities.UpdateAccount(accountUpdateDto)
	account.CreatedAt = accountFound.CreatedAt
	account.Address.CreatedAt = accountFound.Address.CreatedAt
	if err != nil {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}

	err = a.repository.UpdateOne(account)
	if err != nil {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(account)
	if err != nil {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "accounts parsed error", StatusCode: http.StatusInternalServerError})
		return
	}
}

// List retrieves all accounts.
func (a *AccountHandler) List(w http.ResponseWriter, r *http.Request) {
	pkg.DefaultHeaders(w)

	// Fetch the list of accounts from the repository
	accounts, err := a.repository.List()
	if err != nil || len(*accounts) == 0 {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "accounts not found", StatusCode: http.StatusNotFound})
		return
	}

	w.WriteHeader(http.StatusOK)
	var accountOutputDto []dto.AccountOutputDto

	// Convert each account to the output DTO format
	for _, account := range *accounts {
		accountOutputDto = append(accountOutputDto, *mappers.ToAccountOutputDTO(&account))
	}

	err = json.NewEncoder(w).Encode(accountOutputDto) // Respond with the list of accounts
	if err != nil {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: err.Error(), StatusCode: http.StatusNotFound})
	}
}

func (a *AccountHandler) Delete(w http.ResponseWriter, r *http.Request) {
	accountID := chi.URLParam(r, "accountID")
	if accountID == "" {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "account id required", StatusCode: http.StatusBadRequest})
		return
	}
	accountEntity, err := a.repository.SelectOneById(accountID)
	if accountEntity == nil || err != nil {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "account not found", StatusCode: http.StatusNotFound})
		return
	}

	err = a.repository.DeleteById(accountEntity)
	if err != nil {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	w.WriteHeader(http.StatusOK)
}
