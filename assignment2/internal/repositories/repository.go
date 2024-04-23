package repositories

import (
	"context"
	"hacktiv8-course/assignment2/commons/options"
	"hacktiv8-course/assignment2/internal/entities"
	"time"
)

type Repository struct {
	options.Options
}

type RepositoryImpl interface {
	CreateOrderAndItem(ctx context.Context, reqOrder entities.Orders, reqItem entities.Items) error
	GetAllOrders(ctx context.Context) (respData []entities.Orders, err error)
	GetOrderById(ctx context.Context, id uint) (respData entities.Orders, err error)
	UpdateOrderAndItem(ctx context.Context, reqOrder entities.Orders) (err error)
	DeleteOrderAndItem(ctx context.Context, reqOrder entities.Orders) (err error)
}

func NewRepositories(opt options.Options) RepositoryImpl {
	return Repository{opt}
}

func (r Repository) CreateOrderAndItem(ctx context.Context, reqOrder entities.Orders, reqItem entities.Items) (err error) {

	var items []entities.Items
	reqOrder.OrderedAt = time.Now()
	reqOrder.CreatedAt = time.Now()
	reqOrder.OrderedAt = time.Now()

	for _, item := range reqOrder.Items {
		item.CreatedAt = time.Now()
		items = append(items, item)
	}
	reqOrder.Items = items

	err = r.DbGorm.Create(&reqOrder).Error
	if err != nil {
		return
	}

	return
}

func (r Repository) GetAllOrders(ctx context.Context) (respData []entities.Orders, err error) {

	var (
		reqData entities.Orders
	)

	err = r.DbGorm.Model(&reqData).Preload("Items").Find(&respData).Error
	if err != nil {
		return
	}
	return
}

func (r Repository) GetOrderById(ctx context.Context, id uint) (respData entities.Orders, err error) {

	var (
		reqData entities.Orders
	)

	err = r.DbGorm.Model(&reqData).Preload("Items").First(&respData, "order_id=?", id).Error
	if err != nil {
		return
	}
	return
}

func (r Repository) UpdateOrderAndItem(ctx context.Context, reqOrder entities.Orders) (err error) {

	//update order
	err = r.DbGorm.Save(&reqOrder).Error
	if err != nil {
		return
	}

	//update items
	for _, item := range reqOrder.Items {
		if err = r.DbGorm.Save(&item).Error; err != nil {
			return
		}
	}

	return
}

func (r Repository) DeleteOrderAndItem(ctx context.Context, reqOrder entities.Orders) (err error) {

	//delete all item
	orders, _ := r.GetOrderById(ctx, reqOrder.OrderId)
	for _, item := range orders.Items {
		err = r.DbGorm.Delete(&item).Error
		if err != nil {
			return
		}
	}

	//delete order
	err = r.DbGorm.Delete(&reqOrder).Error
	if err != nil {
		return
	}

	return
}
