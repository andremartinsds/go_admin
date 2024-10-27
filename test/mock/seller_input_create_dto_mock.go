package mock

import (
	"github.com/andremartinsds/go_admin/internal/dto"
	"github.com/andremartinsds/go_admin/pkg"
)

func SellerInputCreateDtoMock() dto.SellerInputCreateDTO {
	return dto.SellerInputCreateDTO{
		AccountID:     pkg.NewUUID().String(),
		Nickname:      "Seller nome fantasia",
		CorporateName: "Seller Rz social",
		Document:      "21725051000150",
		Active:        true,
		Address: dto.AddressInputCreateDTO{
			ZipCode:        "32600-212",
			State:          "SP",
			City:           "São Paulo",
			Description:    "Avenida Governador Valadares",
			Number:         "123",
			Complement:     "Perto do BB",
			Neighborhood:   "Centro",
			ReferencePoint: "BB",
			Observation:    "Porta azul",
		},
	}
}
