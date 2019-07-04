package domain

import (
	"github.com/google/uuid"
	"strings"
)

func GenerateUUID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
