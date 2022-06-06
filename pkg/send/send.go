package send

import (
	"net/http"
	"reflect"
	"unsafe"

	"github.com/CharmingCharm/DouSheng/kitex_gen/base"
	"github.com/CharmingCharm/DouSheng/pkg/status"
	"github.com/gin-gonic/gin"
)

func SendStatus(c *gin.Context, err error, data interface{}) {
	st := status.ConvertErrorToStatus(err)
	code := reflect.ValueOf(data).Elem().FieldByName("StatusCode")
	code = reflect.NewAt(code.Type(), unsafe.Pointer(code.UnsafeAddr())).Elem()

	msg := reflect.ValueOf(data).Elem().FieldByName("StatusMsg")
	msg = reflect.NewAt(msg.Type(), unsafe.Pointer(msg.UnsafeAddr())).Elem()

	newCode := reflect.ValueOf(st.StatusCode)
	newMsg := reflect.ValueOf(st.StatusMsg)

	code.Set(newCode)
	msg.Set(newMsg)
	c.JSON(http.StatusOK, data)
}

func SendResp(c *gin.Context, resp base.BaseResp, data interface{}) {
	code := reflect.ValueOf(data).Elem().FieldByName("StatusCode")
	code = reflect.NewAt(code.Type(), unsafe.Pointer(code.UnsafeAddr())).Elem()

	msg := reflect.ValueOf(data).Elem().FieldByName("StatusMsg")
	msg = reflect.NewAt(msg.Type(), unsafe.Pointer(msg.UnsafeAddr())).Elem()

	newCode := reflect.ValueOf(resp.StatusCode)
	newMsg := reflect.ValueOf(resp.StatusMessage)

	code.Set(newCode)
	msg.Set(newMsg)

	c.JSON(http.StatusOK, data)
}
