package minio_config

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client
var qrCodesBucket string

func Connect(endpoint string, bucket string, accessKeyID string, secretAccessKey string) {
	var err error
	minioClient, err = minio.New(endpoint, &minio.Options{
		Secure: false,
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})

	if err != nil {
		panic(err)
	}

	qrCodesBucket = bucket

	log.Printf("Connected to minio server at %s", endpoint)
}

func Client() (*minio.Client, string) {
	return minioClient, qrCodesBucket
}
