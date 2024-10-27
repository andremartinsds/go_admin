package models

import (
	"github.com/andremartinsds/go_admin/pkg"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type SellerModel struct {
	ID            uuid.UUID `gorm:"type:char(36);primaryKey"`
	NickName      string    `gorm:"type:varchar(100);not null"`
	CorporateName string    `gorm:"type:varchar(100);not null"`
	Document      string    `gorm:"type:varchar(20);not null"`
	Active        bool
	AccountID     uuid.UUID      `gorm:"type:char(36);not null"`
	Account       AccountModel   `gorm:"foreignKey:AccountID"`
	AddressID     uuid.UUID      `gorm:"type:char(36)"`
	Address       AddressModel   `gorm:"foreignKey:AddressID"`
	Users         []UserModel    `gorm:"foreignKey:SellerID"`
	Contacts      []ContactModel `gorm:"foreignKey:SellerID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func (s *SellerModel) TableName() string {
	return "sellers"
}

func (s *SellerModel) BeforeCreate(tx *gorm.DB) error {
	s.ID = pkg.NewUUID()
	return nil
}
