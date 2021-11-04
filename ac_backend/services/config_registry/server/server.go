package server

import (
	"alphacare/ac_backend/logging"
	pb "alphacare/ac_backend/proto/generated/config_registry"
	"alphacare/ac_backend/proto/generated/user_registry"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

var (
	log = logging.NewZapLogger()
)

const (
	ConfigTable = "config_table"
)

type Server struct {
	pb.ConfigRegistryServer
	db                 *gorm.DB
	userRegistryClient user_registry.UserRegistryClient
}

func NewServer(database *gorm.DB, urConn *grpc.ClientConn) *Server {
	return &Server{
		db:                 database,
		userRegistryClient: user_registry.NewUserRegistryClient(urConn),
	}
}
