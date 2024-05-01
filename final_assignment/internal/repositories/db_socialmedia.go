package repositories

import (
	"context"
	"hacktiv8-course/final_assignment/internal/entities"
)

// CreateSocialMedia creates a new social media entry in the database
func (r Repository) CreateSocialMedia(ctx context.Context, entry *entities.SocialMedia) error {
	result := r.DbGorm.Create(entry)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAllSocialMedia retrieves all social media entries from the database
func (r Repository) GetAllSocialMedia(ctx context.Context) ([]entities.SocialMedia, error) {
	var entries []entities.SocialMedia
	result := r.DbGorm.Find(&entries)
	if result.Error != nil {
		return nil, result.Error
	}
	return entries, nil
}

// GetSocialMediaByID retrieves a social media entry from the database by ID
func (r Repository) GetSocialMediaByID(ctx context.Context, id uint) (*entities.SocialMedia, error) {
	var entry entities.SocialMedia
	result := r.DbGorm.First(&entry, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entry, nil
}

// UpdateSocialMedia updates an existing social media entry in the database
func (r Repository) UpdateSocialMedia(ctx context.Context, entry *entities.SocialMedia) error {
	result := r.DbGorm.Save(entry)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteSocialMedia deletes a social media entry from the database by ID
func (r Repository) DeleteSocialMedia(ctx context.Context, id uint) error {
	result := r.DbGorm.Delete(&entities.SocialMedia{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
