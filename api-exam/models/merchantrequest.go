package models

import "time"

type MerchantRequest struct {
	MerchantID   int
	MerchantName string
	Commission   float32
}

func (m MerchantRequest) MapCreate() Merchant {
	val := Merchant{
		MerchantID:   m.MerchantID,
		MerchantName: m.MerchantName,
		Commission:   m.Commission,
	}

	val.CreatedAt = time.Now()

	return val
}
func (m MerchantRequest) MapUpdate() Merchant {
	val := Merchant{
		MerchantID:   m.MerchantID,
		MerchantName: m.MerchantName,
		Commission:   m.Commission,
	}

	val.UpdatedAt = time.Now()

	return val
}
