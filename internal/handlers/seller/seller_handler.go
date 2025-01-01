package seller

import (
	"encoding/json"
	"github.com/andremartinsds/go_admin/internal/entities"
	"github.com/andremartinsds/go_admin/internal/mappers"
	"github.com/andremartinsds/go_admin/pkg"
	"net/http"

	"github.com/andremartinsds/go_admin/internal/dto"
	"github.com/andremartinsds/go_admin/internal/infra/repositories"
)

// Handler manages operations related to sellers and their accounts.
type Handler struct {
	sellerRepository  repositories.SellerContract
	accountRepository repositories.AccountContract
}

// Instance creates a new instance of Handler with the required repositories.
func Instance(sellerContract repositories.SellerContract, accountContract repositories.AccountContract) *Handler {
	return &Handler{
		sellerRepository:  sellerContract,
		accountRepository: accountContract,
	}
}

// CreateSeller handles the creation of a new seller.
// Example request payload:
//
//	{
//	  "name": "Example Seller",
//	  "accountID": "12345",
//	  "document": "12345678901234"
//	}
func (sellerHandler *Handler) CreateSeller(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Decode request body into Seller creation DTO.
	var sellerDto dto.SellerInputCreateDTO
	err := json.NewDecoder(r.Body).Decode(&sellerDto)
	if err != nil {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "Invalid input data", StatusCode: http.StatusBadRequest})
		return
	}

	sellerDto.AccountID = r.Header.Get("accountID")

	if !pkg.ValidUUID(sellerDto.AccountID) {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "the current account does not valid", StatusCode: http.StatusBadRequest})
		return
	}

	// Check if account exists.
	accountExists, _ := sellerHandler.accountRepository.ExistsBy(map[string]string{"id": sellerDto.AccountID})
	if !accountExists {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "The account does not exist", StatusCode: http.StatusBadRequest})
		return
	}

	// Create a new seller entity from DTO.
	seller, err := entities.NewSeller(sellerDto)
	if err != nil {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}

	// Check if seller already exists by document.
	sellerExists, _ := sellerHandler.sellerRepository.Exists(map[string]string{"document": seller.Document})
	if sellerExists {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "The seller already exists", StatusCode: http.StatusBadRequest})
		return
	}

	// Save the new seller to the repository.
	err = sellerHandler.sellerRepository.Create(seller)
	if err != nil {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "Error saving seller", StatusCode: http.StatusInternalServerError})
		return
	}

	w.WriteHeader(http.StatusCreated)
	sellerOutputDTO := mappers.SellerEntityToSellerOutputDTO(*seller)
	json.NewEncoder(w).Encode(sellerOutputDTO)
}

// UpdateSeller handles the update of an existing seller's information.
func (sellerHandler *Handler) UpdateSeller(w http.ResponseWriter, r *http.Request) {
	pkg.DefaultHeaders(w)
	sellerID := r.Header.Get("sellerID")
	accountID := r.Header.Get("accountID")

	if !pkg.ValidUUID(sellerID) || !pkg.ValidUUID(accountID) {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "Invalid input data", StatusCode: http.StatusBadRequest})
		return
	}

	// Decode request body into Seller update DTO.
	var sellerInputUpdateDTO dto.SellerInputUpdateDTO
	err := json.NewDecoder(r.Body).Decode(&sellerInputUpdateDTO)
	sellerInputUpdateDTO.ID = sellerID
	if err != nil {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "Invalid input data", StatusCode: http.StatusBadRequest})
		return
	}

	// Check if account exists.
	accountExists, _ := sellerHandler.accountRepository.ExistsBy(map[string]string{"id": accountID})
	if !accountExists {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "The account does not exist", StatusCode: http.StatusBadRequest})
		return
	}

	// Check if seller exists by ID.
	sellerFound, err := sellerHandler.sellerRepository.Select(map[string]string{"id": sellerID})
	if err != nil || sellerFound == nil {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "The seller does not exist", StatusCode: http.StatusBadRequest})
		return
	}

	if !sellerFound.IsAccountIDEqual(accountID) {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "The account does not exist", StatusCode: http.StatusBadRequest})
		return
	}

	// Update seller entity.
	sellerInputUpdateDTO.CreatedAt = sellerFound.CreatedAt
	sellerInputUpdateDTO.Address.CreatedAt = sellerFound.Address.CreatedAt
	sellerInputUpdateDTO.Address.ID = sellerFound.Address.ID.String()
	seller, err := entities.SellerUpdate(sellerInputUpdateDTO)
	if err != nil {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "try again, something wrong", StatusCode: http.StatusBadRequest})
		return
	}

	// Update the seller in the repository.
	err = sellerHandler.sellerRepository.Update(seller)
	if err != nil {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "Error updating seller", StatusCode: http.StatusInternalServerError})
		return
	}

	sellerOutputDTO := mappers.SellerEntityToSellerOutputDTO(*seller)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sellerOutputDTO)
}

// SelectSeller handles the retrieval of a seller by their ID.
func (sellerHandler *Handler) SelectSeller(w http.ResponseWriter, r *http.Request) {
	sellerID := r.Header.Get("sellerID")
	accountID := r.Header.Get("accountID")

	if !pkg.ValidUUID(sellerID) || !pkg.ValidUUID(accountID) {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "invalid UUID", StatusCode: http.StatusBadRequest})
		return
	}

	// Retrieve the seller entity by ID from the repository.
	sellerEntity, err := sellerHandler.sellerRepository.Select(map[string]string{"id": sellerID})
	if err != nil {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "seller not found", StatusCode: http.StatusNotFound})
		json.NewEncoder(w).Encode("")
		return
	}

	accountIDEqual := sellerEntity.IsAccountIDEqual(accountID)

	if !accountIDEqual {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "invalid account ID for the current seller", StatusCode: http.StatusBadRequest})
		return
	}

	sellerOutput := mappers.SellerEntityToSellerOutputDTO(*sellerEntity)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sellerOutput)
}

// ListSeller handles the listing of all sellers.
func (sellerHandler *Handler) ListSeller(w http.ResponseWriter, r *http.Request) {
	accountID := r.Header.Get("accountID")
	if !pkg.ValidUUID(accountID) {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "the accountID is required", StatusCode: http.StatusBadRequest})
		return
	}

	sellers, err := sellerHandler.sellerRepository.List(accountID)
	if err != nil || len(sellers) == 0 {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "does not exist sellers to this account", StatusCode: http.StatusNotImplemented})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sellers)
}

// Inactive the seller with seller ID.
func (sellerHandler *Handler) Inactive(w http.ResponseWriter, r *http.Request) {
	sellerID := r.Header.Get("sellerID")
	accountID := r.Header.Get("accountID")
	if !pkg.ValidUUID(accountID) || !pkg.ValidUUID(sellerID) {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "invalid UUID", StatusCode: http.StatusBadRequest})
		return
	}

	seller, err := sellerHandler.sellerRepository.Select(map[string]string{"id": sellerID})
	if err != nil || seller == nil {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "seller not found", StatusCode: http.StatusBadRequest})
		return
	}
	if !seller.IsAccountIDEqual(accountID) {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: "invalid seller or account", StatusCode: http.StatusBadRequest})
	}

	err = sellerHandler.sellerRepository.DeleteById(*seller)
	if err != nil {
		pkg.ErrorResponse(pkg.InternalError{ResponseWriter: w, Message: err.Error(), StatusCode: http.StatusInternalServerError})
		return
	}
	w.WriteHeader(http.StatusOK)
}
