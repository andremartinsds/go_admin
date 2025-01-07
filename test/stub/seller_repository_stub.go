package stub

import (
	"github.com/andremartinsds/go_admin/internal/entities"
)

type SellerRepositoryStub struct {
	ErrorFnCreate        error
	ErrorFnExists        error
	ErrorFnUpdate        error
	ConditionFnExists    bool
	ErrorFnSelect        error
	Seller               entities.Seller
	ErrorFnDeleteByID    error
	ListSellerFn         []*entities.Seller
	ErrorFnListSeller    error
	SelectOneByIDFN      *entities.Seller
	ErrorFnSelectOneByID error
}

func (s *SellerRepositoryStub) List(accountID string) ([]*entities.Seller, error) {
	return s.ListSellerFn, s.ErrorFnListSeller
}

func (s *SellerRepositoryStub) Exists(param map[string]string) (bool, error) {
	return s.ConditionFnExists, s.ErrorFnExists
}

func (s *SellerRepositoryStub) Update(seller *entities.Seller) error {
	return s.ErrorFnUpdate
}

func (s *SellerRepositoryStub) Select(param map[string]string) (*entities.Seller, error) {
	return &s.Seller, s.ErrorFnSelect
}

func (s *SellerRepositoryStub) Create(seller *entities.Seller) error {
	return s.ErrorFnCreate
}
func (s *SellerRepositoryStub) DeleteById(seller entities.Seller) error {
	return s.ErrorFnDeleteByID
}

func (s *SellerRepositoryStub) SelectOneById(id string) (*entities.Seller, error) {
	return s.SelectOneByIDFN, s.ErrorFnSelectOneByID
}
