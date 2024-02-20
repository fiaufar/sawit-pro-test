package handler

import (
	"github.com/fiaufar/sawit-pro-test/service"
	"github.com/fiaufar/sawit-pro-test/util"
	"github.com/go-playground/validator"
)

type Server struct {
	AuthService    service.AuthServiceInterface
	AccountService service.AccountServiceInterface
	Validator      *validator.Validate
	ResponseUtil   util.ResponseUtil
}

type NewServerOptions struct {
	AuthService    service.AuthServiceInterface
	AccountService service.AccountServiceInterface
	Validator      *validator.Validate
	ResponseUtil   util.ResponseUtil
}

func NewServer(opts NewServerOptions) *Server {
	return &Server{
		AuthService:    opts.AuthService,
		AccountService: opts.AccountService,
		Validator:      opts.Validator,
		ResponseUtil:   opts.ResponseUtil,
	}
}
