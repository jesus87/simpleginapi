package models

import "time"

//Merchant merchant
type Merchant struct {
	MerchantID   int
	MerchantName string
	Commission   float32
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
