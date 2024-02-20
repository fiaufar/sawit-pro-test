package dto

type UpdateProfileRequest struct {
	Fullname    string `json:"fullname" validate:"omitempty,min=3,max=60"`
	PhoneNumber string `json:"phone_number" validate:"omitempty,phone"`
}
