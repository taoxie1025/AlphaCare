package stripe

import (
	"alphacare/ac_backend/logging"
	"github.com/spf13/viper"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/account"
	"github.com/stripe/stripe-go/v72/accountlink"
	"github.com/stripe/stripe-go/v72/checkout/session"
	"github.com/stripe/stripe-go/v72/coupon"
	"github.com/stripe/stripe-go/v72/customer"
	"github.com/stripe/stripe-go/v72/invoice"
	"github.com/stripe/stripe-go/v72/invoiceitem"
	"github.com/stripe/stripe-go/v72/sub"
	"github.com/stripe/stripe-go/v72/webhook"
	"strconv"
	"time"
)

var (
	log = logging.NewZapLogger()
)

type StripeAdapter struct {
	apiKey        string
	webhookSecret string
}

func NewStripeAdapter() *StripeAdapter {
	apiKey := viper.GetString("app.stripe.stripePrivateKey")
	stripe.Key = apiKey
	stripeAdapter := &StripeAdapter{
		apiKey: apiKey,
	}
	stripeAdapter.CreateWebhook()

	return stripeAdapter
}

func (d *StripeAdapter) CreateCustomer(email, firstName, lastName string) (*stripe.Customer, error) {
	log.Infof("CreateCustomer(): email = %s", email)
	params := &stripe.CustomerParams{
		Email: stripe.String(email),
		Name:  stripe.String(firstName + " " + lastName),
	}
	customer, err := customer.New(params)
	if err != nil {
		log.Errorf("CreateCustomer(): error = %v", err)
	}
	return customer, err
}

func (d *StripeAdapter) CreateSubscriptionSession(customerId, priceId, plan, email, coupon, successURL, cancelURL string) (*stripe.CheckoutSession, error) {
	log.Infof("CreateSubscriptionSession(): customerId = %s, priceId = %s, email = %s", customerId, priceId, email)
	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		Mode: stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		SubscriptionData: &stripe.CheckoutSessionSubscriptionDataParams{
			Items: []*stripe.CheckoutSessionSubscriptionDataItemsParams{
				{
					Plan:     stripe.String(priceId),
					Quantity: stripe.Int64(1),
				},
			},
			Metadata: map[string]string{
				"plan":    plan,
				"priceId": priceId,
				"email":   email,
			},
		},
		ClientReferenceID: stripe.String(email + "," + priceId + "," + plan + "," + strconv.FormatInt(time.Now().Unix(), 10)),
		SuccessURL:        stripe.String(successURL),
		CancelURL:         stripe.String(cancelURL),
	}
	params.CustomerEmail = stripe.String(email)

	if coupon != "" {
		params.SubscriptionData.Coupon = stripe.String(coupon)
	}

	checkoutSession, err := session.New(params)
	if err != nil {
		log.Errorf("CreateSubscriptionSession(): error = %v", err)
	}
	return checkoutSession, err
}

func (d *StripeAdapter) CreateWebhook() {
	webhookSecret := viper.GetString("app.stripe.stripeWebhookSecret")
	d.webhookSecret = webhookSecret

	/* create webhook dynamically
	log.Infof("CreateWebhook(): %s%s", httpaddr, "/api/v1/webhook")
	params := &stripe.WebhookEndpointParams{
		URL:           stripe.String(httpaddr + "/api/v1/webhook"),
		EnabledEvents: stripe.StringSlice([]string{"*"}),
		Description:   stripe.String("For user accounts' subscription plan only."),
	}

	endpoint, err := webhookendpoint.New(params)
	if err != nil {
		log.Errorf("CreateWebhook(): error = %v", err)
	}
	d.webhookSecret = endpoint.Secret
	*/
}

func (d *StripeAdapter) ConstructStripeEvent(reqBody []byte, header string) (*stripe.Event, error) {
	log.Infof("ConstructStripeEvent():")
	event, err := webhook.ConstructEvent(reqBody, header, d.webhookSecret)
	if err != nil {
		log.Errorf("ConstructStripeEvent(): error = %v", err)
		return nil, err
	}
	return &event, nil
}

func (d *StripeAdapter) CancelSubscription(subscriptionId string) error {
	log.Infof("CancelSubscription(): subscriptionId = %s", subscriptionId)
	_, err := sub.Cancel(subscriptionId, nil)
	if err != nil {
		log.Errorf("CancelSubscription(): error = %v", err)
	}
	return err
}

func (d *StripeAdapter) UpdateSubscription(subscriptionId, priceId string) error {
	log.Infof("UpdateSubscription(): subscriptionId = %s, priceId = %s", subscriptionId, priceId)
	existingSubscription, err := sub.Get(subscriptionId, nil)
	param := &stripe.SubscriptionParams{
		Items: []*stripe.SubscriptionItemsParams{
			{
				ID:    stripe.String(existingSubscription.Items.Data[0].ID),
				Price: stripe.String(priceId),
			},
		},
		ProrationBehavior: stripe.String(string(stripe.SubscriptionProrationBehaviorCreateProrations)),
		CancelAtPeriodEnd: stripe.Bool(false),
	}
	_, err = sub.Update(subscriptionId, param)
	if err != nil {
		log.Errorf("UpdateSubscription(): error = %v", err)
	}
	return err
}

func (d *StripeAdapter) ReadCoupon(couponId string) (*stripe.Coupon, error) {
	log.Infof("ReadCoupon(%s)", couponId)
	coupon, err := coupon.Get(couponId, nil)
	if err != nil {
		log.Warnf("ReadCoupon(): error = %v", err)
	}
	return coupon, err
}

func (d *StripeAdapter) CreateInvoice(customerId, productId string, amount float64) (*stripe.Invoice, error) {
	log.Infof("CreateInvoice(): customerId = %s, productId = %s, amount = %f", customerId, productId, amount)
	params := &stripe.InvoiceItemParams{
		Customer: stripe.String(customerId),
		PriceData: &stripe.InvoiceItemPriceDataParams{
			Currency:          stripe.String("usd"),
			Product:           stripe.String(productId),
			UnitAmountDecimal: stripe.Float64(amount * 100), // convert cents to dollars, times 100.
		},
	}
	ii, err := invoiceitem.New(params)
	if err != nil {
		log.Error("CreateInvoice(): failed to create invoice item, err = %v", err)
		return nil, err
	}

	invParams := &stripe.InvoiceParams{
		Customer:    stripe.String(ii.Customer.ID),
		AutoAdvance: stripe.Bool(false),
	}
	inv, err := invoice.New(invParams)
	if err != nil {
		log.Error("CreateInvoice(): failed to create invoice, err = %v", err)
		return nil, err
	}

	finalInv, err := invoice.FinalizeInvoice(inv.ID, nil)
	if err != nil {
		log.Error("CreateInvoice(): failed to finalize invoice, err = %v", err)
		return nil, err
	}
	return finalInv, nil
}

func (d *StripeAdapter) CreateConnectedAccount(email, refreshUrl, returnUrl string) (*stripe.Account, *stripe.AccountLink, error) {
	log.Infof("CreateConnectedAccount():")
	acctParams := &stripe.AccountParams{
		Type:  stripe.String(string(stripe.AccountTypeStandard)),
		Email: stripe.String(email),
	}
	acct, err := account.New(acctParams)
	if err != nil {
		log.Error("CreateConnectedAccount(): failed to create new account, err = %v", err)
		return nil, nil, err
	}

	acctLinkParams := &stripe.AccountLinkParams{
		Account:    stripe.String(acct.ID),
		RefreshURL: stripe.String(refreshUrl),
		ReturnURL:  stripe.String(returnUrl),
		Type:       stripe.String("account_onboarding"),
	}
	acctLink, err := accountlink.New(acctLinkParams)
	if err != nil {
		log.Error("CreateConnectedAccount(): failed to create account link, err = %v", err)
		return nil, nil, err
	}
	return acct, acctLink, nil
}
