package pkg

func IsEmptyUUID(id ID) bool {
	UUID, _ := ParseID("00000000-0000-0000-0000-000000000000")
	return id == UUID
}
