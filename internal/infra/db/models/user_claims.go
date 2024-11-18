package models

import (
	"time"

	"github.com/andremartinsds/go_admin/pkg"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserClaimModel struct {
	ID uuid.UUID `gorm:"type:char(36);primaryKey"`

	ClaimID uuid.UUID   `gorm:"type:char(36);not null"`
	Claim   ClaimsModel `gorm:"foreignKey:ClaimID"`

	UserID uuid.UUID `gorm:"type:char(36);not null"`
	User   UserModel `gorm:"foreignKey:UserID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *UserClaimModel) TableName() string {
	return "user_claims"
}

func (u *UserClaimModel) BeforeCreate(tx *gorm.DB) error {
	u.ID = pkg.NewUUID()
	return nil
}
