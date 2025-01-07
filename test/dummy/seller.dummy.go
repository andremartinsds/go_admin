package dummy

import (
	"github.com/andremartinsds/go_admin/internal/dto"
	"github.com/andremartinsds/go_admin/pkg"
	"github.com/andremartinsds/go_admin/test/mock"
)

func CreateInputSellerDTODummy() *dto.SellerInputCreateDTO {
	return &dto.SellerInputCreateDTO{
		AccountID:     pkg.NewUUID().String(),
		Nickname:      "Name Company",
		CorporateName: "Razao Social",
		Document:      "11111111111",
		Active:        true,
		Address:       mock.CreateAddressInputDTOMock(),
	}
}
