package repository

import (
	"receipt-processor-challenge/model"

	"github.com/google/uuid"
)

var (
	receiptsTable = make(map[string]model.Receipt)
	pointsTable   = make(map[string]int64)
)

type ReceiptRepo struct{}

func (s *ReceiptRepo) Store(receipt model.Receipt) string {
	id := s.GenerateUUID()
	receiptsTable[id] = receipt
	return id
}

func (s *ReceiptRepo) StorePoints(id string, points int64) error {
	if _, exists := receiptsTable[id]; !exists {
		// no receipt exists in the system for this id
		return model.ErrReceiptNotFound
	}
	pointsTable[id] = points

	return nil
}

func (s *ReceiptRepo) GetByID(id string) (model.Receipt, error) {
	result, ok := receiptsTable[id]
	if !ok {
		return model.Receipt{}, model.ErrReceiptNotFound
	}

	return result, nil
}

func (s *ReceiptRepo) GetPointsByID(id string) (int64, error) {
	result, ok := pointsTable[id]
	if !ok {
		return 0, model.ErrReceiptNotFound
	}

	return result, nil
}

func (s *ReceiptRepo) GenerateUUID() string {
	return uuid.New().String()
}
