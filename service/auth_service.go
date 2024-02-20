package service

import (
	"github.com/fiaufar/sawit-pro-test/constant"
	dto "github.com/fiaufar/sawit-pro-test/dto/auth"
	"github.com/fiaufar/sawit-pro-test/entity"
	"github.com/fiaufar/sawit-pro-test/repository"
	"github.com/fiaufar/sawit-pro-test/util"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	Login(dto.LoginRequest) (*util.AuthenticationToken, *int64, error)
	Register(dto.RegisterRequest) (*int64, error)
}

type AuthService struct {
	UserRepo           repository.UserRepositoryInterface
	UserCredentialRepo repository.UserCredentialRepositoryInterface
	ErrorUtil          util.ErrorUtil
}

type NewAuthServiceOptions struct {
	UserRepo           repository.UserRepositoryInterface
	UserCredentialRepo repository.UserCredentialRepositoryInterface
	ErrorUtil          util.ErrorUtil
}

func NewAuthService(opts NewAuthServiceOptions) *AuthService {
	return &AuthService{
		UserRepo:           opts.UserRepo,
		UserCredentialRepo: opts.UserCredentialRepo,
		ErrorUtil:          util.NewErrorUtil("Authentication"),
	}
}

func (svc *AuthService) Login(reqBody dto.LoginRequest) (authToken *util.AuthenticationToken, userId *int64, err error) {
	user, err := svc.UserRepo.GetByPhoneNumber(reqBody.PhoneNumber)
	if err != nil {
		util.Log.Error(err)
		err = svc.ErrorUtil.FailedLogin()
		return
	}

	userCr, err := svc.UserCredentialRepo.GetByUserId(user.Id)
	if err != nil {
		util.Log.Error(err)
		err = svc.ErrorUtil.FailedLogin()
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userCr.Password), []byte(reqBody.Password+userCr.Salt)); err != nil {
		util.Log.Infof(err.Error())
		err = svc.ErrorUtil.FailedLogin()
		return
	}

	authToken = util.GetJwtToken(user.Id)
	userId = &user.Id

	user.SuccessfulLoggedIn += 1
	err = svc.UserRepo.Update(user)
	if err != nil {
		util.Log.Error(err)
		err = svc.ErrorUtil.FailedLogin()
		return
	}

	return
}

func (svc *AuthService) Register(reqBody dto.RegisterRequest) (userId *int64, err error) {
	// Generate random string to be a Password salt
	salt := util.GenerateRandomString(constant.PASSWORD_SALT_LENGTH)

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(reqBody.Password+salt), constant.AUTH_BCRYPT_COST)
	if err != nil {
		util.Log.Error(err)
		return
	}

	user := entity.User{
		Fullname:    reqBody.Fullname,
		PhoneNumber: reqBody.PhoneNumber,
	}
	user.Id, err = svc.UserRepo.Create(&user)
	if err != nil {
		util.Log.Error(err)
		return
	}

	userCr := entity.UserCredential{
		UserId:   user.Id,
		Salt:     salt,
		Password: string(hashPassword),
	}

	err = svc.UserCredentialRepo.Create(&userCr)
	if err != nil {
		util.Log.Error(err)
		return
	}

	userId = &user.Id
	return
}
