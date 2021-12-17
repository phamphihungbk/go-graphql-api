package resolver

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/phamphihungbk/go-graphql/internal/model"
	"net/http"
)

type User interface {
}

func (r *QueryResolver) User(ctx *context.Context) (User, error) {
	param := ctx.Value("user_id")
	data, err := r.UserService.GetItem(param)
	return data, err
}

func (r *QueryResolver) Users(ctx *context.Context) ([]User, error) {
	params := ctx.Request.URL.Query()
	data, err := r.UserService.GetAllItems(params)
	return data, err
}

func (r *MutationResolver) CreateUser(ctx *context.Context) User {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}
	userData := append(userData, user)

	return r.UserService.CreateItem(userData)
}

func (r *MutationResolver) UpdateUser(ctx *context.Context) User {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}
	userData := append(userData, user)

	return r.UserService.UpdateItem(userData)
}

func (r *MutationResolver) DeleteUser(ctx *context.Context) error {
	param := ctx.Value("user_id")
	return r.UserService.DeleteItem(param)
}
