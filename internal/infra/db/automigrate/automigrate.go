package automigrate

import (
	"github.com/andremartinsds/go_admin/internal/infra/db/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(models.AccountModel{}, models.SellerModel{}, models.ContactModel{}, models.AddressModel{}, models.SellerModel{}, models.UserModel{})
	if err != nil {
		return err
	}
	return nil
}
