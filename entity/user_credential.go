package entity

type UserCredential struct {
	UserId   int64
	Salt     string
	Password string
}
