package repositories

import (
	"errors"
	"github.com/andremartinsds/go_admin/internal/entities"
	"github.com/andremartinsds/go_admin/internal/infra/db/models"
	"github.com/andremartinsds/go_admin/internal/mappers"
	"github.com/andremartinsds/go_admin/pkg"
	"gorm.io/gorm"
)

type ContractContract interface {
	Create(account *entities.Contact) error
	List() (*[]entities.Contact, error)
}

type ContactRepository struct {
	db *gorm.DB
}

func ContactRepositoryInstancy(connection *gorm.DB) *ContactRepository {
	return &ContactRepository{db: connection}
}

func (c *ContactRepository) Create(contact *entities.Contact) error {
	if err := c.db.Create(
		&models.ContactModel{
			ID:      pkg.NewUUID(),
			Name:    contact.Name,
			Phone:   contact.Phone,
			Email:   contact.Email,
			Message: contact.Message,
		}).Error; err != nil {
		return err
	}
	return nil
}

func (a *ContactRepository) List() (*[]entities.Contact, error) {

	var contacts []models.ContactModel
	err := a.db.Find(&contacts).Error
	if err != nil {
		return nil, err
	}
	if len(contacts) > 0 {
		return nil, errors.New("we do not search any contact")
	}

	var c []entities.Contact
	for _, v := range contacts {
		contactMapper := mappers.ContactModeltoEntity(v)
		c = append(c, contactMapper)
	}

	return &c, nil
}
