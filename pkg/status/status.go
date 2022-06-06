package status

import (
	"errors"
	"fmt"
)

const (
	SuccessCode             = 0
	ServiceErrCode          = 10001
	ParamErrCode            = 10002
	ParamControdictErrCode  = 10003
	LoginErrCode            = 10004
	RegisterErrCode         = 10005
	UserNotExistErrCode     = 10006
	UserAlreadyExistErrCode = 10007
	ActionTypeErrCode       = 10008
	TokenErrCode            = 10009
)

const (
	SuccessMsg             = "Success"
	ServiceErrMsg          = "Service is unable to start successfully"
	ParamErrMsg            = "Wrong Parameter has been given"
	ParamControdictErrMsg  = "Contradictory input parameters in record"
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
	ParamControdictErr  = NewStatus(ParamControdictErrCode, ParamControdictErrMsg)
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
