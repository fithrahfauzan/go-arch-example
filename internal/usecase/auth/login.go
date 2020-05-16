package auth

import (
	"context"

	"github.com/fithrahfauzan/go-arch-example/internal/entity"
	"github.com/fithrahfauzan/go-arch-example/internal/repository"
	"github.com/fithrahfauzan/go-arch-example/internal/usecase"
)

type login struct {
	userRepo  repository.UserRepository
	tokenRepo repository.TokenRepository
}

func NewLoginUseCase(userRepo repository.UserRepository, tokenRepo repository.TokenRepository) usecase.UseCase {
	return &login{userRepo, tokenRepo}
}

func (l *login) Execute(ctx context.Context, input usecase.Input) (usecase.Output, error) {
	req, ok := input.(entity.LoginRequest)
	if !ok {
		return nil, entity.ErrFailedCastingUseCaseInput
	}

	var resp entity.LoginResponse

	user, err := l.userRepo.FindByUsernameAndPassword(ctx, req.Username, req.Password)
	if err != nil {
		return nil, entity.ErrUserNotFound
	}

	token, err := l.tokenRepo.Generate(ctx, user.ID, user.Username)
	if err != nil {
		return resp, entity.ErrFailedGeneratingToken
	}

	return entity.LoginResponse{
		UserID:         user.ID,
		Username:       user.Username,
		Role:           user.Role,
		Token:          token.Value,
		RefreshToken:   token.RefreshToken,
		TokenExpiredAt: token.ExpiredAt,
	}, nil
}
