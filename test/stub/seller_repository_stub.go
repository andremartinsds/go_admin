package stub

import (
	"github.com/andremartinsds/go_admin/internal/entities"
)

type SellerRepositoryStub struct {
	ErrorFnCreate     error
	ErrorFnExists     error
	ErrorFnUpdate     error
	ConditionFnExists bool
	ErrorFnSelect     error
	Seller            entities.Seller
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
