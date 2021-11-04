package server

import (
	pb "alphacare/ac_backend/proto/generated/patient_registry"
	"alphacare/ac_backend/proto/generated/user_registry"
	user_registry_service "alphacare/ac_backend/proto/generated/user_registry"
	"alphacare/ac_backend/services/patient_registry/model"
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	PatientTable = "patient_table"
)

type Server struct {
	pb.PatientRegistryServiceServer
	db                 *gorm.DB
	userRegistryClient user_registry.UserRegistryClient
}

func NewServer(reversedProxyConn *grpc.ClientConn) *Server {
	return &Server{
		db:                 getDBConn(),
		userRegistryClient: user_registry_service.NewUserRegistryClient(reversedProxyConn),
	}
}

func getDBConn() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		viper.GetString("app.db.url"),
		viper.GetString("app.db.user"),
		viper.GetString("app.db.password"),
		viper.GetString("app.db.name"),
		viper.GetString("app.db.port"),
	)

	log.Infof("connecting to db: %s", viper.GetString("app.db.url"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if there is an error opening the connection, handle it
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	db.Table(PatientTable).AutoMigrate(&model.Patient{})
	return db
}
