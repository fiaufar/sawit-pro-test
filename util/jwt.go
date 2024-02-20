package util

import (
	"errors"
	"time"

	"github.com/fiaufar/sawit-pro-test/constant"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AuthenticationToken struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expiredAt"`
}

// func GetJwtToken(userId int64) *AuthenticationToken {
// 	var authToken AuthenticationToken
// 	timeE := time.Now().Add(time.Hour * 12)
// 	timeExp := timeE.Unix()

// 	// Create JWT Token
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"type":   "auth",
// 		"userId": userId,
// 		"exp":    timeExp,
// 	})

// 	// Signing the token with the signing key
// 	secretKey := constant.JWT_SECRET_KEY
// 	var signingkey = []byte(secretKey)
// 	tokenString, _ := token.SignedString(signingkey)

// 	authToken.Token = tokenString
// 	authToken.ExpiredAt = timeE

// 	return &authToken
// }

type JwtCustomClaims struct {
	UserId int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func EchoClaimFunc(ctx echo.Context) jwt.Claims {
	return new(JwtCustomClaims)
}

func ErrorHandler(ctx echo.Context, err error) error {
	Log.Error(err)
	return errors.New("unauthorized")
}

func GetJwtClaims(ctx echo.Context) (*JwtCustomClaims, error) {
	token, ok := ctx.Get("user").(*jwt.Token) // by default token is stored under `user` key
	if !ok {
		return nil, errors.New("failed to get token")
	}

	claims, ok := token.Claims.(*JwtCustomClaims)
	if !ok {
		return nil, errors.New("failed to get token claims")
	}

	return claims, nil
}

func GetJwtToken(userId int64) *AuthenticationToken {

	timeE := time.Now().Add(time.Hour * 12)
	claims := &JwtCustomClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(timeE),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(constant.JWT_SECRET_KEY))
	if err != nil {
		return nil
	}

	return &AuthenticationToken{
		Token:     t,
		ExpiredAt: timeE,
	}
}

// const secret = "secret"

// type JwtCustomClaims struct {
// 	Name   string `json:"name"`
// 	UserId int64  `json:"user_id"`
// 	Admin  bool   `json:"admin"`
// 	jwt.StandardClaims
// }

// func GetJwtTokenEchoV4(userId int64) (*string, error) {
// 	claims := &JwtCustomClaims{
// 		Name:   "Pieter Claerhout",
// 		UserId: userId,
// 		Admin:  true,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: &jwt.Time{
// 				Time: time.Now().Add(time.Hour * 72),
// 			},
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	t, err := token.SignedString([]byte(constant.JWT_SECRET_KEY))
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &t, nil
// }
