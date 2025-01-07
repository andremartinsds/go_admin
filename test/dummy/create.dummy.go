package dummy

import (
	"github.com/andremartinsds/go_admin/internal/dto"
	"github.com/andremartinsds/go_admin/test/mock"
)

func CreateInputAccountDTODummy() *dto.AccountInputCreateDTO {
	return &dto.AccountInputCreateDTO{
		Name:           "Name Company",
		CommercialName: "Razao Social",
		Document:       "11111111111",
		Active:         true,
		AccountType:    "PJ",
		Address:        mock.CreateAddressInputDTOMock(),
	}
}

func UpdateInputAccountDTODummy() *dto.AccountInputUpdateDTO {
	return &dto.AccountInputUpdateDTO{
		Name:           "Name Company",
		CommercialName: "Razao Social",
		Document:       "11111111111",
		Active:         true,
		AccountType:    "PJ",
		Address:        mock.UpdateAddressInputDTOMock(),
	}
}
