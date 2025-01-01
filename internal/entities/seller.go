package entities

import (
	"errors"
	"github.com/andremartinsds/go_admin/internal/dto"
	"github.com/andremartinsds/go_admin/pkg"
	"strings"
	"time"
)

// Seller represents a seller entity with relevant details.
type Seller struct {
	ID        pkg.ID    // Unique identifier for the seller
	AccountID pkg.ID    // Unique identifier for the account associated with the seller
	Account   Account   // Account details associated with the seller
	NickName  string    // Nickname or trade name of the seller
	LegalName string    // Legal name of the seller
	Document  string    // Document number (e.g., CNPJ)
	Active    *bool     // Indicates if the seller is active
	CreatedAt time.Time // Timestamp of when the seller was created
	UpdatedAt time.Time // Timestamp of when the seller was last updated
	Address   *Address  // Address details of the seller
}

// ValidateFieldsToCreate checks the required fields for creating a seller.
// Returns an error if any required field is missing.
func (s Seller) ValidateFieldsToCreate() error {
	message := []string{}

	if s.Document == "" {
		message = append(message, "seller.document is required")
	}
	if s.NickName == "" {
		message = append(message, "seller.nickname is required")
	}
	if s.LegalName == "" {
		message = append(message, "seller.legalName is required")
	}
	if s.Active == nil {
		message = append(message, "seller.active is required")
	}
	if len(message) > 0 {
		return errors.New("[" + strings.Join(message, ",") + "]")
	}

	return nil
}

// NewSeller creates a new Seller entity from the provided SellerInputCreateDTO.
// Returns a pointer to the created Seller and an error if any validations fail.
func NewSeller(sellerDto dto.SellerInputCreateDTO) (*Seller, error) {
	err, address := NewAddress(sellerDto.Address)
	if err != nil {
		return nil, err
	}

	accountID, _ := pkg.StringToUUID(sellerDto.AccountID)

	sellerEntity := Seller{
		ID:        pkg.NewUUID(),
		AccountID: accountID,
		NickName:  sellerDto.Nickname,
		LegalName: sellerDto.CorporateName,
		Document:  sellerDto.Document,
		Active:    &sellerDto.Active,
		Address:   address,
	}

	err = sellerEntity.ValidateFieldsToCreate()
	if err != nil {
		return nil, err
	}

	return &sellerEntity, nil
}

func (s *Seller) IsAccountIDEqual(accountID string) bool {
	return pkg.UUIDToString(s.AccountID) == accountID
}

// NewSellerToUpdate creates or updates an existing Seller entity from the provided SellerInputUpdateDTO.
// Returns a pointer to the updated Seller and an error if any validations fail.
func SellerUpdate(sellerDto dto.SellerInputUpdateDTO) (*Seller, error) {
	err, address := UpdateAddress(sellerDto.Address)
	if err != nil {
		return nil, err
	}

	accountID, _ := pkg.StringToUUID(sellerDto.AccountID)
	sellerID, _ := pkg.StringToUUID(sellerDto.ID)
	sellerEntity := Seller{
		ID:        sellerID,
		AccountID: accountID,
		NickName:  sellerDto.Nickname,
		LegalName: sellerDto.CorporateName,
		Document:  sellerDto.Document,
		Active:    &sellerDto.Active,
		CreatedAt: sellerDto.CreatedAt,
		UpdatedAt: time.Now(),
		Address:   address,
	}

	err = sellerEntity.ValidateFieldsToCreate()
	if err != nil {
		return nil, err
	}

	return &sellerEntity, nil
}
