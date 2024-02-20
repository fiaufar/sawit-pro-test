package handler

import (
	// jwtv3 "github.com/dgrijalva/jwt-go"

	"strings"

	dto "github.com/fiaufar/sawit-pro-test/dto/account"
	"github.com/fiaufar/sawit-pro-test/util"
	"github.com/labstack/echo/v4"
)

func (s *Server) Profile(ctx echo.Context) error {
	claims, err := util.GetJwtClaims(ctx)
	if err != nil {
		util.Log.Error(err.Error())
		return ctx.JSON(s.ResponseUtil.Unauthorized("unauthorized", nil))
	}

	user, err := s.AccountService.GetProfile(claims.UserId)
	if err != nil {
		util.Log.Error(err)
		return ctx.JSON(s.ResponseUtil.BadRequest("failed get profile", err.Error()))
	}

	return ctx.JSON(s.ResponseUtil.Ok("success get profile", dto.CreateGetProfileResponse(user)))
}

func (s *Server) UpdateProfile(ctx echo.Context) error {
	claims, err := util.GetJwtClaims(ctx)
	if err != nil {
		util.Log.Error(err.Error())
		return ctx.JSON(s.ResponseUtil.Unauthorized("unauthorized", nil))
	}

	var req dto.UpdateProfileRequest

	// Bind the request
	if err := ctx.Bind(&req); err != nil {
		util.Log.Error(err)
		return ctx.JSON(s.ResponseUtil.BadRequest("bad request", err.Error()))
	}

	// Validate Form Request
	if err := s.Validator.Struct(req); err != nil {
		util.Log.Error(err)
		return ctx.JSON(s.ResponseUtil.BadRequest("bad request", strings.Split(err.Error(), "\n")))
	}

	err = s.AccountService.UpdateProfile(req, claims.UserId)
	if err != nil {
		util.Log.Error(err)
		return ctx.JSON(s.ResponseUtil.BadRequest("failed update profile", nil))
	}

	return ctx.JSON(s.ResponseUtil.Ok("success update profile", nil))
}
