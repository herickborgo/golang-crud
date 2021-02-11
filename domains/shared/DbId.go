package shared

import "github.com/google/uuid"

// DbID type
type DbID struct {
	value uuid.UUID
	err   error
}

func () New() DbID {
	dbID := DbID{value: uuid.NewUUID}
	return dbID
}
