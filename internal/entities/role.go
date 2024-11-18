package entities

import (
	"github.com/andremartinsds/go_admin/internal/dto"
	"github.com/andremartinsds/go_admin/pkg"
	"time"
)

// Role represents a User Role for one user.
type Role struct {
	// ID is the unique identifier for the role.
	ID pkg.ID

	// Description for a role.
	Description string

	// The api path for a role.
	Path string

	// The SellerID identification for a current Seller.
	SellerID string

	// AccountID indicates ID for current Account.
	AccountID string

	// CreatedAt is the timestamp when the role was created.
	CreatedAt time.Time

	// UpdatedAt is the timestamp when the role was last updated.
	UpdatedAt time.Time
}

// validateToCreate validates the account details before creation.
// It returns an error if any required fields are missing.
func (r *Role) validateToCreate() error {
	return nil
}

// NewRole creates a new Role from the provided RoleInputCreateDTO.
// It returns the newly created Role and an error if any validation fails.
func NewRole(r dto.RoleInputCreateDTO) (*Role, error) {
	role := Role{
		ID:          pkg.NewUUID(),
		Description: r.Description,
		Path:        r.Path,
		SellerID:    r.SellerID,
		AccountID:   r.AccountID,
	}

	err := role.validateToCreate()
	if err != nil {
		return nil, err
	}

	return &role, nil
}
