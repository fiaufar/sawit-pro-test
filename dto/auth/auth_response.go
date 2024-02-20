package dto

import (
	"time"

	"github.com/fiaufar/sawit-pro-test/util"
)

type LoginResponse struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
	Id        int64     `json:"id"`
}

func CreateLoginResponse(token *util.AuthenticationToken, id *int64) *LoginResponse {
	return &LoginResponse{
		Token:     token.Token,
		ExpiredAt: token.ExpiredAt,
		Id:        *id,
	}
}

type RegisterResponse struct {
	Id int64 `json:"id"`
}

func CreateRegisterResponse(id *int64) *RegisterResponse {
	return &RegisterResponse{
		Id: *id,
	}
}
