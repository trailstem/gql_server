package graph

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
)

type DirectiveRoot struct {
	IsAuthenticated func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error)
}

func IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	user := ctx.Value("user")
	if user == nil {
		return nil, errors.New("not authenticated")
	}
	return next(ctx)
}
