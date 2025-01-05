package dto

import "time"

// SellerInputCreateDTO represents the data transfer object used to create a new seller.
// It includes the necessary information to create a seller record.
type SellerInputCreateDTO struct {
	// AccountID is the unique identifier of the associated account.
	AccountID string `json:"accountId"`

	// Nickname is the display name of the seller.
	Nickname string `json:"nickname"`

	// CorporateName is the legal name of the seller.
	CorporateName string `json:"corporateName"`

	// Document is the identifier for the seller, such as a tax identification number.
	Document string `json:"document"`

	// Active indicates whether the seller is currently active.
	Active bool `json:"active"`

	// Address contains the details of the seller's address.
	Address AddressInputCreateDTO `json:"address"`
}

// SellerInputUpdateDTO represents the data transfer object used to update an existing seller.
// It includes the necessary information to modify a seller record.
type SellerInputUpdateDTO struct {
	// ID is the unique identifier of the seller.
	ID string `json:"id"`

	// AccountID is the unique identifier of the associated account.
	AccountID string `json:"accountId"`

	// Nickname is the display name of the seller.
	Nickname string `json:"nickname"`

	// CorporateName is the legal name of the seller.
	CorporateName string `json:"corporateName"`

	// Document is the identifier for the seller, such as a tax identification number.
	Document string `json:"document"`

	// Active indicates whether the seller is currently active.
	Active bool `json:"active"`

	CreatedAt time.Time

	// Address contains the details of the seller's address to be updated.
	Address AddressInputUpdateDTO `json:"address"`
}

// SellerOutputDTO represents the data transfer object used to output seller information.
// It includes details about the seller being returned in responses.
type SellerOutputDTO struct {
	// ID is the unique identifier of the seller.
	ID string `json:"id"`

	// Nickname is the display name of the seller.
	Nickname string `json:"nickname"`

	// Document is the identifier for the seller, such as a tax identification number.
	Document string `json:"document"`

	// Address contains the details of the seller's address.
	Address *AddressOutputDTO `json:"address"`
}
