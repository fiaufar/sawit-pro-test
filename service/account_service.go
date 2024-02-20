package service

import (
	dto "github.com/fiaufar/sawit-pro-test/dto/account"
	"github.com/fiaufar/sawit-pro-test/entity"
	"github.com/fiaufar/sawit-pro-test/repository"
	"github.com/fiaufar/sawit-pro-test/util"
)

type AccountServiceInterface interface {
	GetProfile(userId int64) (*entity.User, error)
	UpdateProfile(reqBody dto.UpdateProfileRequest, userId int64) error
}

type AccountService struct {
	UserRepo  *repository.UserRepository
	ErrorUtil util.ErrorUtil
}

type NewAccountServiceOptions struct {
	UserRepo  *repository.UserRepository
	ErrorUtil util.ErrorUtil
}

func NewAccountService(opts NewAccountServiceOptions) *AccountService {
	return &AccountService{
		UserRepo:  opts.UserRepo,
		ErrorUtil: util.NewErrorUtil("Account"),
	}
}

func (svc *AccountService) GetProfile(userId int64) (*entity.User, error) {
	user, err := svc.UserRepo.GetById(userId)
	if err != nil {
		util.Log.Error(err)
		return nil, err
	}
	return user, nil
}

func (svc *AccountService) UpdateProfile(reqBody dto.UpdateProfileRequest, userId int64) (err error) {
	user, err := svc.UserRepo.GetById(userId)
	if err != nil {
		util.Log.Error(err)
		return err
	}

	if reqBody.Fullname != "" {
		user.Fullname = reqBody.Fullname
	}

	if reqBody.PhoneNumber != "" {
		user.PhoneNumber = reqBody.PhoneNumber
	}

	err = svc.UserRepo.Update(user)
	if err != nil {
		util.Log.Error(err)
		return
	}

	return
}
