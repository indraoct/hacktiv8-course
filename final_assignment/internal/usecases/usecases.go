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
}

func NewUsecase(opt options.Options, repo repositories.RepositoryImpl) UsecaseImpl {
	return Usecases{opt: opt, repo: repo}
}
