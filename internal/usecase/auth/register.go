package auth

import (
	"context"

	"github.com/fithrahfauzan/go-arch-example/internal/entity"
	"github.com/fithrahfauzan/go-arch-example/internal/repository"
	"github.com/fithrahfauzan/go-arch-example/internal/usecase"
	"github.com/fithrahfauzan/go-arch-example/pkg/log"
)

type register struct {
	userRepo  repository.UserRepository
	tokenRepo repository.TokenRepository
}

func NewRegisterUseCase(userRepo repository.UserRepository, tokenRepo repository.TokenRepository) usecase.UseCase {
	return &register{userRepo, tokenRepo}
}

func (r *register) Execute(ctx context.Context, input usecase.Input) (usecase.Output, error) {
	req, ok := input.(entity.RegisterRequest)
	if !ok {
		return nil, entity.ErrFailedCastingUseCaseInput
	}

	log.Debug("request", req)
	// TODO

	return entity.RegisterResponse{}, nil
}
