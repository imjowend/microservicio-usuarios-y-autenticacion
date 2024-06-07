package core

import (
	"context"
	"github.com/imjowend/microservicio-usuarios-y-autenticacion/internal/core/user"
)

type UseCasePort struct {
	CreateUser(context.Context, *user.User) error
}
