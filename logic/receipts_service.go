package logic

import "receipt-processor-challenge/domain"

type IReceiptsService interface {
	ByID(id uint) domain.Receipt
}

type ReceiptsService struct {
}

func NewReceiptsService() *ReceiptsService {
	return &ReceiptsService{}
}

func (s *ReceiptsService) ByID(id uint) domain.Receipt {
	return domain.Receipt{
		ID: id,
	}
}
