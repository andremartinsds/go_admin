package repositories

import (
	"errors"
	"github.com/andremartinsds/go_admin/internal/entities"
	"github.com/andremartinsds/go_admin/internal/infra/db/models"
	"github.com/andremartinsds/go_admin/internal/mappers"
	"github.com/andremartinsds/go_admin/pkg"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"time"
)

type SellerContract interface {
	Create(seller *entities.Seller) error
	Exists(param map[string]string) (bool, error)
	SelectOneById(id string) (*entities.Seller, error)
	Update(seller *entities.Seller) error
	Select(param map[string]string) (*entities.Seller, error)
	List(accountID string) ([]*entities.Seller, error)
}

type SellerRepository struct {
	db *gorm.DB
}

func (s *SellerRepository) SelectOneById(id string) (*entities.Seller, error) {
	var seller models.SellerModel
	s.db.Preload("Endereco").First(&seller, "id = ?", id)
	if lo.IsEmpty(&seller) {
		return nil, errors.New("seller does not found")
	}
	sellerEntity := mappers.SellerModelToEntity(&seller)
	return sellerEntity, nil
}

func (s *SellerRepository) List(accountID string) ([]*entities.Seller, error) {
	var sellers []models.SellerModel
	err := s.db.Preload("Address").Preload("Account").Preload("Account.Address").Find(&sellers).Where("account_id = ?", accountID).Error
	if err != nil || len(sellers) == 0 {
		return nil, errors.New("does not exist seller to this account")
	}
	var sellerEntity []*entities.Seller
	for _, seller := range sellers {
		s := mappers.SellerModelToEntity(&seller)
		sellerEntity = append(sellerEntity, s)
	}
	return sellerEntity, nil
}

func SellerRepositoryInstancy(connection *gorm.DB) *SellerRepository {
	return &SellerRepository{db: connection}
}

func (s *SellerRepository) Select(param map[string]string) (*entities.Seller, error) {
	key, value, _ := pkg.GetKeyValueFromMap(param)
	var sellerModel models.SellerModel
	if err := s.db.Preload("Address").Preload("Account").Preload("Account.Address").First(&sellerModel, key+"=?", value).Error; err != nil {
		return nil, err
	}
	return mappers.SellerModelToEntity(&sellerModel), nil
}

func (s *SellerRepository) Update(seller *entities.Seller) error {
	if err := s.db.Save(mappers.SellerEntityToSellerModel(seller)).Error; err != nil {
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
