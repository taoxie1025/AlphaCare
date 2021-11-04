package main

import (
	"alphacare/ac_backend/logging"
	pb "alphacare/ac_backend/proto/generated/user_registry"
	"alphacare/ac_backend/services/user_registry/model"
	"alphacare/ac_backend/services/user_registry/server"
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net"
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

	db.Table(server.UserTable).AutoMigrate(&model.User{})
	return db
}

func main() {
	LoadConfig()
	lis, err := net.Listen("tcp", viper.GetString("app.server.port"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserRegistryServer(s, server.NewServer(
		getDBConn(),
		viper.GetViper(),
		getReversedProxyConn(),
	))
	log.Infof("user registry svc listening to port: %v", viper.GetString("app.server.port"))
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

// keep the connection open till service shut down, or close the connection right away if dialing failed.
func getReversedProxyConn() *grpc.ClientConn {
	log.Infof("getReversedProxyConn(): dialing up nginx...")
	var conn *grpc.ClientConn
	var err error
	conn, err = grpc.Dial(viper.GetString("app.server.reversedProxy"), grpc.WithInsecure())
	if err != nil {
		log.Errorf("failed to make connection to nginx, err = %v", err)
		defer conn.Close()
		return nil
	}
	return conn
}