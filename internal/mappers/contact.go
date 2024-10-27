package mappers

import (
	"github.com/andremartinsds/go_admin/internal/dto"
	"github.com/andremartinsds/go_admin/internal/entities"
	"github.com/andremartinsds/go_admin/internal/infra/db/models"
)

func ContactEntityToContactOutputDto(contact entities.Contact) dto.ContactsOutputDto {
	return dto.ContactsOutputDto{
		ID:      contact.ID,
		Name:    contact.Name,
		Phone:   contact.Phone,
		Email:   contact.Email,
		Message: contact.Message,
	}
}

func ContactModeltoEntity(c models.ContactModel) entities.Contact {
	return entities.Contact{
		ID:      c.ID.String(),
		Name:    c.Name,
		Phone:   c.Phone,
		Email:   c.Email,
		Message: c.Message,
	}
}
