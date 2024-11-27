package seller

import (
	"encoding/json"
	"github.com/andremartinsds/go_admin/internal/entities"
	"github.com/andremartinsds/go_admin/internal/mappers"
	"github.com/andremartinsds/go_admin/pkg"
	"github.com/go-chi/chi/v5"
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
//	  "cnpj": "12345678901234"
//	}
func (sellerHandler *Handler) CreateSeller(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Decode request body into Seller creation DTO.
	var sellerDto dto.SellerInputCreateDTO
	err := json.NewDecoder(r.Body).Decode(&sellerDto)
	if err != nil {
		pkg.StandardErrorResponse(pkg.StandardError{W: w, Message: "Invalid input data", StatusCode: http.StatusBadRequest})
		return
	}

	// Check if account exists.
	accountExists, _ := sellerHandler.accountRepository.ExistsBy(map[string]string{"id": sellerDto.AccountID})
	if !accountExists {
		pkg.StandardErrorResponse(pkg.StandardError{W: w, Message: "The account does not exist", StatusCode: http.StatusBadRequest})
		return
	}

	// Create a new seller entity from DTO.
	seller, err := entities.NewSeller(sellerDto)
	if err != nil {
		pkg.StandardErrorResponse(pkg.StandardError{W: w, Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}

	// Check if seller already exists by CNPJ.
	sellerExists, _ := sellerHandler.sellerRepository.Exists(map[string]string{"cnpj": seller.Document})
	if sellerExists {
		pkg.StandardErrorResponse(pkg.StandardError{W: w, Message: "The seller already exists", StatusCode: http.StatusBadRequest})
		return
	}

	// Save the new seller to the repository.
	err = sellerHandler.sellerRepository.Create(seller)
	if err != nil {
		pkg.StandardErrorResponse(pkg.StandardError{W: w, Message: "Error saving seller", StatusCode: http.StatusInternalServerError})
		return
	}

	w.WriteHeader(http.StatusCreated)
	sellerOutputDTO := mappers.SellerEntityToSellerOutputDTO(*seller)
	json.NewEncoder(w).Encode(sellerOutputDTO)
}

// UpdateSeller handles the update of an existing seller's information.
func (sellerHandler *Handler) UpdateSeller(w http.ResponseWriter, r *http.Request) {
	pkg.DefaultHeaders(w)
	sellerID := chi.URLParam(r, "sellerID")

	// Decode request body into Seller update DTO.
	var sellerInputUpdateDTO dto.SellerInputUpdateDTO
	err := json.NewDecoder(r.Body).Decode(&sellerInputUpdateDTO)
	sellerInputUpdateDTO.ID = sellerID
	if err != nil {
		pkg.StandardErrorResponse(pkg.StandardError{W: w, Message: "Invalid input data", StatusCode: http.StatusBadRequest})
		return
	}

	// Check if account exists.
	accountExists, _ := sellerHandler.accountRepository.ExistsBy(map[string]string{"id": sellerInputUpdateDTO.AccountID})
	if !accountExists {
		pkg.StandardErrorResponse(pkg.StandardError{W: w, Message: "The account does not exist", StatusCode: http.StatusBadRequest})
		return
	}

	// Check if seller exists by ID.
	sellerFound, err := sellerHandler.sellerRepository.Select(map[string]string{"id": sellerID})
	if err != nil || sellerFound == nil {
		pkg.StandardErrorResponse(pkg.StandardError{W: w, Message: "The seller does not exist", StatusCode: http.StatusBadRequest})
		return
	}

	// Update seller entity.
	sellerInputUpdateDTO.CreatedAt = sellerFound.CreatedAt
	sellerInputUpdateDTO.Address.CreatedAt = sellerFound.Address.CreatedAt
	seller, err := entities.SellerUpdate(sellerInputUpdateDTO)
	if err != nil {
		pkg.StandardErrorResponse(pkg.StandardError{W: w, Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}

	// Update the seller in the repository.
	err = sellerHandler.sellerRepository.Update(seller)
	if err != nil {
		pkg.StandardErrorResponse(pkg.StandardError{W: w, Message: "Error updating seller", StatusCode: http.StatusInternalServerError})
		return
	}

	sellerOutputDTO := mappers.SellerEntityToSellerOutputDTO(*seller)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sellerOutputDTO)
}

// SelectSeller handles the retrieval of a seller by their ID.
func (sellerHandler *Handler) SelectSeller(w http.ResponseWriter, r *http.Request) {
	sellerID := chi.URLParam(r, "sellerID")

	// Retrieve the seller entity by ID from the repository.
	sellerEntity, err := sellerHandler.sellerRepository.Select(map[string]string{"id": sellerID})
	if err != nil {
		pkg.StandardErrorResponse(pkg.StandardError{W: w, Message: "Seller not found", StatusCode: http.StatusNotFound})
		return
	}

	sellerOutput := mappers.SellerEntityToSellerOutputDTO(*sellerEntity)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sellerOutput)
}

// ListSeller handles the listing of all sellers.
func (sellerHandler *Handler) ListSeller(w http.ResponseWriter, r *http.Request) {
	accountID := chi.URLParam(r, "accountID")
	if accountID == "" {
		pkg.StandardErrorResponse(pkg.StandardError{W: w, Message: "the accountID is required", StatusCode: http.StatusBadRequest})
		return
	}

	sellers, err := sellerHandler.sellerRepository.List(accountID)
	if err != nil || len(sellers) == 0 {
		pkg.StandardErrorResponse(pkg.StandardError{W: w, Message: "does not exist sellers to this account", StatusCode: http.StatusNotImplemented})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sellers)
}

// DesactiveSeller handles the deactivation of a seller by their ID.
func (s *Handler) DesactiveSeller(w http.ResponseWriter, r *http.Request) {

	sellerID := chi.URLParam(r, "sellerID")
	if sellerID == "" {
		pkg.StandardErrorResponse(pkg.StandardError{W: w, Message: "the sellerID is required", StatusCode: http.StatusBadRequest})
		return
	}
	accountID := chi.URLParam(r, "accountID")
	if accountID == "" {
		pkg.StandardErrorResponse(pkg.StandardError{W: w, Message: "the accountID is required", StatusCode: http.StatusBadRequest})
		return
	}
	seller, err := s.sellerRepository.Select(map[string]string{"id": sellerID})
	if err != nil || seller == nil {
		pkg.StandardErrorResponse(pkg.StandardError{W: w, Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}

	err = s.sellerRepository.DeleteById(sellerID)
	if err != nil {
		pkg.StandardErrorResponse(pkg.StandardError{W: w, Message: err.Error(), StatusCode: http.StatusInternalServerError})
		return
	}
	w.WriteHeader(http.StatusOK)
}
