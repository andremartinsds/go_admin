package models

import (
	"time"

	"github.com/andremartinsds/go_admin/pkg"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RolesModel struct {
	ID          uuid.UUID    `gorm:"type:char(36);primaryKey"`
	Description string       `gorm:"type:varchar(100);not null"`
	Path        string       `gorm:"type:varchar(15);not null"`
	SellerID    uuid.UUID    `gorm:"type:char(36);not null"`
	Seller      SellerModel  `gorm:"foreignKey:SellerID"`
	AccountID   uuid.UUID    `gorm:"type:char(36);not null"`
	Account     AccountModel `gorm:"foreignKey:AccountID"`
	Users       []UserModel  `gorm:"foreignKey:RoleID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (r *RolesModel) TableName() string {
	return "roles"
}

func (r *RolesModel) BeforeCreate(tx *gorm.DB) error {
	r.ID = pkg.NewUUID()
	return nil
}
