package repositories

import (
	"context"
	"hacktiv8-course/final_assignment/internal/entities"
)

// CreateComment creates a new comment record in the database
func (r Repository) CreateComment(ctx context.Context, comment *entities.Comment) error {
	result := r.DbGorm.Create(comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetCommentByID retrieves a comment record from the database by ID
func (r Repository) GetCommentByID(ctx context.Context, id uint) (*entities.Comment, error) {
	var comment entities.Comment
	result := r.DbGorm.First(&comment, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &comment, nil
}

// UpdateComment updates an existing comment record in the database
func (r Repository) UpdateComment(ctx context.Context, comment *entities.Comment) error {
	result := r.DbGorm.Save(comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAllComments retrieves all comment records from the database
func (r Repository) GetAllComments(ctx context.Context) ([]entities.Comment, error) {
	var comments []entities.Comment
	result := r.DbGorm.Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}

// DeleteComment deletes a comment record from the database by ID
func (r Repository) DeleteComment(ctx context.Context, id uint) error {
	result := r.DbGorm.Delete(&entities.Comment{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r Repository) GetCommentByUserID(ctx context.Context, userId uint) (*[]entities.Comment, error) {
	var comments []entities.Comment
	result := r.DbGorm.Find(&comments, "user_id=?", userId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &comments, nil
}

func (r Repository) GetCommentByIDByUserId(ctx context.Context, userId, id uint) (*entities.Comment, error) {
	var comment entities.Comment
	result := r.DbGorm.First(&comment, "user_id=? AND id=?", userId, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &comment, nil
}
