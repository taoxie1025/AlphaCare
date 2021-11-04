package server

import (
	pb "alphacare/ac_backend/proto/generated/notification"
	"alphacare/ac_backend/services/notification/mail"
)

type Server struct {
	pb.NotificationServer
	SesAdapter *mail.SesAdapter
}

func NewServer() *Server {
	return &Server{
		SesAdapter: mail.NewSesAdapter(),
	}
}
