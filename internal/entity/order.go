package entity

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID          int
	Name        string      `gorm:"column:name;type:varchar(155)"`
	Price       uint32      `gorm:"column:price;type:int"`
	StatusId    int         `gorm:"column:order_status_id"`
	OrderStatus OrderStatus `gorm:"foreignKey:StatusId"`
}

type OrderStatus struct {
	gorm.Model
	ID   int
	Slug string `gorm:"column:slug;type:varchar(155)"`
}
