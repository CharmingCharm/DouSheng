package ostg

import (
	"context"
	"github.com/CharmingCharm/DouSheng/pkg/constants"
	"mime/multipart"
)

var MinIOClient *minio.Client

func Init() {
	MinioClient, err = minio.New(constants.MinIOEndpoint, &minio.Options{
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
	_, err = MinioClient.PutObject(context.TODO(), "video", objectName, src, -1, minio.PutObjectOptions{
		ContentType: ""})
	return err
}
