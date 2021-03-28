package utils

import (
	"github.com/gofrs/uuid"
)

func GenUUID() string{
	v4, _ := uuid.NewV4()
	return v4.String()
}