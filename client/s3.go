package client

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

var Client *s3.Client

func InitS3() error {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}

	Client = s3.NewFromConfig(cfg)
	return err
}

func PutObj(client *s3.Client, key string, file []byte) {
	uploader := manager.NewUploader(client)
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String("media-storage-ether7981"),
		Key:         aws.String(key),
		Body:        bytes.NewReader(file),
		ContentType: aws.String("image/jpg"),
	})

	fmt.Println(result, err)
}

func GetObj(client *s3.Client, key string, file *os.File) {
	downloader := manager.NewDownloader(client)
	_, err := downloader.Download(context.TODO(), file, &s3.GetObjectInput{
		Bucket: aws.String("media-storage-ether7981"),
		Key:    aws.String(key),
	})

	fmt.Println(err)
}
