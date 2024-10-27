package repositories

import (
	"errors"
	"github.com/andremartinsds/go_admin/internal/entities"
	"github.com/andremartinsds/go_admin/internal/infra/db/models"
	"github.com/andremartinsds/go_admin/internal/mappers"
	"github.com/andremartinsds/go_admin/pkg"
	"gorm.io/gorm"
	"time"
)

type SellerContract interface {
	Create(seller *entities.Seller) error
	Exists(param map[string]string) (bool, error)
	Update(seller *entities.Seller) error
	Select(param map[string]string) (*entities.Seller, error)
}

type SellerRepository struct {
	db *gorm.DB
}

func SellerRepositoryInstancy(connection *gorm.DB) *SellerRepository {
	return &SellerRepository{db: connection}
}

func (s *SellerRepository) Select(param map[string]string) (*entities.Seller, error) {
	key, value, _ := pkg.GetKeyValueFromMap(param)
	var sellerModel models.SellerModel
	if err := s.db.Preload("Endereco").Preload("Account").Preload("Account.Endereco").First(&sellerModel, key+" = ?", value).Error; err != nil {
		return nil, err
	}
	return mappers.SellerModelToEntity(&sellerModel), nil
}

func (s *SellerRepository) Update(seller *entities.Seller) error {
	sellerModel := mappers.SellerEntityToSellerModel(seller)
	if err := s.db.Save(sellerModel).Error; err != nil {
		return err
	}
	return nil
}

func (s *SellerRepository) Exists(param map[string]string) (bool, error) {
	key, value, _ := pkg.GetKeyValueFromMap(param)
	var seller models.SellerModel
	s.db.First(&seller, key+" = ?", value)
	if pkg.IsEmptyUUID(seller.ID) {
		return false, errors.New("seller not found")
	}
	return true, nil
}

func (a *SellerRepository) Create(seller *entities.Seller) error {
	enderecoID := pkg.NewUUID()
	err := a.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&models.AddressModel{
			ID:                 enderecoID,
			ZipCode:            seller.Address.ZipCode,
			State:              seller.Address.State,
			City:               seller.Address.City,
			AddressDescription: seller.Address.Description,
			Number:             seller.Address.Number,
			Complement:         seller.Address.Complement,
			Neighborhood:       seller.Address.Neighborhood,
			ReferencePoint:     seller.Address.ReferencePoint,
			Observation:        seller.Address.Observation,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		}).Error; err != nil {
			return err
		}
		if err := tx.Save(&models.SellerModel{
			ID:            seller.ID,
			NickName:      seller.NickName,
			CorporateName: seller.LegalName,
			Document:      seller.Document,
			Active:        *seller.Active,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			AccountID:     seller.AccountID,
			AddressID:     enderecoID,
		}).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
