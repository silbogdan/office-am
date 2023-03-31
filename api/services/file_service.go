package services

import (
	// "am/office-check-in/minio_config"
	"am/office-check-in/minio_config"
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/aidarkhanov/nanoid"
	"github.com/minio/minio-go/v7"
)

const SEVEN_DAYS_IN_NANOSECONDS = 604800000000000

func Upload(file multipart.File, filename string, contentType string) (string, error) {
	// Get minio client
	minioClient, bucket := minio_config.Client()

	// Upload file to minio
	// 1. Read bytes from file
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return "", err
	}

	// Generate random filename
	alphabet := nanoid.DefaultAlphabet
	size := 6

	id, err := nanoid.Generate(alphabet, size)
	if err != nil {
		return "", err
	}

	filename = fmt.Sprintf("%s_%s", id, filename)

	// 2. Upload bytes to minio
	info, err := minioClient.PutObject(context.Background(), bucket, filename, buf, -1, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "", err
	}

	// 3. Return url
	return info.Location, nil
}

func GetFileUrl(filename string) (string, error) {
	// Get minio client
	minioClient, bucket := minio_config.Client()

	// Get file url
	url, err := minioClient.PresignedGetObject(context.Background(), bucket, filename, SEVEN_DAYS_IN_NANOSECONDS, nil)
	if err != nil {
		return "", err
	}

	return url.String(), err
}
