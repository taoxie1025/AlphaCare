package main

import (
	"alphacare/ac_backend/logging"
	pb "alphacare/ac_backend/proto/generated/notification"
	"alphacare/ac_backend/services/notification/server"
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
)

var (
	log = logging.NewZapLogger()
)

func main() {
	LoadConfig()
	lis, err := net.Listen("tcp", viper.GetString("app.server.port"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterNotificationServer(s, server.NewServer())
	log.Infof("notification svc listening to port: %v", viper.GetString("app.server.port"))
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
