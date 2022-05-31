package status

import (
	"errors"
	"fmt"
)

const (
	SuccessCode             = 0
	ServiceErrCode          = 10001
	ParamErrCode            = 10002
	LoginErrCode            = 10003
	RegisterErrCode         = 10004
	UserNotExistErrCode     = 10005
	UserAlreadyExistErrCode = 10006
)

type Status struct {
	StatusCode int64
	StatusMsg  string
}

func (s Status) Error() string {
	return fmt.Sprintf("status code=%d, status msg=%s", s.StatusCode, s.StatusMsg)
}

func NewStatus(code int64, msg string) Status {
	return Status{code, msg}
}

func (s Status) UpdateMessage(msg string) Status {
	s.StatusMsg = msg
	return s
}

var (
	Success             = NewStatus(SuccessCode, "Success")
	ServiceErr          = NewStatus(ServiceErrCode, "Service is unable to start successfully")
	ParamErr            = NewStatus(ParamErrCode, "Wrong Parameter has been given")
	LoginErr            = NewStatus(LoginErrCode, "Wrong username or password")
	RegisterErr         = NewStatus(RegisterErrCode, "Register Error")
	UserNotExistErr     = NewStatus(UserNotExistErrCode, "User does not exists")
	UserAlreadyExistErr = NewStatus(UserAlreadyExistErrCode, "User already exists")
)

// ConvertErr convert error to Errno
func ConvertErrorToStatus(err error) Status {
	st := Status{}
	if errors.As(err, &st) {
		return st
	}

	s := ServiceErr
	s.StatusMsg = err.Error()
	return s
}
