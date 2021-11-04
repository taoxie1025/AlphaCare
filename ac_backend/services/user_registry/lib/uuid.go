package lib

import (
	guuid "github.com/google/uuid"
)

func GenerateUserUUID() string {
	return "User-" + guuid.New().String()[:8]
}
