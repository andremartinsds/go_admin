package pkg

import (
	"github.com/google/uuid"
)

type ID = uuid.UUID

func NewUUID() ID {
	return ID(uuid.New())
}
func ParseID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}
func StrID(id ID) string {
	return id.String()
}
