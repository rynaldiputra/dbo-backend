package order

import (
	"dbo-be/entities"
	"dbo-be/helper"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	SearchOrder(input OrderSearchInput) (*helper.Pagination, error)
	GetOrder() (*helper.Pagination, error)
	FindOrderByID(ID string) (entities.Order, error)
	StoreOrder(order entities.Order) (entities.Order, error)
	UpdateOrder(ID string, input OrderInput) (entities.Order, error)
	DestroyOrder(ID string) (entities.Order, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) SearchOrder(input OrderSearchInput) (*helper.Pagination, error) {
	var orders []*entities.Order
	var pagination helper.Pagination
	pagination.Sort = "id ASC"
	condition := fmt.Sprintf("%s LIKE ?", input.Column)
	value := "%" + input.Value + "%"

	r.db.Scopes(helper.Paginate(orders, &pagination, r.db)).Where(condition, value).Find(&orders)

	pagination.Rows = orders

	return &pagination, nil
}

func (r *repository) GetOrder() (*helper.Pagination, error) {
	var orders []*entities.Order
	var pagination helper.Pagination
	pagination.Sort = "id ASC"

	r.db.Scopes(helper.Paginate(orders, &pagination, r.db)).Find(&orders)

	pagination.Rows = orders

	return &pagination, nil
}

func (r *repository) FindOrderByID(ID string) (entities.Order, error) {
	var order entities.Order

	err := r.db.Where("id = ?", ID).Find(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *repository) StoreOrder(order entities.Order) (entities.Order, error) {
	err := r.db.Create(&order).Error

	if err != nil {
		return order, err
	}
	return order, nil
}

func (r *repository) UpdateOrder(ID string, input OrderInput) (entities.Order, error) {
	var order entities.Order

	err := r.db.Where("id = ?", ID).First(&order).Error

	if err != nil {
		return order, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) == true {
		return order, errors.New("Data Order tidak ditemukan")
	}

	order.Name = input.Name
	order.Type = input.Type
	order.Quantity = input.Quantity
	order.Price = input.Price
	order.UpdatedAt = time.Now()

	err = r.db.Updates(&order).Error

	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *repository) DestroyOrder(ID string) (entities.Order, error) {
	var order entities.Order

	err := r.db.Where("id = ?", ID).First(&order).Error

	if err != nil {
		return order, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) == true {
		return order, errors.New("Data Order tidak ditemukan")
	}

	err = r.db.Delete(&order, ID).Error

	if err != nil {
		return order, err
	}

	return order, nil
}
