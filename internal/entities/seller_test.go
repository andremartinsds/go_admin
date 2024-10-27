package entities

import (
	"testing"

	"github.com/andremartinsds/go_admin/test/mock"
)

func TestCriarUmSeller(t *testing.T) {
	//Arrange
	sellerDto := mock.SellerInputCreateDtoMock()
	// Act
	seller, err := NewSeller(sellerDto)
	// Assert
	if err != nil {
		t.Errorf("Erro ao criar um Seller: %v", err)
	} else {
		t.Logf("Seller criado com sucesso: %v", seller)
	}
}

func TestCriarUmSellerComAccountId(t *testing.T) {
	//Arrange
	sellerDto := mock.SellerInputCreateDtoMock()
	sellerDto.AccountID = ""
	exp := "[seller.accountId is required]"
	//Act
	seller, err := NewSeller(sellerDto)
	//Assert
	if err.Error() != exp {
		t.Errorf("Seller Criado indevidamente: %v", seller)
	}

}
