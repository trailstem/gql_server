package services

import (
	"context"

	"github.com/trailstem/graphql-server/graph/model"
	"github.com/volatiletech/sqlboiler/boil"
)

type Services interface {
	UserService
}

type services struct {
	*userService
}

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		userService: &userService{exec: exec},
	}
}
