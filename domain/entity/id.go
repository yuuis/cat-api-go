package entity

import "github.com/google/uuid"


func GenerateUUID() string {
	return uuid.New().String()
}
