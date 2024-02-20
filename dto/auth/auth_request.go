package dto

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type RegisterRequest struct {
	Fullname    string `json:"fullname" validate:"required,min=3,max=60"`
	PhoneNumber string `json:"phone_number" validate:"required,phone"`
	Password    string `json:"password" validate:"required,min=6,max=64,password"`
}
