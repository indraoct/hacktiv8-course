package usecases

import (
	"context"
	"errors"
	"github.com/astaxie/beego/validation"
	"hacktiv8-course/final_assignment/commons/helper"
	"hacktiv8-course/final_assignment/commons/security"
	"hacktiv8-course/final_assignment/internal/entities"
	"net/http"
	"time"
)

func (u Usecases) Register(ctx context.Context, req entities.Request) (httpStatus int, err error) {
	var (
		user entities.User
	)

	defer func() {
		helper.RecoverPanic()

		httpStatus = http.StatusCreated
		if err != nil {
			httpStatus = http.StatusBadRequest
		}
	}()

	//general field validation
	valid := validation.Validation{}
	valid.Required(req.Email, "email")
	valid.Email(req.Email, "email")
	valid.Required(req.Password, "password")
	valid.Required(req.Username, "username")
	valid.Required(req.Age, "age")
	valid.MinSize(req.Password, 6, "password")
	valid.Min(req.Age, 8, "age")

	if valid.HasErrors() {
		for _, errValidate := range valid.Errors {
			err = errors.New(errValidate.Key + ": " + errValidate.Message)
			return
		}
	}

	users, err := u.repo.GetAllUsers(ctx)
	if err != nil {
		return
	}

	for _, usr := range users {
		if usr.Email == req.Email || usr.Username == req.Username {
			err = errors.New("user: " + req.Email + ", " + req.Username + " is already exist")
			return
		}
	}

	//hashing the password
	passwordHash, err := security.BcryptHashPassword(req.Password)
	if err != nil {
		return
	}

	user = entities.User{
		Username:  req.Username,
		Title:     req.Title,
		Email:     req.Email,
		Password:  passwordHash,
		Age:       req.Age,
		CreatedAt: time.Now(),
	}

	err = u.repo.CreateUser(ctx, user)

	return
}

func (u Usecases) Login(ctx context.Context, req entities.Request) (httpStatus int, err error) {

	return
}
