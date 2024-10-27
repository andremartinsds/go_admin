package mappers

import (
	"github.com/andremartinsds/go_admin/internal/entities"
	"github.com/andremartinsds/go_admin/internal/infra/db/models"
)

func EnderecoModelToEntity(addressModel models.AddressModel) *entities.Address {
	return &entities.Address{
		ID:             addressModel.ID,
		ZipCode:        addressModel.ZipCode,
		State:          addressModel.State,
		City:           addressModel.City,
		Description:    addressModel.AddressDescription,
		Number:         addressModel.Number,
		Complement:     addressModel.Complement,
		Neighborhood:   addressModel.Neighborhood,
		ReferencePoint: addressModel.ReferencePoint,
		Observation:    addressModel.Observation,
	}
}
