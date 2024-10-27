package dummy

import "github.com/andremartinsds/go_admin/internal/dto"

func CreateInputAccountDTODummy() *dto.AccountInputCreateDTO {
	return &dto.AccountInputCreateDTO{
		Name:           "Name Company",
		CommercialName: "Razao Social",
		Document:       "11111111111",
		Active:         true,
		AccountType:    "PJ",
		Address:        *CreateAddressDTODummy(),
	}
}

func UpdateInputAccountDTODummy() *dto.AccountInputUpdateDTO {
	return &dto.AccountInputUpdateDTO{
		Name:           "Name Company",
		CommercialName: "Razao Social",
		Document:       "11111111111",
		Active:         true,
		AccountType:    "PJ",
		Address:        *UpdateAddressDTODummy(),
	}
}
