package core

import (
	"context"
	"fmt"
	"github.com/imjowend/microservicio-usuarios-y-autenticacion/internal/core/user"
)

type UseCase struct {
	user user.RepositoryPort
}

func NewUseCase(u user.RepositoryPort) UseCasePort {
	return &UseCase{
		user: u,
	}
}

func (c *UseCase) CreateUser(ctx context.Context, user *user.User) error {
	fmt.Println(user)
	return nil
}
