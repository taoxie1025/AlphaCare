package server

import (
	pb "alphacare/ac_backend/proto/generated/billing"
	"context"
)

// Create a customer object in stripe to collect money from.
// A medical provider is also a customer to the platform, and they will need to pay monthly subscription using the customerId.
// A client of medical provider is also a customer, and they will need to be billed by the medical provider.
// When a medical provider and a client of the medical provider registers on the platform, this function should be called internally, and update the corresponding entry in db.
func (s *Server) CreateCustomer(ctx context.Context, req *pb.CreateCustomerRequest) (*pb.CreateCustomerResponse, error) {
	log.Infof("CreateCustomer(): request = %v", req)
	resp := &pb.CreateCustomerResponse{}
	customer, err := s.stripeAdapter.CreateCustomer(req.Email, req.FirstName, req.LastName)
	if err != nil {
		errMsg := "failed to create customer"
		log.Infof("CreateCustomer(): %s, error = %+v", errMsg, err)
		return resp, err
	}
	resp.CustomerId = customer.ID

	return resp, nil
}
