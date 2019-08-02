package user

import (
	"context"
	"errors"
	"time"

	"github.com/zainul/arkana-kit/internal/contract"
	"github.com/zainul/arkana-kit/internal/pkg/error/deliveryerror"
	"github.com/zainul/arkana-kit/internal/pkg/error/usecaseerror"
)

// ActivateUser ...
func (u *Usecase) ActivateUser(ctx context.Context, req contract.ActivateUser) *deliveryerror.Error {
	users, err := u.AccountRepository.UserBy("activation_code", req.Code)

	if err != nil {
		return deliveryerror.GetError(usecaseerror.InternalServerError, err)
	}

	if len(users) == 0 {
		return deliveryerror.GetError(usecaseerror.ActivationCodeNotFound, errors.New("code not found"))
	}
	user := users[0]
	now := time.Now()

	if user.ActivateAt != nil {
		return deliveryerror.GetError(usecaseerror.AccountAlreadyActivate, errors.New("account is active"))
	}

	updatedValue := map[string]interface{}{
		"activate_at": now,
	}

	if err = u.AccountRepository.Update("id", user.ID, updatedValue); err != nil {
		return deliveryerror.GetError(usecaseerror.InternalServerError, err)
	}

	return nil
}
