package models

import "time"

type Transaction struct {
	TransactionId int
	MerchantId    int
	Amount        float32
	Commission    float32
	Fee           float32
	Items         []string
	Created_at    time.Time
}
