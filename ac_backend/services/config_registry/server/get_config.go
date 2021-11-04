package server

import (
	pb "alphacare/ac_backend/proto/generated/config_registry"
	"context"
)

func (s *Server) GetConfig(ctx context.Context, req *pb.GetConfigRequest) (*pb.GetConfigResponse, error) {
	log.Infof("POST GetConfig: %v", req)
	res := &pb.GetConfigResponse{}
	result := s.db.Table(ConfigTable).First(&res, req.UserId)
	if result.Error != nil {
		log.Info("Failed to get user config")
		return res, result.Error
	}
	return res, nil
}
