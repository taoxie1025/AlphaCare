package server

import (
	pb "alphacare/ac_backend/proto/generated/user_registry"
	"alphacare/ac_backend/services/user_registry/lib"
	"alphacare/ac_backend/services/user_registry/model"
	"context"
	"errors"
	"strings"
	"time"
)

var (
	InvalidCredsErr = errors.New("Password or Email invalid")
)

func (s *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	log.Infof("POST LoginUser: %v", req)
	resp := &pb.LoginUserResponse{}

	tmpUser := &model.User{}
	result := s.db.Table(UserTable).Where("email = ?", strings.ToLower(req.Email)).Find(tmpUser)
	if result.Error != nil {
		return resp, result.Error
	}
	resp.UserId = tmpUser.Id
	if !lib.CompareHashedStrings(tmpUser.SaltedPassword, req.Password) {
		return resp, InvalidCredsErr
	}

	log.Infof("UserID: %v logged in", tmpUser.Id)
	jwtToken, err := lib.GenJWT(
		s.config.GetString("app.jwtKey"),
		time.Duration(s.config.GetInt64("app.jwtExpiryInMinutes"))*(time.Minute),
		tmpUser)
	if err != nil {
		log.Errorf("Error creating jwt token: %v", err)
		return resp, err
	}

	claims := &pb.Claims{
		Email:     req.GetEmail(),
		UserScope: int32(tmpUser.UserScope),
		Exp:       time.Now().Add(time.Duration(s.config.GetInt64("app.jwtExpiryInMinutes"))*(time.Minute)).Unix(),
		Iat:       time.Now().Unix(),
	}
	resp.JwtToken = jwtToken
	resp.Claims = claims
	return resp, nil
}
