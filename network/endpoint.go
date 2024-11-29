package network

const (
	baseURL = "https://antrpay.com/api/api"

	PaymentP2PURL = baseURL + "/repayment/create_payment_to_card"
	PaymentFPSURL = baseURL + "/repayment/create_payment_fps_h2h"
	PayoutURL     = baseURL + "/repayment/create_payout"
)
