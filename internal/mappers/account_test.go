package mappers

import (
	"github.com/andremartinsds/go_admin/internal/dto"
	"github.com/andremartinsds/go_admin/test/dummy"
	"reflect"
	"testing"
)

func TestAccountMapper(t *testing.T) {
	t.Run("should convert accountEntity to accountOutputDto", func(t *testing.T) {
		accountDummy := dummy.AccountEntitityDummy()

		accountMapper := ToAccountOutputDTO(accountDummy)

		if reflect.TypeOf(*accountMapper) != reflect.TypeOf(dto.AccountOutputDto{}) {
			t.Error("Account Mapper error")
		}
	})
}
