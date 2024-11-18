package dto

// AccountInputCreateDTO represents the data transfer object used to create a new account.
// It includes personal, commercial, and address details required for account setup.
type AccountInputCreateDTO struct {
	// Name is the legal name of the account holder or organization.
	Name string `json:"name"`

	// CommercialName is the business or trading name of the account.
	CommercialName string `json:"commercialName"`

	// Document holds the identification number associated with the account.
	// This could be a tax ID, social security number, or business registration number.
	Document string `json:"document"`

	// Active indicates whether the account is currently active.
	Active bool `json:"active"`

	// AccountType specifies the type of account being created (e.g., "personal", "business").
	AccountType string `json:"accountType"`

	// Address is an embedded struct containing the address details for the account.
	Address AddressInputCreateDTO `json:"address"`
}

// AccountInputDeleteDTO represents the data transfer object used to request the deletion of an existing account.
// It contains the identifier of the account that should be removed.
type AccountInputDeleteDTO struct {
	// Id is the unique identifier of the account to be deleted.
	Id string `json:"id"`
}

// AccountInputSelectDto represents the data transfer object used to select an existing account.
// It contains the identifier of the account that should be retrieved.
type AccountInputSelectDto struct {
	// Id is the unique identifier of the account to be selected.
	Id string `json:"id"`
}

// AccountInputUpdateDTO representa o objeto de transferência de dados usado para atualizar uma conta existente.
// Ele inclui informações pessoais, comerciais e detalhes de endereço necessários para a atualização da conta.
type AccountInputUpdateDTO struct {
	// Id é o identificador único da conta a ser atualizada.
	Id string `json:"id"`

	// Name is the legal name of the account holder or organization.
	Name string `json:"name"`

	// CommercialName is the business or trading name of the account.
	CommercialName string `json:"commercialName"`

	// Document holds the identification number associated with the account.
	// This could be a tax ID, social security number, or business registration number.
	Document string `json:"document"`

	// Active indicates whether the account is currently active.
	Active bool `json:"active"`

	// AccountType specifies the type of account being created (e.g., "personal", "business").
	AccountType string `json:"accountType"`

	// Address represents the data transfer object address for account.
	Address AddressInputUpdateDTO `json:"address"`
}

// AccountOutputDto represents the data transfer object used to output account information.
// It includes details about the account that is being returned in responses.
type AccountOutputDto struct {
	// Id is the unique identifier of the account.
	Id string `json:"id"`

	// Nickname is the display name of the account.
	Nickname string `json:"nickname"`

	Document string `json:"cnpj"`

	// Active indicates whether the account is currently active.
	Active bool `json:"active"`

	// AccountType specifies the type of account.
	AccountType string `json:"accountType"`

	// Address is the output structure containing address details associated with the account.
	Address *AddressOutputDTO `json:"address"`
}
