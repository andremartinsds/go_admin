package entities

import (
	"errors"
	"github.com/andremartinsds/go_admin/internal/dto"
	"strings"
)

// Contact represents the contact information structure.
type Contact struct {
	// ID is the unique identifier for the contact.
	ID string

	// Name is the name of the contact person.
	Name string

	// Phone is the contact's phone number.
	Phone string

	// Email is the contact's email address.
	Email string

	// Message is the message sent by the contact.
	Message string
}

// validateToCreate checks if the contact information is valid for creation.
func (c *Contact) validateToCreate() error {
	errs := []string{}

	if c.Name == "" {
		errs = append(errs, "field: name is required")
	}

	if c.Phone == "" {
		errs = append(errs, "field: phone is required")
	}

	if c.Email == "" {
		errs = append(errs, "field: email is required")
	}

	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "\n"))
	}

	return nil
}

// NewContact creates a new contact from the provided input DTO.
func NewContact(contactInputDTO dto.ContactInputCreateDTO) (*Contact, error) {
	contact := &Contact{
		Name:    contactInputDTO.Name,
		Phone:   contactInputDTO.Phone,
		Email:   contactInputDTO.Email,
		Message: contactInputDTO.Message,
	}
	err := contact.validateToCreate()
	if err != nil {
		return nil, err
	}

	return contact, nil
}
