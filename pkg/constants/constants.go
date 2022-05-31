package constants

const (
	UserTableName     = "user"
	VideoTableName    = "video"
	FavorateTableName = "favorate"
	CommentTableName  = "comment"
	FollowTableName   = "follow"
	// SecretKey               = "secret key"
	// IdentityKey             = "id"
	// Total                   = "total"
	// Notes                   = "notes"
	// NoteID                  = "note_id"
	ApiServiceName    = "ApiService"
	UserServiceName   = "UserService"
	VideoServiceName  = "VideoService"
	ActionServiceName = "ActionService"
	// MySQLDefaultDSN			= "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	MySQLDefaultDSN = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress     = "127.0.0.1:2379"
	// EtcdAddress = ""
	// CPURateLimit float64 = 80.0
	// DefaultLimit         = 10
)
