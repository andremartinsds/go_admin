package mock

import "github.com/andremartinsds/go_admin/internal/dto"

func CreateEnderecoInputCreateDtoMock() dto.AddressInputCreateDTO {
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
