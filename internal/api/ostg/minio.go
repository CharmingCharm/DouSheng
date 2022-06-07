package ostg

import (
	"context"
	"log"
	"mime/multipart"

	"github.com/minio/minio-go/v7"

	"github.com/CharmingCharm/DouSheng/pkg/constants"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinIOClient *minio.Client

func init() {
	var err error
	MinIOClient, err = minio.New(constants.MinIOEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(constants.MinIOId, constants.MinIOSecret, ""),
		Secure: false,
	})
	if err != nil {
		panic(err)
	}
	location := "us-east-1"
	err = MinIOClient.MakeBucket(context.TODO(), "video", minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := MinIOClient.BucketExists(context.TODO(), "video")
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", "video")
		} else {
			panic(err)
		}
	} else {
		policy := `{"Version": "2012-10-17","Statement": [{"Action": ["s3:GetObject"],"Effect": "Allow","Principal": {"AWS": ["*"]},"Resource": ["arn:aws:s3:::video/*"],"Sid": ""}]}`
		MinIOClient.SetBucketPolicy(context.TODO(), "video", policy)
		log.Printf("Successfully create bucket: %s\n", "video")
	}
}

func UploadVideo(objectName string, username string, fileHeader *multipart.FileHeader) error {
	src, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	_, err = MinIOClient.PutObject(context.TODO(), "video", objectName, src, -1, minio.PutObjectOptions{ContentType: "video/mp4"})
	return err
}
