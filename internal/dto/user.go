package dto

import "time"

// UserInputCreateDTO represents the data transfer object used to create a new user.
// It includes the necessary information to create a user account.
type UserInputCreateDTO struct {
	// Name is the full name of the user.
	Name string `json:"name"`

	// Phone is the user's phone number.
	Phone string `json:"phone"`

	// Email is the user's email address.
	Email string `json:"email"`

	// Document is the identifier for the user, such as a tax identification number.
	Document string `json:"document"`

	// BirthDate is the user's date of birth.
	BirthDate time.Time `json:"birthDate"`

	// Provider indicates whether the user is a service provider.
	Provider bool `json:"provider"`

	// Password is the user's password for authentication.
	Password string `json:"password"`

	// PasswordConfirmation is the confirmation of the user's password.
	PasswordConfirmation string `json:"passwordConfirmation"`

	// SellerID is the unique identifier of the associated seller.
	SellerID string `json:"sellerId"`

	// AccountID is the unique identifier of the associated account.
	AccountID string `json:"accountId"`

	// Address contains the details of the user's address.
	Address AddressInputCreateDTO `json:"address"`
}

// UserInputUpdateDTO represents the data transfer object used to update an existing user.
// It includes the necessary information to modify a user's account details.
type UserInputUpdateDTO struct {
	// ID is the unique identifier of the user.
	ID string `json:"id"`

	// Name is the full name of the user.
	Name string `json:"name"`

	// Phone is the user's phone number.
	Phone string `json:"phone"`

	// Email is the user's email address.
	Email string `json:"email"`

	// Document is the identifier for the user, such as a tax identification number.
	Document string `json:"document"`

	// BirthDate is the user's date of birth.
	BirthDate time.Time `json:"birthDate"`

	// Provider indicates whether the user is a service provider.
	Provider bool `json:"provider"`

	// Password is the user's password for authentication. It may be empty if not updating.
	Password string `json:"password"`

	// PasswordConfirmation is the confirmation of the user's password. It may be empty if not updating.
	PasswordConfirmation string `json:"passwordConfirmation"`

	// SellerID is the unique identifier of the associated seller.
	SellerID string `json:"sellerId"`

	// AccountID is the unique identifier of the associated account.
	AccountID string `json:"accountId"`

	// Address contains the details of the user's address to be updated.
	Address AddressInputUpdateDTO `json:"address"`
}

// UserOutputCreateDTO represents the data transfer object used to output user information after creation.
// It includes the necessary details of the user that has been created.
type UserOutputCreateDTO struct {
	// Name is the full name of the user.
	Name string `json:"name"`

	// Phone is the user's phone number.
	Phone string `json:"phone"`

	// Email is the user's email address.
	Email string `json:"email"`

	// Document is the identifier for the user, such as a tax identification number.
	Document string `json:"document"`

	// Provider indicates whether the user is a service provider.
	Provider bool `json:"provider"`

	// SellerID is the unique identifier of the associated seller.
	SellerID string `json:"sellerId"`

	// AccountID is the unique identifier of the associated account.
	AccountID string `json:"accountId"`

	// Address contains the details of the user's address.
	Address AddressInputCreateDTO `json:"address"`
}
