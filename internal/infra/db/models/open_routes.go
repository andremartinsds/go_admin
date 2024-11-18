package models

import (
	"time"

	"github.com/andremartinsds/go_admin/pkg"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OpenRoutes struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	Path      string    `gorm:"type:varchar(255);not null;index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (o *OpenRoutes) TableName() string {
	return "open_routes"
}

func (o *OpenRoutes) BeforeCreate(tx *gorm.DB) error {
	o.ID = pkg.NewUUID()
	return nil
}
