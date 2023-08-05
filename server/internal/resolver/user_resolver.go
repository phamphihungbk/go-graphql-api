package resolver

import (
	"context"

	"github.com/phamphihungbk/go-graphql-api/internal/service"

	"github.com/phamphihungbk/go-graphql-api/internal/model"
)

// ========== Mutation ==========
func (r *mutationResolver) CreateUser(ctx context.Context, payload model.CreateUserPayload) (*model.User, error) {
	return r.userService.CreateUser(payload)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, payload model.UpdateUserPayload) (*model.User, error) {
	return r.userService.UpdateUser(payload)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, email string) (bool, error) {
	return r.userService.DeleteUser(email)
}

// ========== Query ==========
func (r *queryResolver) Users(
	ctx context.Context,
	limit int,
	page int,
	sort string,
) (*model.UsersConnection, error) {
	return r.userService.GetAllUsers(limit, page, sort)
}

func (r *queryResolver) Me(ctx context.Context, email string) (*model.User, error) {
	return r.userService.GetUser(email)
}

func (r *mutationResolver) IssueToken(ctx context.Context, payload model.LoginPayload) (*model.AccessToken, error) {
	return r.userService.IssueToken(payload)
}
