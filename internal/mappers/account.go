package mappers

import (
	"github.com/andremartinsds/go_admin/internal/dto"
	"github.com/andremartinsds/go_admin/internal/entities"
	"github.com/andremartinsds/go_admin/internal/infra/db/models"
	"github.com/andremartinsds/go_admin/pkg"
	"github.com/samber/lo"
)

func ToAccountOutputDTO(account *entities.Account) *dto.AccountOutputDto {
	var enderecoOutput *dto.AddressOutputDTO
	if lo.IsNotEmpty(account.Address) && !lo.Contains([]string{"00000000-0000-0000-0000-000000000000"}, account.Address.ID.String()) {
		enderecoOutput = &dto.AddressOutputDTO{
			ID:             account.Address.ID.String(),
			ZipCode:        account.Address.ZipCode,
			State:          account.Address.State,
			City:           account.Address.City,
			Description:    account.Address.Description,
			Number:         account.Address.Number,
			Complement:     account.Address.Complement,
			Neighborhood:   account.Address.Neighborhood,
			ReferencePoint: account.Address.ReferencePoint,
			Observation:    account.Address.Observation,
		}
	}
	return &dto.AccountOutputDto{
		Id:          account.ID.String(),
		Nickname:    account.NickName,
		Document:    account.Document,
		Active:      *account.Active,
		AccountType: account.AccountType,
		Address:     enderecoOutput,
	}
}

func ToAccountEntity(account models.AccountModel) *entities.Account {
	var address *entities.Address
	if !pkg.IsEmptyUUID(account.Address.ID) {
		address = &entities.Address{
			ID:             account.Address.ID,
			ZipCode:        account.Address.ZipCode,
			State:          account.Address.State,
			City:           account.Address.City,
			Description:    account.Address.AddressDescription,
			Number:         account.Address.Number,
			Complement:     account.Address.Complement,
			Neighborhood:   account.Address.Neighborhood,
			ReferencePoint: account.Address.ReferencePoint,
			Observation:    account.Address.Observation,
		}
	}

	return &entities.Account{
		ID:             account.ID,
		NickName:       account.NickName,
		CommercialName: account.CommercialName,
		Document:       account.Document,
		Active:         &account.Active,
		AccountType:    account.AccountType,
		CreatedAt:      account.CreatedAt,
		UpdatedAt:      account.UpdatedAt,
		Address:        address,
	}
}
