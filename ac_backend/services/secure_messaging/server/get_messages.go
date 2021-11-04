package server

import (
	pb "alphacare/ac_backend/proto/generated/secure_messaging"
	"context"
	"strings"
)

func (s *Server) GetMessage(ctx context.Context, in *pb.GetMessageRequest) (*pb.GetMessageResponse, error) {
	log.Info("GetMessage for receiver Id: %v", in.GetReceiverId())

	res := &pb.GetMessageResponse{}
	var messages []*pb.Message
	receiver := in.GetReceiverId()
	if receiver == "" {
		log.Info("Receiver can not be null")
		return res, nil
	}

	result := s.db.Table(SecureMessagingTable).Where("receiver = ?", strings.ToLower(receiver)).Find(messages)
	if result.Error != nil {
		return res, result.Error
	}
	return &pb.GetMessageResponse{
		Messages: messages,
	}, nil
}
