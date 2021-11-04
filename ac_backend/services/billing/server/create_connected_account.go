package server

import (
	pb "alphacare/ac_backend/proto/generated/billing"
	"context"
)

// Create connected account for the medical providers so that they can collect money from their clients.
// When a new medical provider registers on the platform, this should be called internally to update the user_table in db.
// https://stripe.com/docs/connect/enable-payment-acceptance-guide
func (s *Server) CreateConnectedAccount(ctx context.Context, req *pb.CreateConnectedAccountRequest) (*pb.CreateConnectedAccountResponse, error) {
	log.Infof("CreateConnectedAccount(): request = %v", req)
	resp := &pb.CreateConnectedAccountResponse{}
	acct, acctLink, err := s.stripeAdapter.CreateConnectedAccount(req.Email, req.RefreshUrl, req.ReturnUrl)
	if err != nil {
		errMsg := "failed to create customer"
		log.Infof("CreateConnectedAccount(): %s, error = %+v", errMsg, err)
		return resp, err
	}
	resp.AccountId = acct.ID
	resp.Email = acct.Email
	resp.CreatedAt = acctLink.Created
	resp.ExpiresAt = acctLink.ExpiresAt
	resp.OnboardingUrl = acctLink.URL

	return resp, nil
}
