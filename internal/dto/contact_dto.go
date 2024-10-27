package dto

// ContactInputCreateDTO represents the data transfer object used to create a new contact.
// It includes the necessary information to create a contact record.
type ContactInputCreateDTO struct {
	// Name is the full name of the contact.
	Name string `json:"name"`

	// Phone is the contact's phone number.
	Phone string `json:"phone"`

	// Email is the contact's email address.
	Email string `json:"email"`

	// Message contains any message or notes from the contact.
	Message string `json:"message"`
}

// ContactsOutputDto represents the data transfer object used to output contact information.
// It includes details about the contact being returned in responses.
type ContactsOutputDto struct {
	// ID is the unique identifier of the contact.
	ID string `json:"id"`

	// Name is the full name of the contact.
	Name string `json:"name"`

	// Phone is the contact's phone number.
	Phone string `json:"phone"`

	// Message contains any message or notes associated with the contact.
	Message string `json:"message"`

	// Email is the contact's email address.
	Email string `json:"email"`
}
