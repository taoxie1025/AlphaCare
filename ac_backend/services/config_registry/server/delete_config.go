package server

import (
	pb "alphacare/ac_backend/proto/generated/config_registry"
	"context"
)

func (s *Server) DeleteConfig(ctx context.Context, req *pb.DeleteConfigRequest) (*pb.DeleteConfigResponse, error) {
	log.Infof("POST DeleteConfig: %v", req)
	res := &pb.DeleteConfigResponse{}
	result := s.db.Table(ConfigTable).Delete(req.UserId)
	if result.Error != nil {
		log.Info("Failed to delete user config")
		return res, result.Error
	}
	return res, nil
}
