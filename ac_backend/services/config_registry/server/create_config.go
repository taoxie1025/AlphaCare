package server

import (
	pb "alphacare/ac_backend/proto/generated/config_registry"
	"alphacare/ac_backend/services/config_registry/model"
	"context"
	"strconv"
)

func (s *Server) CreateConfig(ctx context.Context, req *pb.CreateConfigRequest) (*pb.CreateConfigResponse, error) {
	log.Infof("CreateConfig: %v", req)
	res := &pb.CreateConfigResponse{}
	config := &model.Config{
		InstalledPlugins: []model.Plugin{},
	}
	result := s.db.Table(ConfigTable).Create(config)
	if result.Error != nil {
		log.Info("Failed to create user config")
		return res, result.Error
	}

	res.Id = strconv.Itoa(int(config.ID))
	return res, nil
}
