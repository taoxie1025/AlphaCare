package server

import (
	"alphacare/ac_backend/proto/generated/billing"
	pb "alphacare/ac_backend/proto/generated/user_registry"
	"alphacare/ac_backend/services/user_registry/lib"
	"alphacare/ac_backend/services/user_registry/model"
	"context"
	"errors"
	"regexp"
	"strings"
	"time"
)

var (
	emailPattern = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

func validateEmail(email string) bool {
	return emailPattern.MatchString(email)
}

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	log.Infof("POST CreateUser: %v", req)
	res := &pb.CreateUserResponse{}
	if !validateEmail(req.Email) {
		log.Info("Malformed email")
		return res, errors.New("Malformed email string")
	}
	hashedPassword, err := lib.SaltString(req.Password)
	if err != nil {
		log.Info("Failed to hash users password")
		return res, err
	}

	// created stripe connected account
	acctReq := &billing.CreateConnectedAccountRequest{
		Email:      req.Email,
		RefreshUrl: req.RefreshUrl,
		ReturnUrl:  req.ReturnUrl,
	}
	acct, err := s.billingClient.CreateConnectedAccount(ctx, acctReq)
	if err != nil {
		log.Infof("Failed to create new stripe connected account, err = %v", err)
		return nil, err
	}
	custReq := &billing.CreateCustomerRequest{
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	// created stripe customer
	cust, err := s.billingClient.CreateCustomer(ctx, custReq)
	if err != nil {
		log.Infof("Failed to create new stripe customer, err = %v", err)
		return nil, err
	}

	tmpUser := &model.User{
		Id:                           lib.GenerateUserUUID(),
		UserScope:                    model.USER_SCOPE_DOCTOR,
		Firstname:                    strings.ToLower(req.FirstName),
		Lastname:                     strings.ToLower(req.LastName),
		Email:                        strings.ToLower(req.Email),
		SaltedPassword:               hashedPassword,
		ReturnUrl:                    req.ReturnUrl,
		RefreshUrl:                   req.RefreshUrl,
		StripeAccountCreatedAt:       acct.CreatedAt,
		StripeOnboardingUrl:          acct.OnboardingUrl,
		StripeOnboardingUrlExpiresAt: acct.ExpiresAt,
		StripeAccountId:              acct.AccountId,
		StripeCustomerId:             cust.CustomerId,
	}
	result := s.db.Table(UserTable).Create(tmpUser)
	if result.Error != nil {
		log.Info("Failed to create new user")
		return res, result.Error
	}

	res.UserId = tmpUser.Id
	res.StripeAccountId = acct.AccountId
	res.StripeAccountCreatedAt = acct.CreatedAt
	res.StripeOnboardingUrl = acct.OnboardingUrl
	res.StripeOnboardingUrlExpiresAt = acct.ExpiresAt
	res.StripeCustomerId = cust.CustomerId
	res.RefreshUrl = req.RefreshUrl
	res.ReturnUrl = req.ReturnUrl

	jwtToken, err := lib.GenJWT(
		s.config.GetString("app.jwtKey"),
		time.Duration(s.config.GetInt64("app.jwtExpiryInMinutes"))*(time.Minute),
		tmpUser)
	if err != nil {
		log.Errorf("Error creating jwt token: %v", err)
	}

	claims := &pb.Claims{
		Email:     req.GetEmail(),
		UserScope: int32(tmpUser.UserScope),
		Exp:       time.Now().Add(time.Duration(s.config.GetInt64("app.jwtExpiryInMinutes"))*(time.Minute)).Unix(),
		Iat:       time.Now().Unix(),
	}
	res.JwtToken = jwtToken
	res.Claims = claims
	return res, nil
}
