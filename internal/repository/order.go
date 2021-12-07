package repository

import (
	"context"
	"gorm.io/gorm"
	"prj/internal/entity"
)

type OrderRepository interface {
	Get(ctx context.Context, orderID int32) (interface{}, error)
}

type orderRepository struct {
	db *gorm.DB
}

func CreateOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (o *orderRepository) Get(ctx context.Context, orderID int32) (interface{}, error) {
	return o.db.First(&entity.Order{}, orderID), nil
}
