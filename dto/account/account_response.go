package dto

import (
	"github.com/fiaufar/sawit-pro-test/entity"
)

type GetProfileResponse struct {
	Fullname    string `json:"fullname"`
	PhoneNumber string `json:"phone_number"`
}

func CreateGetProfileResponse(user *entity.User) *GetProfileResponse {
	return &GetProfileResponse{
		Fullname:    user.Fullname,
		PhoneNumber: user.PhoneNumber,
	}
}
