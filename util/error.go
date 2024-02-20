package util

import "errors"

type ErrorUtil interface {
	FailedLogin() error
	NotFound() error
	FailedStore() error
	FailedUpdate() error
	FailedDelete() error
	New(msg string) error
}

type ErrorUtilImpl struct {
	objName string
	lang    string
}

func NewErrorUtil(objName string) *ErrorUtilImpl {
	return &ErrorUtilImpl{
		objName: objName,
		lang:    "id",
	}
}

func (e *ErrorUtilImpl) FailedLogin() error {
	return errors.New("failed login")
}

func (e *ErrorUtilImpl) NotFound() error {
	return errors.New(e.objName + " Not found")
}

func (e *ErrorUtilImpl) FailedStore() error {
	return errors.New("failed store " + e.objName)
}

func (e *ErrorUtilImpl) FailedUpdate() error {
	return errors.New("failed update " + e.objName + "!")
}

func (e *ErrorUtilImpl) FailedDelete() error {
	return errors.New("failed delete " + e.objName + "!")
}

func (e *ErrorUtilImpl) New(msg string) error {
	return errors.New(msg)
}
