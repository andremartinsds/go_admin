package mappers

import (
	"github.com/andremartinsds/go_admin/internal/dto"
	"github.com/andremartinsds/go_admin/internal/entities"
	"github.com/andremartinsds/go_admin/internal/infra/db/models"
	"github.com/samber/lo"
)

func FromUserToUserOutputCreateDTO(user *entities.User) *dto.UserOutputCreateDTO {
	return &dto.UserOutputCreateDTO{
		Name:      user.Name,
		Phone:     user.Phone,
		Email:     user.Email,
		Document:  user.Document,
		Provider:  *user.Provider,
		SellerID:  user.SellerID.String(),
		AccountID: user.AccountID.String(),
	}
}

func FromUserModelToUserEntity(userModel *models.UserModel) *entities.User {
	var address entities.Address
	if lo.IsNotEmpty(&userModel.Address) {
		address = *EnderecoModelToEntity(userModel.Address)
	}
	return &entities.User{
		ID:          userModel.ID,
		Name:        userModel.Name,
		Phone:       userModel.Phone,
		Email:       userModel.Email,
		Password:    userModel.Password,
		Document:    userModel.Document,
		DateOfBirth: userModel.BirthDate,
		Provider:    &userModel.Provider,
		SellerID:    userModel.SellerID,
		AccountID:   userModel.AccountID,
		CreatedAt:   userModel.CreatedAt,
		UpdatedAt:   userModel.UpdatedAt,
		Address:     &address,
	}
}
