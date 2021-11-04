package main

import (
	"alphacare/ac_backend/logging"
	pb "alphacare/ac_backend/proto/generated/prescription"
	"alphacare/ac_backend/services/prescription/model"
	"alphacare/ac_backend/services/prescription/server"
	"fmt"
	"net"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	log = logging.NewZapLogger()
)

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

	db.Table(server.PrescriptionTable).AutoMigrate(&model.Prescription{})
	return db
}

func main() {
	LoadConfig()
	lis, err := net.Listen("tcp", viper.GetString("app.server.port"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPrescriptionServiceServer(s, server.NewServer(
		getDBConn(),
	))
	log.Infof("Prescription svc listening to port: %v", viper.GetString("app.server.port"))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func LoadConfig() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
