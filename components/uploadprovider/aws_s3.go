package uploadprovider

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/0xThomas3000/food_delivery/common"
)

// The struct that implements the "UploadProvider" interface (by declaring "SaveFileUploaded()")
type s3Provider struct {
	bucketName string
	region     string
	apiKey     string
	secret     string
	domain     string
	client     *s3.Client
	// PresignClient *s3.PresignClient
}

func NewS3Provider(bucketName string, region string, apiKey string, secret string, domain string) *s3Provider {
	provider := &s3Provider{
		bucketName: bucketName,
		region:     region,
		apiKey:     apiKey,
		secret:     secret,
		domain:     domain,
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Using the Config value, create the s3 client
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.Region = provider.region
		o.Credentials = aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(
			provider.apiKey, // Access key ID
			provider.secret, // Secret access key
			"",              // Token can be ignore
		))
	})

	provider.client = client // client chỉ dùng 1 lần => NewFromConfig(...) ra rồi gán xài luôn ko cần New lại

	return provider
}

func (provider *s3Provider) SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error) {
	fileBytes := bytes.NewReader(data) // bọc data vào NewReader
	fileType := http.DetectContentType(data)

	_, err := provider.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(provider.bucketName),
		Key:         aws.String(dst), // Key chính là file name
		ACL:         "private",
		ContentType: aws.String(fileType),
		Body:        fileBytes, // Body: chính là io của buffer data (mảng []byte "data" muốn dc chuyển thành Body)
	})

	if err != nil {
		return nil, err
	}

	/***** FOR Create a presigned URL *****/
	// var lifetimeSecs int64 = 5
	// _, err := provider.PresignClient.PresignPutObject(context.TODO(), &s3.PutObjectInput{
	// 	Bucket: aws.String(provider.bucketName),
	// 	Key:    aws.String(dst),
	// 	ACL:    "private",
	// }, func(opts *s3.PresignOptions) {
	// 	opts.Expires = time.Duration(lifetimeSecs * int64(time.Second))
	// })
	// if err != nil {
	// 	log.Printf("Couldn't get a presigned request to put %v:%v. Here's why: %v\n",
	// 		provider.bucketName, dst, err)
	// }

	img := &common.Image{
		Url:       fmt.Sprintf("%s/%s", provider.domain, dst),
		CloudName: "s3",
	}

	return img, nil
}
