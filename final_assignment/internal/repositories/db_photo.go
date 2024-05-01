package repositories

import (
	"context"
	"hacktiv8-course/final_assignment/internal/entities"
)

// CreatePhoto creates a new photo record in the database
func (r Repository) CreatePhoto(ctx context.Context, photo *entities.Photo) error {
	result := r.DbGorm.Create(photo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetPhotoByID retrieves a photo record from the database by ID
func (r Repository) GetPhotoByID(ctx context.Context, id uint) (*entities.Photo, error) {
	var photo entities.Photo
	result := r.DbGorm.First(&photo, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &photo, nil
}

// UpdatePhoto updates an existing photo record in the database
func (r Repository) UpdatePhoto(ctx context.Context, photo *entities.Photo) error {
	result := r.DbGorm.Save(photo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeletePhoto deletes a photo record from the database by ID
func (r Repository) DeletePhoto(ctx context.Context, id uint) error {
	result := r.DbGorm.Delete(&entities.Photo{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAllPhotos retrieves all photo records from the database
func (r Repository) GetAllPhotos(ctx context.Context) ([]entities.Photo, error) {
	var photos []entities.Photo
	result := r.DbGorm.Find(&photos)
	if result.Error != nil {
		return nil, result.Error
	}
	return photos, nil
}
