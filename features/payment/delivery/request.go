package delivery

type Request struct {
	Bank  string `json:"bank" form:"bank" validate:"required"`
	Kurir string `json:"kurir" form:"kurir" validate:"required"`
}

type MidtransHookRequest struct {
	TransactionTime   string `form:"transaction_time" json:"transaction_time"`
	TransactionStatus string `form:"transaction_status" json:"transaction_status"`
	OrderID           string `form:"order_id" json:"order_id"`
	MerchantID        string `form:"merchant_id" json:"merchant_id"`
	GrossAmount       string `form:"gross_amount" json:"gross_amount"`
	FraudStatus       string `form:"fraud_status" json:"fraud_status"`
	Currency          string `form:"currency" json:"currency"`
}
