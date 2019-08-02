package user

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/zainul/arkana-kit/internal/pkg/error/usecaseerror"
	"github.com/zainul/arkana-kit/internal/pkg/hashing"

	"github.com/zainul/arkana-kit/internal/pkg/error/deliveryerror"

	"github.com/zainul/arkana-kit/internal/constant"

	"github.com/zainul/ark/xrandom"

	"github.com/zainul/arkana-kit/internal/contract"
	"github.com/zainul/arkana-kit/internal/entity"
	"github.com/zainul/arkana-kit/internal/repository"
)

// Usecase ...
type Usecase struct {
	AccountRepository repository.Account
}

// RegisterUser is register user account
func (u *Usecase) RegisterUser(ctx context.Context, req contract.RegisterUser) *deliveryerror.Error {

	e := entity.User{}

	users, err := u.AccountRepository.UserBy("email", req.Email)

	if err != nil {
		return deliveryerror.GetError(usecaseerror.InternalServerError, err)
	}

	if len(users) > 0 {
		return deliveryerror.GetError(usecaseerror.UserAlreadyExist, errors.New("user found"))
	}

	accType, ok := constant.UserType[strings.ToUpper(req.AccountType)]

	if !ok {
		return deliveryerror.GetError(usecaseerror.InvalidAccountType, errors.New("cannot cast the user type map"))
	}

	str, _ := xrandom.GenerateRandomString(50)
	e.ID = xrandom.GenerateGeneralID()
	e.FirstName = req.FirstName
	e.LastName = req.LastName
	e.Email = req.Email

	var strPwd string
	if strPwd, err = hashing.HashPassword(req.Password); err != nil {
		return deliveryerror.GetError(usecaseerror.InternalServerError, err)
	}

	e.EncryptedPass = strPwd
	e.PhoneNumber = req.PhoneNumber
	e.AccountType = accType
	e.Username = xrandom.GenerateAlias(req.FirstName)

	// e.InvitationCode = req
	// e.ActivateAt = req.

	e.ActivationCode = str
	now := time.Now()
	e.ModifiedAt = &now

	if err := u.AccountRepository.SaveUser(e); err != nil {
		return deliveryerror.GetError(usecaseerror.InternalServerError, err)
	}

	return nil
}
