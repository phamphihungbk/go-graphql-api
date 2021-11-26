package resolver

import (
	"context"
	"github.com/phamphihungbk/go-graphql/internal/model"
	"github.com/phamphihungbk/go-graphql/internal/service"
)

type User model.User

type UserResolver struct {
	*service.UserService
}

type QueryResolver struct {
	*UserResolver
}

type MutationResolver struct {
	*UserResolver
}

func NewUserResolver(userService *service.UserService) *UserResolver {
	return &UserResolver{userService}
}

func (r *UserResolver) Query() *QueryResolver {
	return &QueryResolver{r}
}

func (r *QueryResolver) User(ctx context.Context, id uint) (User, error) {
	data, err := r.UserService.GetItem(id)
	return data, err
}

func (r *UserResolver) Mutation() *MutationResolver {
	return &MutationResolver{r}
}

func (r *MutationResolver) Create(ctx context.Context, user User) User {
	return r.UserService.Create(user)
}

func (r *MutationResolver) Update(ctx context.Context, data User) User {
	return r.UserService.Update(data)
}

func (r *MutationResolver) Delete(ctx context.Context, id uint) error {
	return r.UserService.Delete(id)
}
