package repositories

import (
	"errors"
	"fmt"
	"github.com/andremartinsds/go_admin/internal/entities"
	"time"

	"github.com/andremartinsds/go_admin/internal/infra/db/models"
	"github.com/andremartinsds/go_admin/internal/mappers"
	"github.com/andremartinsds/go_admin/pkg"
	"gorm.io/gorm"
)

type AccountContract interface {
	Create(account *entities.Account) error
	ExistsByField(param map[string]string) (bool, error)
	SelectOneById(id string) (*entities.Account, error)
	UpdateOne(account *entities.Account) error
	List() (*[]entities.Account, error)
}

type AccountRepository struct {
	db *gorm.DB
}

func AccountRepositoryInstance(connection *gorm.DB) *AccountRepository {
	return &AccountRepository{db: connection}
}

func (a *AccountRepository) List() (*[]entities.Account, error) {
	var accountsModel []models.AccountModel
	err := a.db.Preload("Address").Find(&accountsModel).Error
	if err != nil {
		return nil, err
	}
	var accounts []entities.Account
	for _, account := range accountsModel {
		accountsEntity := mappers.ToAccountEntity(account)
		accounts = append(accounts, *accountsEntity)
	}

	return &accounts, nil
}

func (a *AccountRepository) UpdateOne(account *entities.Account) error {
	if (entities.Address{}) != *account.Address {
		a.db.Debug()
		err := a.db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Save(&models.AddressModel{
				ID:                 account.Address.ID,
				ZipCode:            account.Address.ZipCode,
				State:              account.Address.State,
				City:               account.Address.City,
				AddressDescription: account.Address.Description,
				Number:             account.Address.Number,
				Complement:         account.Address.Complement,
				Neighborhood:       account.Address.Neighborhood,
				ReferencePoint:     account.Address.ReferencePoint,
				CreatedAt:          account.Address.CreatedAt,
				Observation:        account.Address.Observation,
				UpdatedAt:          time.Now(),
			}).Error; err != nil {
				return err
			}
			if err := tx.Save(&models.AccountModel{
				ID:             account.ID,
				NickName:       account.NickName,
				CommercialName: account.CommercialName,
				Document:       account.Document,
				Active:         *account.Active,
				AccountType:    account.AccountType,
				CreatedAt:      account.CreatedAt,
				UpdatedAt:      time.Now(),
				AddressID:      account.Address.ID,
			}).Error; err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *AccountRepository) SelectOneById(id string) (*entities.Account, error) {
	var account models.AccountModel
	a.db.Preload("Address").First(&account, "id=?", id)
	if account.Document == "" {
		return nil, errors.New("account does not found")
	}
	accountEntity := mappers.ToAccountEntity(account)
	return accountEntity, nil
}

func (a *AccountRepository) ExistsByField(param map[string]string) (bool, error) {
	var account models.AccountModel
	parameter, value, err := pkg.GetKeyValueFromMap(param)
	if err != nil {
		return false, fmt.Errorf("map error")
	}
	fmt.Println(parameter + " ?")
	a.db.Preload("Address").First(&account, parameter+"= ?", value)
	if account.Document == "" {
		return false, fmt.Errorf("user does not found for %s and %s", parameter, value)
	}
	return true, nil
}

func (a *AccountRepository) Create(account *entities.Account) error {
	var addressID = pkg.NewUUID()
	a.db.Debug()
	err := a.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&models.AddressModel{
			ID:                 addressID,
			ZipCode:            account.Address.ZipCode,
			State:              account.Address.State,
			City:               account.Address.City,
			AddressDescription: account.Address.Description,
			Number:             account.Address.Number,
			Complement:         account.Address.Complement,
			Neighborhood:       account.Address.Neighborhood,
			ReferencePoint:     account.Address.ReferencePoint,
			Observation:        account.Address.Observation,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		}).Error; err != nil {
			return err
		}
		if err := tx.Save(&models.AccountModel{
			NickName:       account.NickName,
			CommercialName: account.CommercialName,
			Document:       account.Document,
			Active:         *account.Active,
			AccountType:    account.AccountType,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			AddressID:      addressID,
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
