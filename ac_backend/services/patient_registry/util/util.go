package util

import (
	guuid "github.com/google/uuid"
)

func GeneratePatientUUID() string {
	return "patient-" + guuid.New().String()[:8]
}

func GenerateInitialPassword() string {
	return guuid.New().String()[:4]
}
