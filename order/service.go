package order

import (
	"dbo-be/entities"
	"dbo-be/helper"
	"errors"
	"time"
)

type Service interface {
	SearchOrders(input OrderSearchInput) (*helper.Pagination, error)
	GetOrders() (*helper.Pagination, error)
	GetOrderById(ID string) (entities.Order, error)
	CreateOrder(input OrderInput) (entities.Order, error)
	EditOrder(ID string, input OrderInput) (entities.Order, error)
	DeleteOrder(ID string) (entities.Order, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) SearchOrders(input OrderSearchInput) (*helper.Pagination, error) {
	order, err := s.repository.SearchOrder(input)
	if err != nil {
		return order, err
	}

	return order, nil
}

func (s *service) GetOrders() (*helper.Pagination, error) {
	orders, err := s.repository.GetOrder()

	if err != nil {
		return nil, err
	}

	return orders, err
}

func (s *service) GetOrderById(ID string) (entities.Order, error) {
	order, err := s.repository.FindOrderByID(ID)
	if err != nil {
		return order, err
	}

	if order.ID == 0 {
		return order, errors.New("data transaksi tidak ditemukan berdasarkan ID")
	}

	return order, nil
}

func (s *service) CreateOrder(input OrderInput) (entities.Order, error) {
	var order entities.Order

	order.UserID = input.UserID
	order.Name = input.Name
	order.Type = input.Type
	order.Quantity = input.Quantity
	order.Price = input.Price
	order.CreatedAt = time.Now()

	newTransaction, err := s.repository.StoreOrder(order)

	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}

func (s *service) EditOrder(ID string, input OrderInput) (entities.Order, error) {
	order, err := s.repository.UpdateOrder(ID, input)
	if err != nil {
		return order, err
	}

	return order, nil
}

func (s *service) DeleteOrder(ID string) (entities.Order, error) {
	order, err := s.repository.DestroyOrder(ID)
	if err != nil {
		return order, err
	}

	return order, nil
}
