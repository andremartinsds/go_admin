package models

import (
	"time"

	"github.com/andremartinsds/go_admin/pkg"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClaimsModel struct {
	ID          uuid.UUID    `gorm:"type:char(36);primaryKey"`
	Description string       `gorm:"type:varchar(100);default null"`
	Path        string       `gorm:"type:varchar(255);not null"`
	SellerID    uuid.UUID    `gorm:"type:char(36);not null"`
	Seller      SellerModel  `gorm:"foreignKey:SellerID"`
	AccountID   uuid.UUID    `gorm:"type:char(36);not null"`
	Account     AccountModel `gorm:"foreignKey:AccountID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (c *ClaimsModel) TableName() string {
	return "claims"
}

func (c *ClaimsModel) BeforeCreate(tx *gorm.DB) error {
	c.ID = pkg.NewUUID()
	return nil
}
