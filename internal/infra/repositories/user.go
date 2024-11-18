package repositories

import (
	"errors"
	"github.com/andremartinsds/go_admin/internal/entities"
	"time"

	"github.com/andremartinsds/go_admin/internal/infra/db/models"
	"github.com/andremartinsds/go_admin/internal/mappers"
	"github.com/andremartinsds/go_admin/pkg"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type UserContract interface {
	Create(account *entities.User) error
	UserExists(param map[string]string) error
	SelectOneById(id string) (*entities.User, error)
	Login(username, password string) (*entities.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func UserRepositoryInstancy(connection *gorm.DB) *UserRepository {
	return &UserRepository{db: connection}
}

func (u *UserRepository) Login(username, password string) (*entities.User, error) {
	var userModel models.UserModel
	err := u.db.Where("username = ?", username).Where("password =?", password).First(&userModel)
	if err != nil {
		return nil, errors.New("the username " + username + " or password " + password + " are incorrect")
	}
	return mappers.FromUserModelToUserEntity(&userModel), nil
}

func (u *UserRepository) UserExists(param map[string]string) error {
	parameter, value, err := pkg.GetKeyValueFromMap(param)
	if err != nil {
		return err
	}
	var userModel *models.UserModel
	u.db.Debug()
	u.db.First(&userModel, parameter+"=?", value)
	if lo.IsNotEmpty(userModel.Email) {
		return errors.New("the user already exists")
	}
	return nil
}

func (u *UserRepository) Create(user *entities.User) error {
	enderecoId := pkg.NewUUID()

	u.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Save(&models.AddressModel{
			ZipCode:            user.Address.ZipCode,
			State:              user.Address.State,
			City:               user.Address.City,
			AddressDescription: user.Address.Description,
			Number:             user.Address.Number,
			Complement:         user.Address.Complement,
			Neighborhood:       user.Address.Neighborhood,
			ReferencePoint:     user.Address.ReferencePoint,
			Observation:        user.Address.Observation,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		}).Error; err != nil {
			return err
		}

		if err := tx.Save(&models.UserModel{
			Name:      user.Name,
			Phone:     user.Phone,
			Email:     user.Email,
			Provider:  *user.Provider,
			Password:  user.Password,
			Document:  user.Document,
			AccountID: user.AccountID,
			SellerID:  user.SellerID,
			AddressID: enderecoId,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}).Error; err != nil {
			return err
		}
		return nil
	})
	return nil
}

func (u *UserRepository) SelectOneById(id string) (*entities.User, error) {
	var user models.UserModel

	if err := u.db.First(&user, "id=?", id).Error; err != nil {
		return nil, err
	}
	userEntity := mappers.FromUserModelToUserEntity(&user)

	if lo.IsEmpty(userEntity) {
		return nil, errors.New("the user does not found")
	}
	return *&userEntity, nil
}
