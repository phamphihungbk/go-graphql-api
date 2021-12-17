package resolver

import (
	"github.com/phamphihungbk/go-graphql/internal/service"
)

type Resolver struct {
	UserService service.UserService
}

type QueryResolver struct {
	*Resolver
}

type MutationResolver struct {
	*Resolver
}

func (r *Resolver) Query() *QueryResolver {
	return &QueryResolver{r}
}

func (r *Resolver) Mutation() *MutationResolver {
	return &MutationResolver{r}
}
