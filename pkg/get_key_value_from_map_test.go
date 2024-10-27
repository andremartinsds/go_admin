package pkg

import (
	"github.com/andremartinsds/go_admin/pkg/tt"
	"testing"
)

func TestGetKeyValueFromMap(t *testing.T) {

	t.Run("should not receive more than one prop", func(t *testing.T) {
		_, _, err := GetKeyValueFromMap(map[string]string{
			"key1": "value1",
			"key2": "value2",
		})
		tt.AssertTest(t, err.Error(), "limit exceeded")

	})

	t.Run("should get the first prop", func(t *testing.T) {
		key, value, _ := GetKeyValueFromMap(map[string]string{
			"key1": "value1",
		})
		tt.AssertTest(t, key, "key1")
		tt.AssertTest(t, value, "value1")
	})

}
