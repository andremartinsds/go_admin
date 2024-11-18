package repositories

import (
	"github.com/andremartinsds/go_admin/internal/entities"
	"github.com/andremartinsds/go_admin/internal/infra/db/models"
	"github.com/andremartinsds/go_admin/pkg"
	"gorm.io/gorm"
	"time"
)

type RoleContract interface {
	CreateDefaultsRoles(r *entities.Role) error
}

type RoleRepository struct {
	db *gorm.DB
}

func RoleRepositoryInstancy(connection *gorm.DB) *RoleRepository {
	return &RoleRepository{db: connection}
}

func (r *RoleRepository) CreateDefaultsRoles(role *entities.Role) error {
	sellerID, _ := pkg.ParseID(role.SellerID)
	accountID, _ := pkg.ParseID(role.AccountID)

	if err := r.db.Create(&models.RolesModel{
		ID:          pkg.NewUUID(),
		Description: role.Description,
		Path:        role.Path,
		SellerID:    sellerID,
		AccountID:   accountID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}).Error; err != nil {
		return err
	}
	return nil
}
