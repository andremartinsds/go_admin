package dto

// RoleInputCreateDTO represents the data transfer object used to create a new role.
// It includes the necessary information to create a role.
type RoleInputCreateDTO struct {
	// Name is the full name of the role.
	Description string `json:"description"`

	// Phone is the role's phone number.
	Path string `json:"path"`

	// SellerID is the unique identifier of the associated seller.
	SellerID string `json:"sellerId"`

	// AccountID is the unique identifier of the associated account.
	AccountID string `json:"accountId"`
}
