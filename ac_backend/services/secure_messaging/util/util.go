package util

import (
	guuid "github.com/google/uuid"
)

func GenerateSecureMessageUUID() string {
	return "secure-message-" + guuid.New().String()[:14]
}
