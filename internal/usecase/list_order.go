package usecase

import "github.com/devfullcycle/20-CleanArch/internal/entity"

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(repository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: repository,
	}
}

func (uc *ListOrdersUseCase) Execute() ([]OrderOutputDTO, error) {
	orders, err := uc.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var orderOutputDTOS []OrderOutputDTO
	for _, order := range orders {
		orderDto := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		orderOutputDTOS = append(orderOutputDTOS, orderDto)
	}
	return orderOutputDTOS, nil
}
