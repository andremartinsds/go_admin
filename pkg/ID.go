package pkg

import (
	"github.com/google/uuid"
)

type ID = uuid.UUID

func NewUUID() ID {
	return ID(uuid.New())
}
func StringToUUID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}
func UUIDToString(id ID) string {
	return id.String()
}

func ValidUUID(s string) bool {
	return len(s) == SizeGuid
}
