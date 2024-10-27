package handlers

import (
	"encoding/json"
	"github.com/andremartinsds/go_admin/internal/entities"
	"net/http"

	"github.com/andremartinsds/go_admin/internal/errs"
	"github.com/andremartinsds/go_admin/internal/mappers"

	"github.com/andremartinsds/go_admin/internal/dto"

	"github.com/andremartinsds/go_admin/internal/infra/repositories"
)

// ContactHandler handles operations related to contacts.
type ContactHandler struct {
	repository repositories.ContractContract
}

// ContactHandlerInstancy creates an instance of ContactHandler.
func ContactHandlerInstancy(contract repositories.ContractContract) *ContactHandler {
	return &ContactHandler{
		repository: contract,
	}
}

// CreateContact handles the creation of a new contact.
// Example request payload:
//
//	{
//	  "name": "John Doe",
//	  "email": "john.doe@example.com",
//	  "phone": "123-456-7890"
//	}
func (c *ContactHandler) CreateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Decode request body into contact creation DTO.
	var contactDto dto.ContactInputCreateDTO
	err := json.NewDecoder(r.Body).Decode(&contactDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errs.HttpResponse{ErrorCode: http.StatusBadRequest, Message: "Invalid input data"})
		return
	}

	// Create contact entity from DTO.
	contact, err := entities.NewContact(contactDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errs.HttpResponse{ErrorCode: http.StatusBadRequest, Message: err.Error()})
		return
	}

	// Insert the contact into the repository.
	err = c.repository.Create(contact)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs.HttpResponse{ErrorCode: http.StatusInternalServerError, Message: "Error saving contact"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(mappers.ContactEntityToContactOutputDto(*contact))
}

// List handles listing all existing contacts.
func (c *ContactHandler) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Retrieve the list of contacts from the repository.
	contactsFound, err := c.repository.List()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errs.HttpResponse{ErrorCode: http.StatusNotFound, Message: "Contacts not found"})
		return
	}

	w.WriteHeader(http.StatusOK)
	var contactsOutputDTO []dto.ContactsOutputDto

	// Map each contact entity to the output DTO.
	for _, contact := range *contactsFound {
		c := mappers.ContactEntityToContactOutputDto(contact)
		contactsOutputDTO = append(contactsOutputDTO, c)
	}

	// Encode the list of output DTOs as JSON.
	err = json.NewEncoder(w).Encode(contactsOutputDTO)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs.HttpResponse{ErrorCode: http.StatusInternalServerError, Message: "Error encoding response"})
	}
}
