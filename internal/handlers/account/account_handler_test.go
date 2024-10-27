package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/andremartinsds/go_admin/internal/entities"
	"github.com/andremartinsds/go_admin/pkg/tt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andremartinsds/go_admin/internal/dto"

	"github.com/andremartinsds/go_admin/test/dummy"
	"github.com/andremartinsds/go_admin/test/stub"
)

func TestCreateAccount(t *testing.T) {
	t.Run("should create new account", func(t *testing.T) {
		//Arange
		accountHandlerInstance := AccountHandlerInstancy(&stub.AccountRepositoryStub{
			ErrorFnCreate: nil,
			ErrorFnExists: nil,
			Condition:     false,
		})

		createInputAccountDTO := dummy.CreateInputAccountDTODummy()
		body, err := json.Marshal(createInputAccountDTO)
		if err != nil {
			t.Error(err)
		}

		//Act
		req, err := http.NewRequest("POST", "/accounts", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal("failed to create request")
		}
		rr := httptest.NewRecorder()

		accountHandlerInstance.CreateAccount(rr, req)

		//Assert
		if status := rr.Code; status != http.StatusCreated {
			t.Error("handler returned wrong status code")
		}
		var returnedAccount entities.Account
		err = json.NewDecoder(rr.Body).Decode(&returnedAccount)
		if err != nil {
			t.Fatalf("failed to decode response body: %v", err)
		}
		tt.AssertTest(t, returnedAccount.Name, createInputAccountDTO.Name)
		tt.AssertTest(t, returnedAccount.Document, createInputAccountDTO.Document)
	})

	t.Run("should return error if account already exists", func(t *testing.T) {
		//Arrange
		handler := AccountHandlerInstancy(&stub.AccountRepositoryStub{
			ErrorFnCreate: nil,
			ErrorFnExists: errors.New("the account already exists"),
			Condition:     true,
		})
		createAccountDTO := dummy.CreateInputAccountDTODummy()

		body, err := json.Marshal(createAccountDTO)
		if err != nil {
			t.Error(err)
		}
		req, err := http.NewRequest("POST", "/accounts", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal("failed to create request")
		}
		rr := httptest.NewRecorder()

		//Act
		handler.CreateAccount(rr, req)

		//Assert
		if status := rr.Code; status != http.StatusBadRequest {
			fmt.Println(status)
			t.Error("the status code should have been http.StatusBadRequest")
		}
	})

	t.Run("should return an account dto like a response", func(t *testing.T) {
		//Arrange
		want := dto.AccountOutputDto{
			Nickname:    "Name Company",
			Document:    "11111111111",
			Active:      true,
			AccountType: "PJ",
		}

		accountHandlerInstance := AccountHandlerInstancy(&stub.AccountRepositoryStub{
			ErrorFnCreate: nil,
			ErrorFnExists: nil,
			Condition:     false,
		})

		createAccountDto := dummy.CreateInputAccountDTODummy()

		body, err := json.Marshal(createAccountDto)
		if err != nil {
			t.Fatal("failed to create request")
		}
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("POST", "/accounts", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal("failed to handle request")
		}

		//Act
		accountHandlerInstance.CreateAccount(rr, req)

		var got dto.AccountOutputDto
		tt.UnmarshalTest(t, rr, &got)

		want.Id = got.Id

		//Assert
		tt.AssertTest(t, got, want)
	})
}

func TestSelectOneById(t *testing.T) {
	t.Run("should select one account by id", func(t *testing.T) {
		//Arrange
		handler := AccountHandlerInstancy(&stub.AccountRepositoryStub{
			AccountEntityFnSelectOneById: *dummy.AccountEntitityDummy(),
		})

		req, _ := http.NewRequest("GET", "/accounts/1", nil)
		rr := httptest.NewRecorder()

		//Act
		handler.SelectAccount(rr, req)

		tt.AssertTest(t, rr.Code, http.StatusOK)
	})

	t.Run("should return error if account not found", func(t *testing.T) {
		//Arrange
		handler := AccountHandlerInstancy(&stub.AccountRepositoryStub{
			ErrorFnSelectOneById: errors.New("account not found"),
		})

		req, _ := http.NewRequest("GET", "/accounts/22", nil)

		rr := httptest.NewRecorder()

		//Act
		handler.SelectAccount(rr, req)

		//Assert
		tt.AssertTest(t, rr.Code, http.StatusNotFound)
	})
}

func TestUpdateOneById(t *testing.T) {
	t.Run("if account don't exists should return 404", func(t *testing.T) {
		//Arrange
		handler := AccountHandlerInstancy(&stub.AccountRepositoryStub{
			ErrorFnUpdateOne:     nil,
			ErrorFnSelectOneById: errors.New("account not found"),
		})

		accountToUpdateDTO := dummy.UpdateInputAccountDTODummy()

		account, err := json.Marshal(accountToUpdateDTO)
		if err != nil {
			t.Error(err)
		}
		req, _ := http.NewRequest("PUT", "/accounts/1", bytes.NewBuffer(account))

		//Act
		rr := httptest.NewRecorder()
		handler.UpdateAccount(rr, req)

		//Assert
		tt.AssertTest(t, rr.Code, http.StatusNotFound)
	})

	t.Run("should update one account by id", func(t *testing.T) {
		//Arrange
		handler := AccountHandlerInstancy(&stub.AccountRepositoryStub{
			ErrorFnUpdateOne:     nil,
			ErrorFnSelectOneById: nil,
		})

		accountToUpdateDTO := dummy.UpdateInputAccountDTODummy()

		account, err := json.Marshal(accountToUpdateDTO)
		if err != nil {
			t.Error(err)
		}
		req, _ := http.NewRequest("PUT", "/accounts/1", bytes.NewBuffer(account))

		//Act
		rr := httptest.NewRecorder()
		handler.UpdateAccount(rr, req)

		//Assert
		tt.AssertTest(t, rr.Code, http.StatusOK)
	})
}
