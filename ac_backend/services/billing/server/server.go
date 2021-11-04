package server

import (
	"alphacare/ac_backend/logging"
	"alphacare/ac_backend/proto/generated/billing"
	pb "alphacare/ac_backend/proto/generated/billing"
	"alphacare/ac_backend/services/billing/stripe"
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
	pb.BillingServer
	db            *gorm.DB
	billingClient billing.BillingClient
	config        *viper.Viper
	stripeAdapter *stripe.StripeAdapter
}

func NewServer(database *gorm.DB, config *viper.Viper) *Server {
	return &Server{
		db:            database,
		config:        config,
		stripeAdapter: stripe.NewStripeAdapter(),
	}
}
