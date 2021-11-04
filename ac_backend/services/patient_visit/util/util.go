package util

import (
	guuid "github.com/google/uuid"
)

func GeneratePatientVisitUUID() string {
	return "patient-visit-" + guuid.New().String()[:14]
}
