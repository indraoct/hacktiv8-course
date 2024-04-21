package usecases

import (
	"context"
	"hacktiv8-course/assignment2/commons/options"
	"hacktiv8-course/assignment2/internal/entities"
	"hacktiv8-course/assignment2/internal/repositories"
	"net/http"
)

type Usecases struct {
	opt  options.Options
	repo repositories.RepositoryImpl
}

type UsecaseImpl interface {
	CreateOrder(ctx context.Context, orders entities.Orders) (message string, httpCode int, err error)
}

func NewUsecase(opt options.Options, repo repositories.RepositoryImpl) UsecaseImpl {
	return Usecases{opt: opt, repo: repo}
}

func (u Usecases) CreateOrder(ctx context.Context, orders entities.Orders) (message string, httpCode int, err error) {

	var (
		reqOrder entities.Orders
		reqItems entities.Items
	)

	httpCode = http.StatusCreated
	message = "Success!"

	defer func() {
		if err != nil {
			httpCode = http.StatusBadRequest
			message = err.Error()
		}
	}()

	reqOrder = orders

	for _, item := range orders.Items {
		reqItems = item
		break
	}

	err = u.repo.CreateOrderAndItem(ctx, reqOrder, reqItems)
	return
}

func (u Usecases) GetAllOrderAndItem(ctx context.Context) (orders []entities.Orders, message string, httpCode int, err error) {

	httpCode = http.StatusCreated
	message = "Success!"

	defer func() {
		if err != nil {
			httpCode = http.StatusBadRequest
			message = err.Error()
		}
	}()

	orders, err = u.repo.GetAllOrders(ctx)
	return
}

func (u Usecases) UpdateOrder(ctx context.Context, orders entities.Orders) (message string, httpCode int, err error) {

	var (
		reqOrder entities.Orders
		reqItems entities.Items
	)

	httpCode = http.StatusCreated
	message = "Success!"

	defer func() {
		if err != nil {
			httpCode = http.StatusBadRequest
			message = err.Error()
		}
	}()

	reqOrder = orders

	for _, item := range orders.Items {
		reqItems = item
		break
	}

	err = u.repo.UpdateOrderAndItem(ctx, reqOrder, reqItems)
	return
}

func (u Usecases) DeleteOrder(ctx context.Context, orders entities.Orders) (message string, httpCode int, err error) {

	var (
		reqOrder entities.Orders
		reqItems entities.Items
	)

	httpCode = http.StatusCreated
	message = "Success!"

	defer func() {
		if err != nil {
			httpCode = http.StatusBadRequest
			message = err.Error()
		}
	}()

	reqOrder = orders

	for _, item := range orders.Items {
		reqItems = item
		break
	}

	err = u.repo.UpdateOrderAndItem(ctx, reqOrder, reqItems)
	return
}
