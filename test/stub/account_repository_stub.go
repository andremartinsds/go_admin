package stub

import (
	"github.com/andremartinsds/go_admin/internal/entities"
)

type AccountRepositoryStub struct {
	ErrorFnCreate                error
	ErrorFnExists                error
	ErrorFnSelectOneById         error
	AccountEntityFnSelectOneById entities.Account
	Condition                    bool
	ErrorFnUpdateOne             error
	ErrorFnList                  error
	AccounntEntityFnList         []entities.Account
}

func (a *AccountRepositoryStub) Create(account *entities.Account) error {
	return a.ErrorFnCreate
}
func (a *AccountRepositoryStub) ExistsByField(param map[string]string) (bool, error) {
	return a.Condition, a.ErrorFnExists
}
func (a *AccountRepositoryStub) SelectOneById(id string) (*entities.Account, error) {
	return &a.AccountEntityFnSelectOneById, a.ErrorFnSelectOneById
}
func (a *AccountRepositoryStub) UpdateOne(account *entities.Account) error {
	return a.ErrorFnUpdateOne
}
func (a *AccountRepositoryStub) List() (*[]entities.Account, error) {
	return &a.AccounntEntityFnList, a.ErrorFnList
}
