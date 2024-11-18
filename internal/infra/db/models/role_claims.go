package models

import (
	"time"

	"github.com/andremartinsds/go_admin/pkg"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleClaimsModel struct {
	ID        uuid.UUID   `gorm:"type:char(36);primaryKey"`
	RoleID    uuid.UUID   `gorm:"type:char(36);not null"`
	Role      RolesModel  `gorm:"foreignKey:RoleID"`
	ClaimID   uuid.UUID   `gorm:"type:char(36);not null"`
	Claim     ClaimsModel `gorm:"foreignKey:ClaimID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (r *RoleClaimsModel) TableName() string {
	return "role_claims"
}

func (r *RoleClaimsModel) BeforeCreate(tx *gorm.DB) error {
	r.ID = pkg.NewUUID()
	return nil
}
