package factory

import (
	"fmt"
	"gorm.io/gorm"
	"prj/internal/entity"
	pb "prj/proto/order"
)

type OrderFactory interface {
	Create(orderData *pb.OrderPayload) (*entity.Order, error)
}

type orderFactory struct {
	db *gorm.DB
}

func CreateOrderFactory(db *gorm.DB) OrderFactory {
	return &orderFactory{
		db: db,
	}
}

func (o orderFactory) Create(orderData *pb.OrderPayload) (*entity.Order, error) {
	order := &entity.Order{
		Name:     orderData.Name,
		StatusId: int(orderData.Status),
		Price:    orderData.Price,
	}
	result := o.db.Create(order)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
	}
	return order, nil
}
