package user

import (
	"context"

	"github.com/fithrahfauzan/go-arch-example/internal/usecase"
)

type GetUser struct {
}

func NewGetUserUseCase() usecase.UseCase {
	return &GetUser{}
}

func (s *GetUser) Execute(ctx context.Context, input usecase.Input) (usecase.Output, error) {
	return nil, nil
}
