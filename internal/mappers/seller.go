package mappers

import (
	"github.com/andremartinsds/go_admin/internal/dto"
	"github.com/andremartinsds/go_admin/internal/entities"
	"github.com/andremartinsds/go_admin/internal/infra/db/models"
	"github.com/andremartinsds/go_admin/pkg"
)

func SellerEntityToSellerModel(s *entities.Seller) *models.SellerModel {
	var accountModel models.AccountModel
	if !pkg.IsEmptyUUID(s.Account.ID) {

		var enderecoAccountModel models.AddressModel
		if !pkg.IsEmptyUUID(s.Account.Address.ID) {
			enderecoAccountModel = models.AddressModel{
				ID:                 s.Account.Address.ID,
				ZipCode:            s.Account.Address.ZipCode,
				State:              s.Account.Address.State,
				City:               s.Account.Address.City,
				AddressDescription: s.Account.Address.Description,
				Number:             s.Account.Address.Number,
				Complement:         s.Account.Address.Complement,
				Neighborhood:       s.Account.Address.Neighborhood,
				ReferencePoint:     s.Account.Address.ReferencePoint,
				Observation:        s.Account.Address.Observation,
				CreatedAt:          s.Account.CreatedAt,
				UpdatedAt:          s.Account.UpdatedAt,
			}
		}

		accountModel = models.AccountModel{
			ID:             s.Account.ID,
			NickName:       s.Account.NickName,
			CommercialName: s.Account.CommercialName,
			Document:       s.Account.Document,
			Active:         *s.Account.Active,
			AccountType:    s.Account.AccountType,
			AddressID:      s.Account.Address.ID,
			Address:        enderecoAccountModel,
			CreatedAt:      s.Account.CreatedAt,
			UpdatedAt:      s.Account.UpdatedAt,
		}
	}
	var enderecoSellerModel models.AddressModel
	if !pkg.IsEmptyUUID(s.Address.ID) {
		enderecoSellerModel = models.AddressModel{
			ID:                 s.Address.ID,
			ZipCode:            s.Address.ZipCode,
			State:              s.Address.State,
			City:               s.Address.City,
			AddressDescription: s.Address.Description,
			Number:             s.Address.Number,
			Complement:         s.Address.Complement,
			Neighborhood:       s.Address.Neighborhood,
			ReferencePoint:     s.Address.ReferencePoint,
			Observation:        s.Address.Observation,
			CreatedAt:          s.CreatedAt,
			UpdatedAt:          s.UpdatedAt,
		}
	}
	return &models.SellerModel{
		ID:            s.ID,
		NickName:      s.NickName,
		CorporateName: s.LegalName,
		Document:      s.Document,
		Active:        *s.Active,
		AccountID:     s.AccountID,
		Account:       accountModel,
		AddressID:     s.Address.ID,
		Address:       enderecoSellerModel,
		CreatedAt:     s.CreatedAt,
		UpdatedAt:     s.UpdatedAt,
	}
}

func SellerEntityToSellerOutputDTO(seller entities.Seller) dto.SellerOutputDto {
	var address dto.AddressOutputDTO
	if !pkg.IsEmptyUUID(seller.Address.ID) {
		address = dto.AddressOutputDTO{
			ID:             seller.Address.ID.String(),
			ZipCode:        seller.Address.ZipCode,
			State:          seller.Address.State,
			City:           seller.Address.City,
			Description:    seller.Address.Description,
			Number:         seller.Address.Number,
			Complement:     seller.Address.Complement,
			Neighborhood:   seller.Address.Neighborhood,
			ReferencePoint: seller.Address.ReferencePoint,
			Observation:    seller.Address.Observation,
		}
	}

	return dto.SellerOutputDto{
		ID:       pkg.StrID(seller.ID),
		Nickname: seller.NickName,
		Document: seller.Document,
		Address:  &address,
	}
}

func SellerModelToEntity(s *models.SellerModel) *entities.Seller {
	var enderecoSeller entities.Address
	if !pkg.IsEmptyUUID(s.Address.ID) {
		enderecoSeller = entities.Address{
			ID:             s.Address.ID,
			ZipCode:        s.Address.ZipCode,
			Description:    s.Address.AddressDescription,
			State:          s.Address.State,
			City:           s.Address.City,
			Number:         s.Address.Number,
			Complement:     s.Address.Complement,
			Neighborhood:   s.Address.Neighborhood,
			ReferencePoint: s.Address.ReferencePoint,
			Observation:    s.Address.Observation,
		}
	}
	var account entities.Account
	if !pkg.IsEmptyUUID(s.AccountID) {
		var address entities.Address
		if !pkg.IsEmptyUUID(s.Account.Address.ID) {
			address = entities.Address{
				ID:             s.Account.Address.ID,
				ZipCode:        s.Account.Address.ZipCode,
				Description:    s.Account.Address.AddressDescription,
				State:          s.Account.Address.State,
				City:           s.Account.Address.City,
				Number:         s.Account.Address.Number,
				Complement:     s.Account.Address.Complement,
				Neighborhood:   s.Account.Address.Neighborhood,
				ReferencePoint: s.Account.Address.ReferencePoint,
				Observation:    s.Account.Address.Observation,
			}
		}

		account = entities.Account{
			ID:             s.AccountID,
			NickName:       s.Account.NickName,
			CommercialName: s.Account.CommercialName,
			Document:       s.Account.Document,
			Active:         &s.Account.Active,
			AccountType:    s.Account.AccountType,
			CreatedAt:      s.Account.CreatedAt,
			UpdatedAt:      s.Account.UpdatedAt,
			Address:        &address,
		}
	}

	return &entities.Seller{
		ID:        s.ID,
		AccountID: s.AccountID,
		Account:   account,
		NickName:  s.NickName,
		LegalName: s.CorporateName,
		Document:  s.Document,
		Active:    &s.Active,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
		Address:   &enderecoSeller,
	}
}
