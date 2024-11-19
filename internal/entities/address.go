package entities

import (
	"errors"
	"strings"
	"time"

	"github.com/andremartinsds/go_admin/internal/dto"
	"github.com/andremartinsds/go_admin/pkg"
)

// Address represents a location with detailed information.
type Address struct {
	// ID is the unique identifier for the address.
	ID pkg.ID

	// ZipCode is the postal code for the address.
	ZipCode string

	// Description provides a detailed description of the address.
	Description string

	// State is the state where the address is located.
	State string

	// City is the city where the address is located.
	City string

	// Number is the street number of the address.
	Number string

	// Complement is additional information for the address.
	Complement string

	// Neighborhood is the neighborhood of the address.
	Neighborhood string

	// ReferencePoint provides a point of reference for locating the address.
	ReferencePoint string

	// Observation contains any additional notes about the address.
	Observation string

	// CreatedAt is the timestamp when the account was created.
	CreatedAt time.Time

	// UpdatedAt is the timestamp when the account was last updated.
	UpdatedAt time.Time
}

// Validate checks if the address fields are valid.
// It returns an error if any required fields are missing.
func (e *Address) Validate() error {
	messages := []string{}

	if e.ZipCode == "" {
		messages = append(messages, "address.zipCode is required")
	}
	if len(e.ZipCode) > 8 {
		messages = append(messages, "address.zipCode is too long")
	}
	if e.State == "" {
		messages = append(messages, "address.state is required")
	}
	if e.City == "" {
		messages = append(messages, "address.city is required")
	}
	if e.Number == "" {
		messages = append(messages, "address.number is required")
	}
	if e.Complement == "" {
		messages = append(messages, "address.complement is required")
	}
	if e.Neighborhood == "" {
		messages = append(messages, "address.neighborhood is required")
	}
	if e.ReferencePoint == "" {
		messages = append(messages, "address.referencePoint is required")
	}
	if e.Observation == "" {
		messages = append(messages, "address.observation is required")
	}
	if len(messages) > 0 {
		return errors.New("[" + strings.Join(messages, ",") + "]")
	}

	return nil
}

// NewAddress creates a new Address from the provided AddressInputCreateDTO.
// It returns the newly created Address and an error if validation fails.
func NewAddress(addressDto dto.AddressInputCreateDTO) (error, *Address) {
	address := Address{
		ID:             pkg.NewUUID(),
		ZipCode:        addressDto.ZipCode,
		State:          addressDto.State,
		City:           addressDto.City,
		Description:    addressDto.Description,
		Number:         addressDto.Number,
		Complement:     addressDto.Complement,
		Neighborhood:   addressDto.Neighborhood,
		ReferencePoint: addressDto.ReferencePoint,
		Observation:    addressDto.Observation,
	}
	address.clearZipCode()
	err := address.Validate()
	if err != nil {
		return err, nil
	}

	return nil, &address
}

func (a *Address) clearZipCode() {
	a.ZipCode = strings.ReplaceAll(a.ZipCode, "-", "")
	a.ZipCode = strings.ReplaceAll(a.ZipCode, ".", "")
	a.ZipCode = strings.ReplaceAll(a.ZipCode, "/", "")
	a.ZipCode = strings.ReplaceAll(a.ZipCode, "_", "")
}

// UpdateAddress updates an existing Address with the provided AddressInputUpdateDTO.
// It returns the updated Address and an error if the update fails.
func UpdateAddress(addressDto dto.AddressInputUpdateDTO) (error, *Address) {
	addressID, _ := pkg.ParseID(addressDto.ID)
	address := Address{
		ID:             addressID,
		ZipCode:        addressDto.ZipCode,
		State:          addressDto.State,
		City:           addressDto.City,
		Description:    addressDto.Description,
		Number:         addressDto.Number,
		Complement:     addressDto.Complement,
		Neighborhood:   addressDto.Neighborhood,
		ReferencePoint: addressDto.ReferencePoint,
		CreatedAt:      addressDto.CreatedAt,
		UpdatedAt:      time.Now(),
		Observation:    addressDto.Observation,
	}
	address.clearZipCode()
	err := address.Validate()
	if err != nil {
		return err, nil
	}

	return nil, &address
}
