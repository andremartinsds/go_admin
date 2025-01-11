package entities

import (
	"errors"
	"strings"
	"time"

	"github.com/andremartinsds/go_admin/internal/dto"
	"github.com/andremartinsds/go_admin/pkg"
	"github.com/andremartinsds/go_admin/pkg/auth"
)

// User represents a user entity with relevant fields.
type User struct {
	ID          pkg.ID    // Unique identifier for the user
	Name        string    // Name of the user
	Phone       string    // Phone number of the user
	Email       string    // Email address of the user
	Password    string    // Password for the user account
	Document    string    // Document identifier (e.g., CPF or ID number)
	DateOfBirth time.Time // Date of birth of the user
	Provider    *bool     // Indicates if the user is a provider
	SellerID    pkg.ID    // Identifier for the seller associated with the user
	RoleID      pkg.ID    // Identifier for the seller associated with the user
	AccountID   pkg.ID    // Identifier for the account associated with the user
	CreatedAt   time.Time // Timestamp for when the user was created
	UpdatedAt   time.Time // Timestamp for when the user was last updated
	Address     *Address  // Pointer to the user's address
	Claims      *Claims   //Pointer to the user's claims
}

// validateToCreate checks the User fields for required values before creation.
func (a *User) validateToCreate() error {
	errs := []string{}

	if a.Name == "" {
		errs = append(errs, "field: name is required")
	}

	if a.Phone == "" {
		errs = append(errs, "field: phone is required")
	}

	if a.Email == "" {
		errs = append(errs, "field: email is required")
	}

	if a.Password == "" {
		errs = append(errs, "field: password is required")
	}

	if a.Document == "" {
		errs = append(errs, "field: document is required")
	}

	if a.Provider == nil {
		errs = append(errs, "field: provider is required")
	}

	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "\n"))
	}

	return nil
}

// ValidatePassword checks if the provided password and password confirmation match.
func (a *User) ValidatePassword(requestPass string) bool {
	return auth.IsValidPassword(a.Password, requestPass)
}

func (a *User) HandleBirthDate(date string) {
	time, _ := time.Parse(time.RFC3339, date)
	a.DateOfBirth = time.UTC()
}

// CreateUser initializes a new User based on the provided input DTO.
func CreateUser(userInputDTO dto.UserInputCreateDTO) (*User, error) {
	// Create a new address entity from the input DTO
	err, addressEntity := NewAddress(userInputDTO.Address)
	if err != nil {
		return nil, err
	}

	// Parse SellerID and AccountID from strings to IDs
	SellerID, _ := pkg.StringToUUID(userInputDTO.SellerID)
	accountID, _ := pkg.StringToUUID(userInputDTO.AccountID)
	roleID, _ := pkg.StringToUUID(userInputDTO.RoleID)

	passwordEncripted, err := auth.Encrypt(userInputDTO.Password)
	if err != nil {
		return nil, errors.New("we have a problemn to ecript password")
	}

	// Initialize a new User instance
	user := User{
		ID:        pkg.NewUUID(),
		Name:      userInputDTO.Name,
		Phone:     userInputDTO.Phone,
		Email:     userInputDTO.Email,
		Document:  userInputDTO.Document,
		Provider:  &userInputDTO.Provider,
		Password:  passwordEncripted,
		SellerID:  SellerID,
		RoleID:    roleID,
		AccountID: accountID,
		Address:   addressEntity,
	}

	user.HandleBirthDate(userInputDTO.BirthDate)

	// Validate the user before creation
	err = user.validateToCreate()
	if err != nil {
		return nil, err
	}

	// Validate the password confirmation
	if err != nil {
		return nil, err
	}

	return &user, nil // Return the newly created user
}
