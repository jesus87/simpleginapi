package store

import "github.com/banwire/api-exam/models"

type TransactionRepository struct {
	transactions []models.Transaction
}
                           
func (t *TransactionRepository) Create(tr models.Transaction, merchantCommission float32) error {
	tr.Commission = merchantCommission
	t.transactions = append(t.transactions, tr)

	return nil
}

func (t *TransactionRepository) GetAll() ([]models.Transaction, error) {

	return t.transactions, nil
}

func (t *TransactionRepository) GetRevenue() float32 {
	total := float32(0)
	for _, item := range t.transactions {
		total += item.Commission
	}

	return total
}

func (t *TransactionRepository) GetRevenueByMerchant() map[int]float32 {
	result := make(map[int]float32)
	for _, item := range t.transactions {
		if val, exists := result[item.MerchantId]; exists {
			result[item.MerchantId] += val
		} else {
			result[item.MerchantId] = val
		}
	}

	return result
}
