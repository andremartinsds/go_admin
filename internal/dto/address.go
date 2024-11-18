package dto

// AddressInputCreateDTO represents the data transfer object used to create a new address.
// It includes detailed information about the address that needs to be created.
type AddressInputCreateDTO struct {
	// ZipCode is the postal code of the address.
	ZipCode string `json:"zipCode"`

	// Description provides a detailed description of the address.
	Description string

	// State is the name of the state where the address is located.
	State string `json:"state"`

	// City is the name of the city where the address is located.
	City string `json:"city"`

	// Street is the street name or description of the address.
	Street string `json:"street"`

	// Number is the house number associated with the address.
	Number string `json:"number"`

	// Complement provides additional address information, such as apartment number or suite.
	Complement string `json:"complement"`

	// Neighborhood is the neighborhood where the address is located.
	Neighborhood string `json:"neighborhood"`

	// ReferencePoint is a landmark or reference point near the address.
	ReferencePoint string `json:"referencePoint"`

	// Observation contains any additional observations or notes regarding the address.
	Observation string `json:"observation"`
}

// AddressInputUpdateDTO represents the data transfer object used to update an existing address.
// It includes detailed information about the address that needs to be updated.
type AddressInputUpdateDTO struct {
	// ID is the unique identifier of the address to be updated.
	ID string `json:"id"`

	// ZipCode is the postal code of the address.
	ZipCode string `json:"zipCode"`

	// Description is a detailed description of the address.
	Description string `json:"description"`

	// State is the name of the state where the address is located.
	State string `json:"state"`

	// City is the name of the city where the address is located.
	City string `json:"city"`

	// Street is the street name or description of the address.
	Street string `json:"street"`

	// Number is the house number associated with the address.
	Number string `json:"number"`

	// Complement provides additional address information, such as apartment number or suite.
	Complement string `json:"complement"`

	// Neighborhood is the neighborhood where the address is located.
	Neighborhood string `json:"neighborhood"`

	// ReferencePoint is a landmark or reference point near the address.
	ReferencePoint string `json:"referencePoint"`

	// Observation contains any additional observations or notes regarding the address.
	Observation string `json:"observation"`
}

// AddressOutputDTO represents the data transfer object used to output address information.
// It includes detailed information about the address being returned in responses.
type AddressOutputDTO struct {
	// ID is the unique identifier of the address.
	ID string `json:"id"`

	// ZipCode is the postal code of the address.
	ZipCode string `json:"zipCode"`

	// Description is the description of the address.
	Description string `json:"description"`

	// State is the name of the state where the address is located.
	State string `json:"state"`

	// City is the name of the city where the address is located.
	City string `json:"city"`

	// Street is the street name or description of the address.
	Street string `json:"street"`

	// Number is the house number associated with the address.
	Number string `json:"number"`

	// Complement provides additional address information, such as apartment number or suite.
	Complement string `json:"complement"`

	// Neighborhood is the neighborhood where the address is located.
	Neighborhood string `json:"neighborhood"`

	// ReferencePoint is a landmark or reference point near the address.
	ReferencePoint string `json:"referencePoint"`

	// Observation contains any additional observations or notes regarding the address.
	Observation string `json:"observation"`
}
