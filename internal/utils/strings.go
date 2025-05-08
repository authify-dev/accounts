package utils

import (
	"fmt"

	"github.com/google/uuid"
)

func GenerateRandomUserName() string {
	id := uuid.New()

	return fmt.Sprintf("User_%s", id.String())
}
