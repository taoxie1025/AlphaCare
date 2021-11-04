package server

import (
	"alphacare/ac_backend/logging"
	pb "alphacare/ac_backend/proto/generated/notification"
	"context"
	"errors"
	"github.com/spf13/viper"
)

var (
	log = logging.NewZapLogger()
)

func (s *Server) SendEmail(ctx context.Context, in *pb.SendEmailRequest) (*pb.SendEmailResponse, error) {
	log.Info("SendEmail(): Received: %v", in)
	fromEmail := viper.GetString("app.ses.defaultOutgoingEmail")
	switch in.GetType() {
	case pb.EmailType_EmailTypeBilling:
		fromEmail = viper.GetString("app.ses.billingOutgoingEmail")
	case pb.EmailType_EmailTypeAccount:
		fromEmail = viper.GetString("app.ses.accountOutgoingEmail")
	}

	if yes, err := s.SesAdapter.IsEmailVerified(fromEmail); err == nil {
		if !yes {
			return nil, errors.New("fromAddress is not verified")
		}
	} else {
		log.Error("SendEmail(): failed to validate fromAddress error = %v")
		return nil, err
	}

	for _, toEmail := range in.GetTo() {
		err := s.SesAdapter.SendEmail(fromEmail, toEmail, in.GetSubject(), in.GetHtmlBody())
		if err != nil {
			log.Errorf("SendEmail(): failed to send email from %s to %s, error = %v", fromEmail, toEmail, err)
			return nil, err
		}
	}

	return &pb.SendEmailResponse{}, nil
}
