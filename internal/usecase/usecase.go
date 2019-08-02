package usecase

import (
	"context"

	"github.com/zainul/arkana-kit/internal/contract"
	"github.com/zainul/arkana-kit/internal/pkg/error/deliveryerror"
	"github.com/zainul/arkana-kit/internal/repository"
	"github.com/zainul/arkana-kit/internal/usecase/user"
)

// User ...
type User interface {
	RegisterUser(ctx context.Context, req contract.RegisterUser) *deliveryerror.Error
	ActivateUser(ctx context.Context, req contract.ActivateUser) *deliveryerror.Error
}

// NewUser ...
func NewUser(
	AccountRepository repository.Account,
) User {
	return &user.Usecase{
		AccountRepository: AccountRepository,
	}
}
