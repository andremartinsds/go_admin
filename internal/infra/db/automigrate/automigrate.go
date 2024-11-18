package automigrate

import (
	"github.com/andremartinsds/go_admin/internal/infra/db/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(models.AddressModel{}, models.AccountModel{}, models.SellerModel{}, models.ContactModel{},
		models.SellerModel{}, models.UserModel{}, models.RolesModel{}, models.ClaimsModel{},
		models.UserClaimModel{}, models.RoleClaimsModel{}, models.OpenRoutes{})
	if err != nil {
		return err
	}
	return nil
}
