package server

import (
	pb "alphacare/ac_backend/proto/generated/prescription"
	"gorm.io/gorm"
)

const (
	PrescriptionTable = "prescription"
)

type Server struct {
	pb.PrescriptionServiceServer
	db *gorm.DB
}

func NewServer(database *gorm.DB) *Server {
	return &Server{
		db: database,
	}
}
