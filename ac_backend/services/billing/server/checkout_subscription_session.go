package server

import (
	pb "alphacare/ac_backend/proto/generated/billing"
	"context"
)

func (s *Server) CheckoutSubscriptionSession(ctx context.Context, req *pb.CheckoutSubscriptionSessionRequest) (*pb.CheckoutSubscriptionSessionResponse, error) {
	log.Infof("CheckoutSubscriptionSession(): request = %v", req)
	sess, err := s.stripeAdapter.CreateSubscriptionSession(req.StripeCustomerId, req.PriceId, req.Plan, req.Email, req.Coupon, req.SuccessUrl, req.CancelUrl)
	if err != nil {
		errMsg := "failed to create checkout session"
		log.Errorf("%s, error = %+v", errMsg, err)
	}
	resp := &pb.CheckoutSubscriptionSessionResponse{
		SessionId: sess.ID,
	}
	return resp, nil
}
