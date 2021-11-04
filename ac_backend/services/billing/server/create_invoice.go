package server

import (
	pb "alphacare/ac_backend/proto/generated/billing"
	"context"
	"github.com/spf13/viper"
)

func (s *Server) CreateInvoice(ctx context.Context, req *pb.CreateInvoiceRequest) (*pb.CreateInvoiceResponse, error) {
	log.Infof("CreateInvoice(): request = %v", req)
	resp := &pb.CreateInvoiceResponse{}
	inv, err := s.stripeAdapter.CreateInvoice(req.CustomerId, viper.GetString("app.stripe.generalInvoiceProductId"), req.Amount)
	if err != nil {
		errMsg := "failed to read coupon"
		log.Infof("CreateInvoice(): %s, error = %+v", errMsg, err)
		return resp, err
	}
	resp.InvoiceId = inv.ID
	resp.InvoicePdfUrl = inv.InvoicePDF
	resp.PaymentUrl = inv.InvoicePDF[:len(inv.InvoicePDF)-4]
	resp.CustomerId = inv.Customer.ID
	resp.CustomerEmail = inv.CustomerEmail

	return resp, nil
}
