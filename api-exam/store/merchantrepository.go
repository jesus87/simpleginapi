package store

import (
	"fmt"
	"time"
                        
	"github.com/banwire/api-exam/models"
)                                              
type MerchantRepository struct {
	merchants []models.Merchant 
}
              
func (m *MerchantRepository) Create(mr models.Merchant) error {
	if item, err := m.Get(mr.MerchantID); item != nil {
		return fmt.Errorf("Id %v exists", mr.MerchantID)
	} else {
		if err != nil {
			return err
		}
	}

	if mr.Commission < 1 {
		mr.Commission = 1                                                                
	}
                                               
	if mr.Commission > 100 {
		mr.Commission = 100
	}

	m.merchants = append(m.merchants, mr)

	return nil
}                                                    
                                
func (m *MerchantRepository) Edit(mr models.Merchant) error {
	for i := 0; i < len(m.merchants); i++ {
		if m.merchants[i].MerchantID == mr.MerchantID {
			m.merchants[i].Commission = mr.Commission
			m.merchants[i].MerchantName = mr.MerchantName
			m.merchants[i].UpdatedAt = time.Now()
			return nil
		}
	}

	return fmt.Errorf("Merchant with ID: %v, not exists", mr.MerchantID)
}

func (m *MerchantRepository) Get(id int) (*models.Merchant, error) {
	for _, item := range m.merchants {
		if item.MerchantID == id {
			return &item, nil
		}
	}

	return nil, nil
}

func (m *MerchantRepository) GetAll() ([]models.Merchant, error) {

	return m.merchants, nil
}
