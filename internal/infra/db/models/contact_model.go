package models

import (
	"github.com/andremartinsds/go_admin/pkg"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ContactModel struct {
	ID        uuid.UUID    `gorm:"type:char(36);primaryKey"`
	Name      string       `gorm:"type:varchar(100);not null"`
	Phone     string       `gorm:"type:varchar(20);not null"`
	Email     string       `gorm:"type:varchar(50);not null"`
	Message   string       `gorm:"type:varchar(255);not null"`
	AccountID uuid.UUID    `gorm:"type:char(36);not null"`
	Account   AccountModel `gorm:"foreignKey:AccountID"`
	SellerID  uuid.UUID    `gorm:"type:char(36)"`
	Seller    SellerModel  `gorm:"foreignKey:SellerID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (c *ContactModel) TableName() string {
	return "contacts"
}

func (c *ContactModel) BeforeCreate(tx *gorm.DB) error {
	c.ID = pkg.NewUUID()
	return nil
}
