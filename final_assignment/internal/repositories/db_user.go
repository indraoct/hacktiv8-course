package repositories

import (
	"context"
	"hacktiv8-course/final_assignment/commons/helper"
	"hacktiv8-course/final_assignment/internal/entities"
)

// CreateUser creates a new user record in the database
func (r Repository) CreateUser(ctx context.Context, user entities.User) error {

	defer func() {
		helper.RecoverPanic()
	}()

	result := r.DbGorm.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetUserByID retrieves a user record from the database by ID
func (r Repository) GetUserByID(ctx context.Context, id uint) (*entities.User, error) {
	var user entities.User
	result := r.DbGorm.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// UpdateUser updates an existing user record in the database
func (r Repository) UpdateUser(ctx context.Context, user *entities.User) error {
	result := r.DbGorm.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteUser deletes a user record from the database by ID
func (r Repository) DeleteUser(ctx context.Context, id uint) error {
	result := r.DbGorm.Delete(&entities.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r Repository) GetAllUsers(ctx context.Context) ([]entities.User, error) {
	var user []entities.User
	result := r.DbGorm.Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
