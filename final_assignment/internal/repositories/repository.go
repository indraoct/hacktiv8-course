package repositories

import (
	"context"
	"hacktiv8-course/final_assignment/commons/options"
	"hacktiv8-course/final_assignment/internal/entities"
)

type Repository struct {
	options.Options
}

type RepositoryImpl interface {
	CreateUser(ctx context.Context, user entities.User) error
	GetUserByID(ctx context.Context, id uint) (*entities.User, error)
	GetAllUsers(ctx context.Context) ([]entities.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
	UpdateUser(ctx context.Context, user *entities.User) error
	DeleteUser(ctx context.Context, id uint) error

	CreatePhoto(ctx context.Context, photo *entities.Photo) error
	GetPhotoByID(ctx context.Context, id uint) (*entities.Photo, error)
	UpdatePhoto(ctx context.Context, photo *entities.Photo) error
	DeletePhoto(ctx context.Context, id uint) error
	GetAllPhotos(ctx context.Context) ([]entities.Photo, error)

	CreateComment(ctx context.Context, comment *entities.Comment) error
	GetCommentByID(ctx context.Context, id uint) (*entities.Comment, error)
	UpdateComment(ctx context.Context, comment *entities.Comment) error
	GetAllComments(ctx context.Context) ([]entities.Comment, error)
	DeleteComment(ctx context.Context, id uint) error

	CreateSocialMedia(ctx context.Context, entry *entities.SocialMedia) error
	GetAllSocialMedia(ctx context.Context) ([]entities.SocialMedia, error)
	GetSocialMediaByID(ctx context.Context, id uint) (*entities.SocialMedia, error)
	UpdateSocialMedia(ctx context.Context, entry *entities.SocialMedia) error
	DeleteSocialMedia(ctx context.Context, id uint) error
}

func NewRepositories(opt options.Options) RepositoryImpl {
	return Repository{opt}
}
