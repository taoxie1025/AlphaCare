package server

import (
	pb "alphacare/ac_backend/proto/generated/patient_visit"
	"gorm.io/gorm"
)

const (
	PatientVisitTable = "patient_visit"
	UserTable         = "user_table"
	PatientTable      = "patient_table"
)

type Server struct {
	pb.PatientVisitServiceServer
	db *gorm.DB
}

func NewServer(database *gorm.DB) *Server {
	return &Server{
		db: database,
	}
}
