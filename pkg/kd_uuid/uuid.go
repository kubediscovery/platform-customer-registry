package kd_uuid

import (
	"github.com/google/uuid"
)

func NewUuid() uuid.UUID {
	return uuid.New()
}

func NewUuidString() string {
	u, _ := uuid.NewV7()
	return u.String()
}

func ParseUuid(id string) (uuid.UUID, error) {
	return uuid.Parse(id)
}

func IsValid(id *string) bool {
	u, err := uuid.Parse(*id)

	if err != nil {
		return false
	}

	if u.String() != *id {
		return false
	}

	return true
}
