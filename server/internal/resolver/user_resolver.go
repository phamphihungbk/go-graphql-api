package resolver

import (
	"context"
	"github.com/phamphihungbk/go-graphql/internal/model"
)

func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	return r.userService.GetItem(id)
}

func (r *queryResolver) Users(
	ctx context.Context,
	limit int,
	page int,
	sort string,
) ([]*model.User, error) {
	data, err := r.userService.GetAllItems(limit, page, sort)
	return data, err
}

func (r *mutationResolver) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	return r.userService.CreateItem(user)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, user *model.User) *model.User {
	return r.userService.UpdateItem(user)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int) error {
	return r.userService.DeleteItem(id)
}
