package service

import (
	"bytes"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/globalsign/mgo/bson"
)

func PreAWS(fileExtension, folder string) (*session.Session, string, error) {
	tempFileUrl := folder + "/" + bson.NewObjectId().Hex() + fileExtension
	session, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_SECRET_ID"),
			os.Getenv("AWS_SECRET_KEY"),
			os.Getenv("AWS_TOKEN"),
		),
	})

	return session, tempFileUrl, err
}

func UploadFileToS3(s *session.Session, file multipart.File, filename string, size int64) (url string, err error) {
	bucketName := os.Getenv("S3_BUCKET_NAME")
	buffer := make([]byte, size)
	file.Read(buffer)
	url = "https://s3-eu-west-3.amazonaws.com/" + bucketName + "/" + filename
	log.Println(url)

	_, err = s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(bucketName),
		Key:           aws.String(filename),
		ACL:           aws.String("public-read"),
		Body:          bytes.NewReader(buffer),
		ContentLength: aws.Int64(size),
		ContentType: aws.String(http.DetectContentType(buffer)),
		ContentDisposition: aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass: aws.String("INTELLIGENT_TIERING"),
	})
	return url, err
}


