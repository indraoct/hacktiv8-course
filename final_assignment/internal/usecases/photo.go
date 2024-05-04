package usecases

import (
	"context"
	"errors"
	"github.com/astaxie/beego/validation"
	"hacktiv8-course/final_assignment/cmd/middleware"
	"hacktiv8-course/final_assignment/commons/helper"
	"hacktiv8-course/final_assignment/internal/entities"
	"net/http"
	"time"
)

func (u Usecases) CreatePhoto(ctx context.Context, req entities.Request) (httpStatus int, err error) {

	var (
		user  entities.User
		photo entities.Photo
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
	valid.Required(req.PhotoTitle, "photoTitle")
	valid.Required(req.PhotoUrl, "photoUrl")

	if valid.HasErrors() {
		for _, errValidate := range valid.Errors {
			err = errors.New(errValidate.Key + ": " + errValidate.Message)
			return
		}
	}

	user = middleware.GetAuth(ctx)
	photo = entities.Photo{
		Caption:   req.PhotoTitle,
		PhotoURL:  req.PhotoUrl,
		CreatedAt: time.Now(),
		UserID:    user.ID,
	}

	err = u.repo.CreatePhoto(ctx, &photo)

	return
}

func (u Usecases) GetAllPhoto(ctx context.Context) (photos *[]entities.Photo, httpStatus int, err error) {

	var (
		user entities.User
	)

	defer func() {
		helper.RecoverPanic()

		httpStatus = http.StatusOK
		if err != nil {
			httpStatus = http.StatusBadRequest
		}
	}()

	user = middleware.GetAuth(ctx)

	photos, err = u.repo.GetPhotoByUserID(ctx, user.ID)
	return
}

func (u Usecases) GetPhotoById(ctx context.Context, req entities.Request) (photo *entities.Photo, httpStatus int, err error) {

	var (
		user entities.User
	)

	defer func() {
		helper.RecoverPanic()

		httpStatus = http.StatusOK
		if err != nil {
			httpStatus = http.StatusBadRequest
		}
	}()

	valid := validation.Validation{}
	valid.Required(req.PhotoId, "photoId")

	if valid.HasErrors() {
		for _, errValidate := range valid.Errors {
			err = errors.New(errValidate.Key + ": " + errValidate.Message)
			return
		}
	}

	user = middleware.GetAuth(ctx)

	photo, err = u.repo.GetPhotoByIDAndUserID(ctx, user.ID, req.PhotoId)
	return
}

func (u Usecases) UpdatePhoto(ctx context.Context, req entities.Request) (photo *entities.Photo, httpStatus int, err error) {

	var (
		user entities.User
	)

	defer func() {
		helper.RecoverPanic()

		httpStatus = http.StatusOK
		if err != nil {
			httpStatus = http.StatusBadRequest
		}
	}()

	//general field validation
	valid := validation.Validation{}
	valid.Required(req.PhotoId, "photoId")
	valid.Required(req.PhotoTitle, "photoTitle")
	valid.Required(req.PhotoUrl, "photoUrl")

	if valid.HasErrors() {
		for _, errValidate := range valid.Errors {
			err = errors.New(errValidate.Key + ": " + errValidate.Message)
			return
		}
	}

	user = middleware.GetAuth(ctx)

	photo, err = u.repo.GetPhotoByIDAndUserID(ctx, user.ID, req.PhotoId)
	if err != nil || photo == nil {
		err = errors.New("photo you want to edit is not your photo")
		return
	}

	photo = &entities.Photo{
		ID:        req.PhotoId,
		Caption:   req.PhotoTitle,
		PhotoURL:  req.PhotoUrl,
		CreatedAt: time.Now(),
		UserID:    user.ID,
	}

	err = u.repo.UpdatePhoto(ctx, photo)

	return
}

func (u Usecases) DeletePhoto(ctx context.Context, req entities.Request) (httpStatus int, err error) {

	var (
		user     entities.User
		photo    entities.Photo
		photoPtr *entities.Photo
	)

	defer func() {
		helper.RecoverPanic()

		httpStatus = http.StatusOK
		if err != nil {
			httpStatus = http.StatusBadRequest
		}
	}()

	//general field validation
	valid := validation.Validation{}
	valid.Required(req.PhotoId, "photoId")

	if valid.HasErrors() {
		for _, errValidate := range valid.Errors {
			err = errors.New(errValidate.Key + ": " + errValidate.Message)
			return
		}
	}

	user = middleware.GetAuth(ctx)

	photoPtr, err = u.repo.GetPhotoByIDAndUserID(ctx, user.ID, req.PhotoId)
	if err != nil || photoPtr == nil {
		err = errors.New("photo you want to delete is not your photo")
		return
	}

	photo = entities.Photo{
		ID:        req.PhotoId,
		Caption:   req.PhotoTitle,
		PhotoURL:  req.PhotoUrl,
		CreatedAt: time.Now(),
		UserID:    user.ID,
	}

	err = u.repo.DeletePhoto(ctx, photo.ID)

	return
}
