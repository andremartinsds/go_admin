package entities

import (
	"github.com/google/uuid"
	"time"
)

type Claims struct {
	ID          uuid.UUID
	Description string
	Path        string
	SellerID    uuid.UUID
	AccountID   uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserClaims struct {
	UserID  uuid.UUID
	ClaimID uuid.UUID
}
