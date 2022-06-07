package constants

import (
	"github.com/CharmingCharm/DouSheng/pkg/status"
)

const (
	UserTableName     = "user"
	VideoTableName    = "video"
	FavorateTableName = "favorate"
	CommentTableName  = "comment"
	RelationTableName = "relation"
	ApiServiceName    = "ApiService"
	UserServiceName   = "UserService"
	VideoServiceName  = "VideoService"
	ActionServiceName = "ActionService"
	MySQLDefaultDSN   = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress       = "127.0.0.1:2379"

	DefaultStatusCode  = status.ServiceErrCode
	DefaultStatusMsg   = status.ServiceErrMsg
	DefaultErrPosInt64 = -1
	DefaultErrString   = ""
	DefaultVideoTitle  = "MyVideo"
	DefaultCoverUrl    = "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg"

	// minio监听的地址
	MinIOEndpoint = "127.0.0.1:9000"
	MinIOPos      = "http://192.168.1.68:9000"
	// minio的用户名
	MinIOId = "minioadmin"
	// minio的密码
	MinIOSecret = "minioadmin"
)

type Response struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}
