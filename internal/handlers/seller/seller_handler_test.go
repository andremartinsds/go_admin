package seller

import (
	"bytes"
	"encoding/json"
	"github.com/andremartinsds/go_admin/pkg/tt"
	"github.com/andremartinsds/go_admin/test/dummy"
	"github.com/andremartinsds/go_admin/test/stub"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSellerHandler_CreateSeller(t *testing.T) {
	t.Run("should create a new seller", func(t *testing.T) {
		//Arrange
		sellerInstance := Instance(
			&stub.SellerRepositoryStub{
				ErrorFnCreate: nil,
			}, &stub.AccountRepositoryStub{
				ErrorFnExists: nil,
				ErrorFnCreate: nil,
				Condition:     true,
			})

		sellerDto, err := json.Marshal(dummy.CreateInputSellerDTODummy())
		if err != nil {
			t.Error(err)
		}

		//Act
		req, err := http.NewRequest("POST", "/sellers", bytes.NewBuffer(sellerDto))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		sellerInstance.CreateSeller(rr, req)

		//Assert
		tt.AssertTest(t, rr.Code, http.StatusCreated)
	})

	t.Run("should return 400 bad request when account does not exists", func(t *testing.T) {
		//Arrange
		sellerInstance := Instance(
			&stub.SellerRepositoryStub{
				ErrorFnCreate: nil,
			}, &stub.AccountRepositoryStub{
				ErrorFnExists: nil,
				ErrorFnCreate: nil,
				Condition:     false,
			})

		sellerDto, err := json.Marshal(dummy.CreateInputSellerDTODummy())
		if err != nil {
			t.Error(err)
		}

		req, err := http.NewRequest("POST", "/sellers", bytes.NewBuffer(sellerDto))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		//Act
		sellerInstance.CreateSeller(rr, req)

		//Assert
		tt.AssertTest(t, rr.Code, http.StatusBadRequest)
	})

	t.Run("should return 400 bad request when seller already exists", func(t *testing.T) {
		//Arrange
		sellerInstance := Instance(
			&stub.SellerRepositoryStub{
				ErrorFnCreate:     nil,
				ErrorFnExists:     nil,
				ConditionFnExists: true,
			},
			&stub.AccountRepositoryStub{
				ErrorFnCreate: nil,
				ErrorFnExists: nil,
				Condition:     true,
			})

		sellerDto, err := json.Marshal(dummy.CreateInputSellerDTODummy())
		if err != nil {
			t.Error(err)
		}

		req, err := http.NewRequest("POST", "/sellers", bytes.NewBuffer(sellerDto))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		//Act
		sellerInstance.CreateSeller(rr, req)

		//Assert
		tt.AssertTest(t, rr.Code, http.StatusBadRequest)
	})
}
