package dummy

import (
	"github.com/andremartinsds/go_admin/internal/entities"
	"github.com/andremartinsds/go_admin/pkg"
	"time"
)

func AccountEntitityDummy() *entities.Account {
	ativo := true
	return &entities.Account{
		ID:             pkg.NewUUID(),
		NickName:       "any valid nome fantasia",
		CommercialName: "any valid razao social",
		Document:       "any valid cnpj",
		Active:         &ativo,
		AccountType:    "PJ",
		CreatedAt:      time.Date(2024, time.July, 2, 12, 2, 1, 1, time.UTC),
		UpdatedAt:      time.Date(2024, time.July, 2, 12, 2, 1, 1, time.UTC),
		Address:        AddressDummy(),
	}
}
