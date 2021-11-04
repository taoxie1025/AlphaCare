package server

import (
	pb "alphacare/ac_backend/proto/generated/user_registry"
	"alphacare/ac_backend/services/user_registry/model"
	"context"
	"strings"
)

func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	log.Infof("POST GetUSer: %v", req)
	resp := &pb.GetUserResponse{}

	tmpUser := &model.User{}
	result := s.db.Table(UserTable).Where("id = ?", strings.ToLower(req.UserId)).Find(tmpUser)
	if result.Error != nil {
		return resp, result.Error
	}
	return &pb.GetUserResponse{
		UserId:    tmpUser.Id,
		FirstName: tmpUser.Firstname,
		LastName:  tmpUser.Lastname,
		Email:     tmpUser.Email,
		UserScope: int32(tmpUser.UserScope),
	}, nil
}
