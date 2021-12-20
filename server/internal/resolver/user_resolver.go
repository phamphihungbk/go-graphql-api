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

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	return r.userService.CreateItem(input)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input model.UpdateUserInput) (*model.User, error) {
	return r.userService.UpdateItem(id, input)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (bool, error) {
	return r.userService.DeleteItem(id)
}
