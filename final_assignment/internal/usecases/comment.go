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

func (u Usecases) CreateComment(ctx context.Context, req entities.Request) (httpStatus int, err error) {

	var (
		user    entities.User
		comment entities.Comment
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
	valid.Required(req.Message, "message")
	valid.Required(req.PhotoId, "photoId")

	if valid.HasErrors() {
		for _, errValidate := range valid.Errors {
			err = errors.New(errValidate.Key + ": " + errValidate.Message)
			return
		}
	}

	photo, err := u.repo.GetPhotoByID(ctx, req.PhotoId)
	if err != nil || photo == nil {
		err = errors.New("invalid photo id")
		return
	}

	user = middleware.GetAuth(ctx)
	comment = entities.Comment{
		Message:   req.Message,
		PhotoID:   req.PhotoId,
		CreatedAt: time.Now(),
		UserID:    user.ID,
	}

	err = u.repo.CreateComment(ctx, &comment)

	return
}

func (u Usecases) GetAllComment(ctx context.Context) (comments *[]entities.Comment, httpStatus int, err error) {

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

	comments, err = u.repo.GetCommentByUserID(ctx, user.ID)
	return
}

func (u Usecases) GetCommentById(ctx context.Context, req entities.Request) (comment *entities.Comment, httpStatus int, err error) {

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
	valid.Required(req.CommentId, "commentId")

	if valid.HasErrors() {
		for _, errValidate := range valid.Errors {
			err = errors.New(errValidate.Key + ": " + errValidate.Message)
			return
		}
	}

	user = middleware.GetAuth(ctx)

	comment, err = u.repo.GetCommentByIDByUserId(ctx, user.ID, req.CommentId)
	return
}

func (u Usecases) UpdateComment(ctx context.Context, req entities.Request) (comment *entities.Comment, httpStatus int, err error) {

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
	valid.Required(req.CommentId, "commentId")
	valid.Required(req.Message, "message")

	if valid.HasErrors() {
		for _, errValidate := range valid.Errors {
			err = errors.New(errValidate.Key + ": " + errValidate.Message)
			return
		}
	}

	user = middleware.GetAuth(ctx)

	comment, err = u.repo.GetCommentByIDByUserId(ctx, user.ID, req.CommentId)
	if err != nil || comment == nil {
		err = errors.New("comment you want to edit is not your comment")
		return
	}

	photo, err := u.repo.GetPhotoByID(ctx, req.PhotoId)
	if err != nil || photo == nil {
		err = errors.New("invalid photo id")
		return
	}

	comment = &entities.Comment{
		ID:        req.CommentId,
		Message:   req.Message,
		PhotoID:   req.PhotoId,
		UpdatedAt: time.Now(),
		UserID:    user.ID,
	}

	err = u.repo.UpdateComment(ctx, comment)

	return
}

func (u Usecases) DeleteComment(ctx context.Context, req entities.Request) (httpStatus int, err error) {

	var (
		user       entities.User
		comment    entities.Comment
		commentPtr *entities.Comment
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
	valid.Required(req.CommentId, "commentId")

	if valid.HasErrors() {
		for _, errValidate := range valid.Errors {
			err = errors.New(errValidate.Key + ": " + errValidate.Message)
			return
		}
	}

	user = middleware.GetAuth(ctx)

	commentPtr, err = u.repo.GetCommentByIDByUserId(ctx, user.ID, req.CommentId)
	if err != nil || commentPtr == nil {
		err = errors.New("comment you want to delete is not your comment")
		return
	}

	comment = entities.Comment{
		ID:        req.CommentId,
		CreatedAt: time.Now(),
		UserID:    user.ID,
	}

	err = u.repo.DeleteComment(ctx, comment.ID)

	return
}
