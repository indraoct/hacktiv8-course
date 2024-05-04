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

func (u Usecases) CreateSocial(ctx context.Context, req entities.Request) (httpStatus int, err error) {

	var (
		user   entities.User
		social entities.SocialMedia
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
	valid.Required(req.SocialName, "socialName")
	valid.Required(req.SocialUrl, "socialUrl")

	if valid.HasErrors() {
		for _, errValidate := range valid.Errors {
			err = errors.New(errValidate.Key + ": " + errValidate.Message)
			return
		}
	}

	user = middleware.GetAuth(ctx)
	social = entities.SocialMedia{
		Name:           req.SocialName,
		SocialMediaURL: req.SocialUrl,
		CreatedAt:      time.Now(),
		UserID:         user.ID,
	}

	err = u.repo.CreateSocialMedia(ctx, &social)

	return
}

func (u Usecases) GetAllSocial(ctx context.Context) (socials *[]entities.SocialMedia, httpStatus int, err error) {

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

	socials, err = u.repo.GetSocialByUserID(ctx, user.ID)
	return
}

func (u Usecases) GetSocialById(ctx context.Context, req entities.Request) (social *entities.SocialMedia, httpStatus int, err error) {

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
	valid.Required(req.SocialId, "socialId")

	if valid.HasErrors() {
		for _, errValidate := range valid.Errors {
			err = errors.New(errValidate.Key + ": " + errValidate.Message)
			return
		}
	}

	user = middleware.GetAuth(ctx)

	social, err = u.repo.GetSocialByIDAndUserID(ctx, user.ID, req.SocialId)
	return
}

func (u Usecases) UpdateSocial(ctx context.Context, req entities.Request) (social *entities.SocialMedia, httpStatus int, err error) {

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
	valid.Required(req.SocialId, "socialId")
	valid.Required(req.SocialName, "socialName")
	valid.Required(req.SocialUrl, "socialUrl")

	if valid.HasErrors() {
		for _, errValidate := range valid.Errors {
			err = errors.New(errValidate.Key + ": " + errValidate.Message)
			return
		}
	}

	user = middleware.GetAuth(ctx)

	social, err = u.repo.GetSocialByIDAndUserID(ctx, user.ID, req.SocialId)
	if err != nil || social == nil {
		err = errors.New("social media you want to edit is not your social media")
		return
	}

	social = &entities.SocialMedia{
		ID:             req.SocialId,
		Name:           req.SocialName,
		SocialMediaURL: req.SocialUrl,
		UpdatedAt:      time.Now(),
		UserID:         user.ID,
	}

	err = u.repo.UpdateSocialMedia(ctx, social)

	return
}

func (u Usecases) DeleteSocial(ctx context.Context, req entities.Request) (httpStatus int, err error) {

	var (
		user      entities.User
		social    entities.SocialMedia
		socialPtr *entities.SocialMedia
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
	valid.Required(req.SocialId, "socialId")

	if valid.HasErrors() {
		for _, errValidate := range valid.Errors {
			err = errors.New(errValidate.Key + ": " + errValidate.Message)
			return
		}
	}

	user = middleware.GetAuth(ctx)

	socialPtr, err = u.repo.GetSocialByIDAndUserID(ctx, user.ID, req.SocialId)
	if err != nil || socialPtr == nil {
		err = errors.New("social media you want to delete is not your social media")
		return
	}

	social = entities.SocialMedia{
		ID: req.SocialId,
	}

	err = u.repo.DeleteSocialMedia(ctx, social.ID)

	return
}
