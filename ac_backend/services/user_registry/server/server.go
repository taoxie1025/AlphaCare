package server

import (
	"alphacare/ac_backend/logging"
	"alphacare/ac_backend/proto/generated/billing"
	pb "alphacare/ac_backend/proto/generated/user_registry"
	"google.golang.org/grpc"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	log = logging.NewZapLogger()
)

const (
	UserTable = "user_table"
)

type Server struct {
	pb.UserRegistryServer
	db            *gorm.DB
	config        *viper.Viper
	billingClient billing.BillingClient
}

func NewServer(database *gorm.DB, config *viper.Viper, nginxConn *grpc.ClientConn) *Server {
	return &Server{
		db:            database,
		config:        config,
		billingClient: billing.NewBillingClient(nginxConn),
	}
}
