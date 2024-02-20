package handler

import (
	"strings"

	dto "github.com/fiaufar/sawit-pro-test/dto/auth"
	"github.com/fiaufar/sawit-pro-test/util"
	"github.com/labstack/echo/v4"
)

func (s *Server) Login(ctx echo.Context) error {
	var req dto.LoginRequest

	// Bind the request
	if err := ctx.Bind(&req); err != nil {
		util.Log.Error(err)
		return ctx.JSON(s.ResponseUtil.BadRequest("bad request", err.Error()))
	}

	// Validate Form Request
	if err := s.Validator.Struct(req); err != nil {
		util.Log.Error(err)
		return ctx.JSON(s.ResponseUtil.BadRequest("failed login", strings.Split(err.Error(), "\n")))
	}

	token, userId, err := s.AuthService.Login(req)
	if err != nil {
		util.Log.Error(err)
		return ctx.JSON(s.ResponseUtil.BadRequest("failed login", err.Error()))
	}

	return ctx.JSON(s.ResponseUtil.Ok("success login", dto.CreateLoginResponse(token, userId)))
}

func (s *Server) Register(ctx echo.Context) error {
	var req dto.RegisterRequest

	// Bind the request
	if err := ctx.Bind(&req); err != nil {
		util.Log.Error(err)
		return ctx.JSON(s.ResponseUtil.BadRequest("failed register", err.Error()))
	}

	// Validate Form Request
	if err := s.Validator.Struct(req); err != nil {
		util.Log.Error(err)
		return ctx.JSON(s.ResponseUtil.BadRequest("failed register", strings.Split(err.Error(), "\n")))
	}

	userId, err := s.AuthService.Register(req)
	if err != nil {
		util.Log.Error(err)
		return ctx.JSON(s.ResponseUtil.BadRequest("failed register", nil))
	}

	return ctx.JSON(s.ResponseUtil.Ok("success register new user", dto.CreateRegisterResponse(userId)))
}
