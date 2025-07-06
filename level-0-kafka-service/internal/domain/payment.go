package domain

type Payment struct {
	OrderUID     string `json:"-" db:"order_uid"`
	Transaction  string `json:"transaction" db:"transaction"`
	RequestID    string `json:"request_id" db:"request_id"`
	Currency     string `json:"currency" db:"currency"`
	Provider     string `json:"provider" db:"provider"`
	Amount       int64  `json:"amount" db:"amount"`
	PaymentDt    int64  `json:"payment_dt" db:"payment_dt"`
	Bank         string `json:"bank" db:"bank"`
	DeliveryCost int64  `json:"delivery_cost" db:"delivery_cost"`
	GoodsTotal   int64  `json:"goods_total" db:"goods_total"`
	CustomFee    int64  `json:"custom_fee" db:"custom_fee"`
}
