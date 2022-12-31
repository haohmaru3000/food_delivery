package uploadprovider

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/0xThomas3000/food_delivery/common"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// The struct that implements the "UploadProvider" interface (by declaring "SaveFileUploaded()")
type s3Provider struct {
	bucketName string
	region     string
	apiKey     string
	secret     string
	domain     string
	session    *session.Session
}

func NewS3Provider(bucketName string, region string, apiKey string, secret string, domain string) *s3Provider {
	provider := &s3Provider{
		bucketName: bucketName,
		region:     region,
		apiKey:     apiKey,
		secret:     secret,
		domain:     domain,
	}

	s3Session, err := session.NewSession(&aws.Config{
		Region: aws.String(provider.region),
		Credentials: credentials.NewStaticCredentials(
			provider.apiKey, // Access key ID
			provider.secret, // Secret access key
			""),             // Token can be ignore
	})

	if err != nil {
		log.Fatalln(err)
	}

	provider.session = s3Session // session chỉ dùng 1 lần => NewSession ra rồi gán xài luôn ko cần New lại

	return provider
}

func (provider *s3Provider) SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error) {
	fileBytes := bytes.NewReader(data) // bọc data vào NewReader
	fileType := http.DetectContentType(data)

	_, err := s3.New(provider.session).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(provider.bucketName),
		Key:         aws.String(dst), // Key chính là file name
		ACL:         aws.String("private"),
		ContentType: aws.String(fileType),
		Body:        fileBytes, // Body: chính là io của buffer data (mảng []byte "data" muốn dc chuyển thành Body)
	})

	if err != nil {
		return nil, err
	}

	img := &common.Image{
		Url:       fmt.Sprintf("%s/%s", provider.domain, dst),
		CloudName: "s3",
	}

	return img, nil
}
