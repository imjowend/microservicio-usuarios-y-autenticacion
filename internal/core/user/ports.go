package user

import (
	"context"
)

type RepositoryPort interface {
	CreateUser(context.Context) error
}
