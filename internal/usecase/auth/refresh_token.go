package auth

import (
	"context"

	"github.com/fithrahfauzan/go-arch-example/internal/entity"
	"github.com/fithrahfauzan/go-arch-example/internal/repository"
	"github.com/fithrahfauzan/go-arch-example/internal/usecase"
	"github.com/fithrahfauzan/go-arch-example/pkg/log"
)

type refreshToken struct {
	tokenRepo repository.TokenRepository
}

func NewRefreshTokenUseCase(tokenRepo repository.TokenRepository) usecase.UseCase {
	return &refreshToken{tokenRepo}
}

func (s *refreshToken) Execute(ctx context.Context, input usecase.Input) (usecase.Output, error) {
	req, ok := input.(entity.RefreshTokenRequest)
	if !ok {
		return nil, entity.ErrFailedCastingUseCaseInput
	}

	log.Debug("request", req)
	// TODO

	return entity.RefreshTokenResponse{}, nil
}
