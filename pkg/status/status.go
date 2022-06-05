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
	ActionTypeErrCode       = 10007
	TokenErrCode            = 10008
)

const (
	SuccessMsg             = "Success"
	ServiceErrMsg          = "Service is unable to start successfully"
	ParamErrMsg            = "Wrong Parameter has been given"
	LoginErrMsg            = "Wrong username or password"
	RegisterErrMsg         = "Register Error"
	UserNotExistErrMsg     = "User does not exists"
	UserAlreadyExistErrMsg = "User already exists"
	ActionTypeErrMsg       = "Action invalid"
	TokenErrMsg            = "Token invalid"
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
	Success             = NewStatus(SuccessCode, SuccessMsg)
	ServiceErr          = NewStatus(ServiceErrCode, ServiceErrMsg)
	ParamErr            = NewStatus(ParamErrCode, ParamErrMsg)
	LoginErr            = NewStatus(LoginErrCode, LoginErrMsg)
	RegisterErr         = NewStatus(RegisterErrCode, RegisterErrMsg)
	UserNotExistErr     = NewStatus(UserNotExistErrCode, UserNotExistErrMsg)
	UserAlreadyExistErr = NewStatus(UserAlreadyExistErrCode, UserAlreadyExistErrMsg)
	ActionTypeErr       = NewStatus(ActionTypeErrCode, ActionTypeErrMsg)
	TokenErr            = NewStatus(TokenErrCode, TokenErrMsg)
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
