package pack

import (
	"errors"
	"time"

	"github.com/CharmingCharm/DouSheng/kitex_gen/user"
	"github.com/CharmingCharm/DouSheng/pkg/status"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *user.BaseResp {
	if err == nil {
		return baseResp(status.Success)
	}

	s := status.Status{}
	if errors.As(err, &s) {
		return baseResp(s)
	}

	st := status.ServiceErr.UpdateMessage(err.Error())
	return baseResp(st)
}

func baseResp(st status.Status) *user.BaseResp {
	return &user.BaseResp{StatusCode: st.StatusCode, StatusMessage: st.StatusMsg, ServiceTime: time.Now().Unix()}
}
