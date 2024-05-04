package usecases

import (
	"context"
	"hacktiv8-course/final_assignment/commons/options"
	"hacktiv8-course/final_assignment/internal/entities"
	"hacktiv8-course/final_assignment/internal/repositories"
)

type Usecases struct {
	opt  options.Options
	repo repositories.RepositoryImpl
}

type UsecaseImpl interface {
	Register(ctx context.Context, req entities.Request) (httpStatus int, err error)
	Login(ctx context.Context, req entities.Request) (token string, httpStatus int, err error)

	CreatePhoto(ctx context.Context, req entities.Request) (httpStatus int, err error)
	GetAllPhoto(ctx context.Context) (photos *[]entities.Photo, httpStatus int, err error)
	GetPhotoById(ctx context.Context, req entities.Request) (photo *entities.Photo, httpStatus int, err error)
	UpdatePhoto(ctx context.Context, req entities.Request) (photo *entities.Photo, httpStatus int, err error)
	DeletePhoto(ctx context.Context, req entities.Request) (httpStatus int, err error)

	CreateComment(ctx context.Context, req entities.Request) (httpStatus int, err error)
	GetAllComment(ctx context.Context) (comments *[]entities.Comment, httpStatus int, err error)
	GetCommentById(ctx context.Context, req entities.Request) (comment *entities.Comment, httpStatus int, err error)
	UpdateComment(ctx context.Context, req entities.Request) (comment *entities.Comment, httpStatus int, err error)
	DeleteComment(ctx context.Context, req entities.Request) (httpStatus int, err error)
}

func NewUsecase(opt options.Options, repo repositories.RepositoryImpl) UsecaseImpl {
	return Usecases{opt: opt, repo: repo}
}
