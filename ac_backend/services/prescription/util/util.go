package util

import (
	guuid "github.com/google/uuid"
)

func GeneratePrescriptionUUID() string {
	return "prescription-" + guuid.New().String()[:14]
}
