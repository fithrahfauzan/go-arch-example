package repository

import (
	"context"

	"github.com/fithrahfauzan/go-arch-example/internal/entity"
)

type TokenRepository interface {
	Generate(ctx context.Context, userID string, username string) (entity.Token, error)
	Update(ctx context.Context, token string) (entity.Token, error)
	Delete(ctx context.Context, token string) error
}
