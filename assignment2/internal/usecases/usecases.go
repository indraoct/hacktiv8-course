package usecases

import (
	"hacktiv8-course/assignment2/commons/options"
	"hacktiv8-course/assignment2/internal/repositories"
)

type Usecases struct {
	opt  options.Options
	repo repositories.RepositoryImpl
}

type UsecaseImpl interface {
}

func NewUsecase(opt options.Options, repo repositories.RepositoryImpl) UsecaseImpl {
	return Usecases{opt: opt, repo: repo}
}

func (u Usecases) Test() {

}
