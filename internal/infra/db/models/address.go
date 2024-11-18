package models

import (
	"github.com/andremartinsds/go_admin/pkg"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type AddressModel struct {
	ID                 uuid.UUID      `gorm:"type:char(36);primaryKey"`
	ZipCode            string         `gorm:"type:varchar(8);not null"`
	State              string         `gorm:"type:varchar(20);not null"`
	City               string         `gorm:"type:varchar(20);not null"`
	AddressDescription string         `gorm:"type:varchar(100);not null"`
	Neighborhood       string         `gorm:"type:varchar(100);not null"`
	Sellers            []SellerModel  `gorm:"foreignKey:AddressID"`
	Accounts           []AccountModel `gorm:"foreignKey:AddressID"`
	Users              []UserModel    `gorm:"foreignKey:AddressID"`
	Number             string         `gorm:"type:varchar(10)"`
	Complement         string         `gorm:"type:varchar(100);not null"`
	ReferencePoint     string         `gorm:"type:varchar(100);not null"`
	Observation        string         `gorm:"type:varchar(100);not null"`
	District           string         `gorm:"type:varchar(30)"`
	Landmark           string         `gorm:"type:varchar(100)"`
	Notes              string         `gorm:"type:varchar(100)"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}

func (e *AddressModel) TableName() string {
	return "addresses"
}

func (e *AddressModel) BeforeCreate(tx *gorm.DB) error {
	e.ID = pkg.NewUUID()
	return nil
}
