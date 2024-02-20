package service

import (
	"testing"

	"github.com/fiaufar/sawit-pro-test/constant"
	dto "github.com/fiaufar/sawit-pro-test/dto/auth"
	"github.com/fiaufar/sawit-pro-test/entity"
	"github.com/fiaufar/sawit-pro-test/repository"
	"github.com/fiaufar/sawit-pro-test/repository/mocks"
	"github.com/fiaufar/sawit-pro-test/util"
	"github.com/golang/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthService_Login(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserRepo := mocks.NewMockUserRepositoryInterface(mockCtrl)
	expUserRes := &entity.User{
		Id:                 1,
		Fullname:           "Budiman",
		PhoneNumber:        "+628123123123",
		SuccessfulLoggedIn: 0,
	}
	var wantUserId int64 = 1

	mockUserRepo.EXPECT().GetByPhoneNumber("+628123123123").Return(expUserRes, nil).AnyTimes()

	mockUserCredentialRepo := mocks.NewMockUserCredentialRepositoryInterface(mockCtrl)
	mockUserCredentialRepo.EXPECT().GetByUserId(expUserRes.Id).Return(&entity.UserCredential{
		UserId:   1,
		Salt:     "BpLnfgDsc3WD9F3",
		Password: "$2a$08$H1qMoVHMAtVYHw1Df9iz4Ofn/84WM0h4Cn3hoWChz66nKPwFcNmPe",
	}, nil).AnyTimes()

	expUserRes.SuccessfulLoggedIn += 1
	mockUserRepo.EXPECT().Update(expUserRes).Return(nil).AnyTimes()

	type fields struct {
		UserRepo           repository.UserRepositoryInterface
		UserCredentialRepo repository.UserCredentialRepositoryInterface
		ErrorUtil          util.ErrorUtil
	}
	type args struct {
		reqBody dto.LoginRequest
	}
	type svc struct {
		authService AuthService
	}
	authSvc := svc{
		authService: AuthService{
			UserRepo:           mockUserRepo,
			UserCredentialRepo: mockUserCredentialRepo,
			ErrorUtil:          util.NewErrorUtil("Auth"),
		},
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		svc           svc
		wantAuthToken util.AuthenticationToken
		wantUserId    int64
		wantErr       bool
	}{
		{
			name: "Test login",
			fields: fields{
				UserRepo:           mockUserRepo,
				UserCredentialRepo: mockUserCredentialRepo,
				ErrorUtil:          util.NewErrorUtil("Auth"),
			},
			args: args{
				reqBody: dto.LoginRequest{
					PhoneNumber: "+628123123123",
					Password:    "Ahmad123!",
				},
			},
			svc:           authSvc,
			wantAuthToken: util.AuthenticationToken{},
			wantUserId:    wantUserId,
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAuthToken, gotUserId, err := tt.svc.authService.Login(tt.args.reqBody)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthService.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAuthToken == nil {
				t.Errorf("AuthService.Login() gotAuthToken = %v, want %v", gotAuthToken, tt.wantAuthToken)
			}
			if *gotUserId != tt.wantUserId {
				t.Errorf("AuthService.Login() gotUserId = %v, want %v", gotUserId, tt.wantUserId)
			}
		})
	}
}

func TestAuthService_Register(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserRepo := mocks.NewMockUserRepositoryInterface(mockCtrl)
	reqRegister := dto.RegisterRequest{
		Fullname:    "Budiman",
		PhoneNumber: "+628123123123",
		Password:    "Ahmad123!",
	}

	expUserRes := &entity.User{
		Fullname:    "Budiman",
		PhoneNumber: "+628123123123",
	}
	var wantUserId int64 = 1
	mockUserRepo.EXPECT().Create(expUserRes).Return(wantUserId, nil).AnyTimes()

	mockUserCredentialRepo := mocks.NewMockUserCredentialRepositoryInterface(mockCtrl)
	salt := util.GenerateRandomString(constant.PASSWORD_SALT_LENGTH)
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(reqRegister.Password+salt), constant.AUTH_BCRYPT_COST)
	userCr := entity.UserCredential{
		UserId:   1,
		Salt:     salt,
		Password: string(hashPassword),
	}
	mockUserCredentialRepo.EXPECT().Create(&userCr).Return(nil).AnyTimes()

	type fields struct {
		UserRepo           repository.UserRepositoryInterface
		UserCredentialRepo repository.UserCredentialRepositoryInterface
		ErrorUtil          util.ErrorUtil
	}
	type args struct {
		reqBody dto.RegisterRequest
	}
	type svc struct {
		authService AuthService
	}
	authSvc := svc{
		authService: AuthService{
			UserRepo:           mockUserRepo,
			UserCredentialRepo: mockUserCredentialRepo,
			ErrorUtil:          util.NewErrorUtil("Auth"),
		},
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		svc        svc
		wantUserId *int64
		wantErr    bool
	}{
		{
			name: "Test register",
			fields: fields{
				UserRepo:           mockUserRepo,
				UserCredentialRepo: mockUserCredentialRepo,
				ErrorUtil:          util.NewErrorUtil("Auth"),
			},
			args: args{
				reqBody: reqRegister,
			},
			svc:        authSvc,
			wantUserId: &wantUserId,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUserId, err := tt.svc.authService.Register(tt.args.reqBody)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthService.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if *gotUserId != *tt.wantUserId {
				t.Errorf("AuthService.Register() = %v, want %v", gotUserId, tt.wantUserId)
			}
		})
	}
}
