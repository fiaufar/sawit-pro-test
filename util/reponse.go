package util

import (
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"messages"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type ResponseUtil interface {
	CreateResponse(status int, msg string, data *any, err interface{}) (int, Response)
	Ok(msg string, data interface{}) (int, Response)
	Created(msg string, data interface{}) (int, Response)
	NotFound(msg string) (int, Response)
	BadRequest(msg string, err interface{}) (int, Response)
	Forbidden(msg string, err interface{}) (int, Response)
	Unauthorized(msg string, err interface{}) (int, Response)
}

type ResponseUtilImpl struct {
	Response
}

func NewResponseUtil() *ResponseUtilImpl {
	return &ResponseUtilImpl{
		Response{},
	}
}

func (r *ResponseUtilImpl) CreateResponse(status int, msg string, data *any, err interface{}) (int, Response) {
	return status, Response{
		Code:    status,
		Status:  http.StatusText(status),
		Errors:  err,
		Message: msg,
		Data:    data,
	}
}

func (r *ResponseUtilImpl) Ok(msg string, data interface{}) (int, Response) {
	return r.CreateResponse(http.StatusOK, msg, &data, nil)
}

func (r *ResponseUtilImpl) Created(msg string, data interface{}) (int, Response) {
	return r.CreateResponse(http.StatusCreated, msg, &data, nil)
}

func (r *ResponseUtilImpl) NotFound(msg string) (int, Response) {
	return r.CreateResponse(http.StatusNotFound, msg, nil, nil)
}

func (r *ResponseUtilImpl) BadRequest(msg string, err interface{}) (int, Response) {
	return r.CreateResponse(http.StatusBadRequest, msg, nil, err)
}

func (r *ResponseUtilImpl) Forbidden(msg string, err interface{}) (int, Response) {
	return r.CreateResponse(http.StatusForbidden, msg, nil, err)
}

func (r *ResponseUtilImpl) Unauthorized(msg string, err interface{}) (int, Response) {
	return r.CreateResponse(http.StatusUnauthorized, msg, nil, err)
}
