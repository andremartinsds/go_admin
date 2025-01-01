package entities

import (
	"errors"
	"github.com/andremartinsds/go_admin/internal/dto"
	"github.com/andremartinsds/go_admin/pkg"
	"strings"
	"time"
)

// Account represents a financial account with relevant details.
type Account struct {
	// ID is the unique identifier for the account.
	ID pkg.ID

	// Name is the fantasy name of the account.
	NickName string

	// CommercialName is the legal name of the account.
	CommercialName string

	// Document is the identifier for the account, such as a tax identification number.
	Document string

	// Active indicates whether the account is active or not.
	Active *bool

	// AccountType represents the type of the account.
	AccountType string

	// CreatedAt is the timestamp when the account was created.
	CreatedAt time.Time

	// UpdatedAt is the timestamp when the account was last updated.
	UpdatedAt time.Time

	// Address contains the details of the account's address.
	Address *Address
}

// validateToCreate validates the account details before creation.
// It returns an error if any required fields are missing.
func (a *Account) validateToCreate() error {
	errs := []string{}

	if a.NickName == "" {
		errs = append(errs, "field: Name is required")
	}

	if len(a.AccountType) > 2 {
		errs = append(errs, "field: accountType accepted only two character")
	}

	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "\n"))
	}

	return nil
}

// NewAccount creates a new Account from the provided AccountInputCreateDTO.
// It returns the newly created Account and an error if any validation fails.
func NewAccount(accountCreateDto dto.AccountInputCreateDTO) (*Account, error) {
	err, addressEntity := NewAddress(accountCreateDto.Address)
	if err != nil {
		return nil, err
	}

	account := Account{
		ID:             pkg.NewUUID(),
		NickName:       accountCreateDto.Name,
		CommercialName: accountCreateDto.CommercialName,
		Document:       accountCreateDto.Document,
		Active:         &accountCreateDto.Active,
		AccountType:    accountCreateDto.AccountType,
		Address:        addressEntity,
	}

	err = account.validateToCreate()
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// UpdateAccount updates an existing Account with the provided AccountInputUpdateDTO.
// It returns the updated Account and an error if the update fails.
func UpdateAccount(accountUpdateDto dto.AccountInputUpdateDTO) (*Account, error) {
	err, addressEntity := UpdateAddress(accountUpdateDto.Address)
	if err != nil {
		return nil, err
	}
	accountID, _ := pkg.StringToUUID(accountUpdateDto.Id)
	account := Account{
		ID:             accountID,
		NickName:       accountUpdateDto.Name,
		CommercialName: accountUpdateDto.CommercialName,
		Document:       accountUpdateDto.Document,
		Active:         &accountUpdateDto.Active,
		AccountType:    accountUpdateDto.AccountType,
		Address:        addressEntity,
	}

	return &account, nil
}
