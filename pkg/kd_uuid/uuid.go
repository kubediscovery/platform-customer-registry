package kd_uuid

import (
	"github.com/google/uuid"
)

func NewUuid() uuid.UUID {
	return uuid.New()
}

func NewUuidString() string {
	return uuid.New().String()
}

func ParseUuid(id string) (uuid.UUID ,error ){
	return uuid.Parse(id)
}
