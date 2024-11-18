package models

import (
	"github.com/andremartinsds/go_admin/pkg"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type AccountModel struct {
	ID             uuid.UUID      `gorm:"type:char(36);primaryKey"`
	NickName       string         `gorm:"type:varchar(100);not null"`
	CommercialName string         `gorm:"type:varchar(100);not null"`
	TaxID          string         `gorm:"type:varchar(20);not null"`
	Document       string         `gorm:"type:varchar(100);not null"`
	AccountType    string         `gorm:"type:char(2)"`
	Active         bool           `gorm:"not null"`
	AddressID      uuid.UUID      `gorm:"type:char(36);default null"`
	Address        AddressModel   `gorm:"foreignKey:AddressID"`
	Users          []UserModel    `gorm:"foreignKey:AccountID"`
	Sellers        []SellerModel  `gorm:"foreignKey:AccountID"`
	Contacts       []ContactModel `gorm:"foreignKey:AccountID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func (a *AccountModel) TableName() string {
	return "accounts"
}

func (a *AccountModel) BeforeCreate(tx *gorm.DB) error {
	a.ID = pkg.NewUUID()
	return nil
}
