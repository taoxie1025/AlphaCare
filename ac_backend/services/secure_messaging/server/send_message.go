package server

import (
	"alphacare/ac_backend/logging"
	pb "alphacare/ac_backend/proto/generated/secure_messaging"
	"alphacare/ac_backend/services/secure_messaging/util"
	user_model "alphacare/ac_backend/services/user_registry/model"
	"context"
	"strings"
	"time"
)

var (
	log = logging.NewZapLogger()
)

func (s *Server) SendMessage(ctx context.Context, in *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	log.Info("SendMessage from user Id: %v to user Id: %v", in.GetMessage().GetSender(), in.GetMessage().GetReceiver())

	res := &pb.SendMessageResponse{}
	sender := in.GetMessage().GetSender()
	receiver := in.GetMessage().GetReceiver()
	if sender == "" {
		log.Info("Sender can not be null")
		return res, nil
	}
	if receiver == "" {
		log.Info("Receiver can not be null")
		return res, nil
	}
	tmpUser := &user_model.User{}
	result := s.db.Table(UserTable).Where("id = ?", strings.ToLower(sender)).Find(tmpUser)
	if result.Error != nil {
		log.Info("Sender %v does not exist", sender)
		return res, result.Error
	}
	result = s.db.Table(UserTable).Where("id = ?", strings.ToLower(receiver)).Find(tmpUser)
	if result.Error != nil {
		log.Info("Receiver %v does not exist", receiver)
		return res, result.Error
	}
	message := in.GetMessage()
	message.Id = util.GenerateSecureMessageUUID()
	message.Timestamp = time.Now().UnixNano() / int64(time.Millisecond)
	result = s.db.Table(SecureMessagingTable).Create(in.GetMessage())
	if result.Error != nil {
		log.Info("Failed to send message")
		return res, result.Error
	}
	return res, nil
}
