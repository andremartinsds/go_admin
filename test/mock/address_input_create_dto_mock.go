package mock

import (
	"github.com/andremartinsds/go_admin/internal/dto"
	"github.com/andremartinsds/go_admin/pkg"
)

func CreateAddressInputDTOMock() dto.AddressInputCreateDTO {
	return dto.AddressInputCreateDTO{
		ZipCode:        "32600-212",
		State:          "Sp",
		City:           "Sp",
		Description:    "Avenida Governador Valadares",
		Number:         "123",
		Complement:     "Perto do BB",
		Neighborhood:   "Centro",
		ReferencePoint: "BB",
		Observation:    "Porta azul",
	}
}

func UpdateAddressInputDTOMock() dto.AddressInputUpdateDTO {
	return dto.AddressInputUpdateDTO{
		ID:             pkg.UUIDToString(pkg.NewUUID()),
		ZipCode:        "32600-212",
		State:          "Sp",
		City:           "Sp",
		Description:    "Avenida Governador Valadares",
		Number:         "123",
		Complement:     "Perto do BB",
		Neighborhood:   "Centro",
		ReferencePoint: "BB",
		Observation:    "Porta azul",
	}
}
