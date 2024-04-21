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
	UpdateOrderAndItem(ctx context.Context, order entities.Orders, items entities.Items) (err error)
	DeleteOrderAndItem(ctx context.Context, order entities.Orders, items entities.Items) (err error)
}

func NewRepositories(opt options.Options) RepositoryImpl {
	return Repository{opt}
}

func (r Repository) CreateOrderAndItem(ctx context.Context, reqOrder entities.Orders, reqItem entities.Items) (err error) {

	//start transaction
	r.DbGorm.Begin()

	defer func() {
		//if there is error then rollback
		if err == nil {
			r.DbGorm.Rollback()
			return
		}
		//if everything is ok then will be transaction commit
		r.DbGorm.Commit()

	}()

	reqOrder.OrderedAt = time.Now()
	reqOrder.CreatedAt = time.Now()
	reqOrder.OrderedAt = time.Now()

	err = r.DbGorm.Create(&reqOrder).Error
	if err != nil {
		return
	}

	err = r.DbGorm.Create(&reqItem).Error
	return
}

func (r Repository) GetAllOrders(ctx context.Context) (respData []entities.Orders, err error) {

	var (
		reqData entities.Orders
	)

	rows, err := r.DbGorm.Model(&reqData).Rows()
	if err != nil {
		return
	}

	for rows.Next() {
		var order entities.Orders
		r.DbGorm.ScanRows(rows, &order)
		respData = append(respData, order)
	}

	return
}

func (r Repository) UpdateOrderAndItem(ctx context.Context, order entities.Orders, items entities.Items) (err error) {

	defer func() {
		//if there is error then rollback
		if err != nil {
			r.DbGorm.Rollback()
			return
		}
		//if everything is ok then will be transaction commit
		r.DbGorm.Commit()
	}()

	r.DbGorm.Begin()
	order.UpdatedAt = time.Now()

	if err = r.DbGorm.Where("order_id = ?", order.OrderId).Updates(&order).Error; err != nil {
		return
	}

	if err = r.DbGorm.Where("item_id = ?", items.ItemId).Updates(&items).Error; err != nil {
		return
	}

	return
}

func (r Repository) DeleteOrderAndItem(ctx context.Context, order entities.Orders, items entities.Items) (err error) {
	defer func() {
		//if there is error then rollback
		if err != nil {
			r.DbGorm.Rollback()
			return
		}
		//if everything is ok then will be transaction commit
		r.DbGorm.Commit()
	}()

	r.DbGorm.Begin()
	if err = r.DbGorm.Where("order_id = ?", order.OrderId).Delete(&order).Error; err != nil {
		return
	}

	if err = r.DbGorm.Where("item_id = ?", items.ItemId).Delete(&items).Error; err != nil {
		return
	}
	return
}
