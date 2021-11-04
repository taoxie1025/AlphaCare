package server

import (
	pb "alphacare/ac_backend/proto/generated/secure_messaging"
	"gorm.io/gorm"
)

const (
	SecureMessagingTable = "secure_messaging"
	UserTable            = "user_table"
	PatientTable         = "patient_table"
)

type Server struct {
	pb.SecureMessagingServiceServer
	db *gorm.DB
}

func NewServer(database *gorm.DB) *Server {
	return &Server{
		db: database,
	}
}
