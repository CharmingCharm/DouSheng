package ostg

import (
	"mime/multipart"

	"github.com/CharmingCharm/DouSheng/pkg/constants"
	"github.com/minio/minio-go"
	"github.com/minio/minio-go/pkg/credentials"
)

var MinIOClient *minio.Client

func Init() {
	var err error
	MinIOClient, err = minio.NewWithOptions(constants.MinIOEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(constants.MinIOId, constants.MinIOSecret, ""),
		Secure: false,
	})
	if err != nil {
		panic(err)
	}
}

func UploadVideo(objectName string, fileHeader *multipart.FileHeader) error {
	src, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	_, err = MinIOClient.PutObject("video", objectName, src, -1, minio.PutObjectOptions{
		ContentType: ""})
	return err
}
