package main

import (
	"alphacare/ac_backend/logging"
	pb "alphacare/ac_backend/proto/generated/patient_registry"
	"alphacare/ac_backend/services/patient_registry/server"
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
	pb.RegisterPatientRegistryServiceServer(s, server.NewServer(
		getReversedProxyConn(),
	))
	log.Infof("patient registry svc listening to port: %v", viper.GetString("app.server.port"))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
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

func LoadConfig() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
