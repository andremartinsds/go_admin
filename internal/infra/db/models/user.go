package models

import (
	"time"

	"github.com/andremartinsds/go_admin/pkg"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserModel struct {
	ID        uuid.UUID    `gorm:"type:char(36);primaryKey"`
	Name      string       `gorm:"type:varchar(100);not null"`
	Phone     string       `gorm:"type:varchar(15);default null"`
	Email     string       `gorm:"type:varchar(100);not null"`
	Password  string       `gorm:"type:varchar(255);not null"`
	Document  string       `gorm:"type:varchar(20);not null"`
	BirthDate time.Time    `gorm:"type:varchar(20);default null"`
	Provider  bool         `gorm:"not null"`
	SellerID  uuid.UUID    `gorm:"type:char(36);not null"`
	Seller    SellerModel  `gorm:"foreignKey:SellerID"`
	AccountID uuid.UUID    `gorm:"type:char(36);not null"`
	Account   AccountModel `gorm:"foreignKey:AccountID"`
	AddressID uuid.UUID    `gorm:"type:char(36);default null"`
	Address   AddressModel `gorm:"foreignKey:AddressID"`
	RoleID    uuid.UUID    `gorm:"type:char(36);default null"`
	Role      RolesModel   `gorm:"foreignKey:RoleID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *UserModel) TableName() string {
	return "users"
}

func (u *UserModel) BeforeCreate(tx *gorm.DB) error {
	u.ID = pkg.NewUUID()
	return nil
}
