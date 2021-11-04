package server

import (
	pb "alphacare/ac_backend/proto/generated/billing"
	"context"
)

func (s *Server) ValidateCoupon(ctx context.Context, req *pb.ValidateCouponRequest) (*pb.ValidateCouponResponse, error) {
	log.Infof("ValidateCoupon(): request = %v", req)
	resp := &pb.ValidateCouponResponse{}
	coupon, err := s.stripeAdapter.ReadCoupon(req.CouponId)
	if err != nil {
		errMsg := "failed to read coupon"
		log.Infof("ValidateCoupon(): %s, error = %+v", errMsg, err)
		return resp, err
	}
	resp.IsCouponValid = true
	resp.AmountOff = coupon.AmountOff
	resp.PercentOff = coupon.PercentOff

	return resp, nil
}
