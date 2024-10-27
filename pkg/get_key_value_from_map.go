package pkg

import (
	"errors"
)

func GetKeyValueFromMap(m map[string]string) (string, string, error) {
	limit := 0
	var value string
	var key string
	var err error
	for k, v := range m {
		limit++
		if limit > 1 {
			err = errors.New("limit exceeded")
			break
		}
		key = k
		value = v
	}
	return key, value, err

}
