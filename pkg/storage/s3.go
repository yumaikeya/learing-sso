package storage

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewLocalS3Client() *s3.Client {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			PartitionID:   "aws",
			URL:           "http://s3:4566",
			SigningRegion: "localhost:4566",
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("access_key", "secret_access", os.Getenv("S3_REGION"))),
		config.WithEndpointResolverWithOptions(customResolver),
	)
	if err != nil {
		fmt.Printf("Fail to create AWS config: %#v", err.Error())
		panic(err)
	}

	return s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.Region = os.Getenv("S3_REGION")
		o.UsePathStyle = true
	})
}

func NewS3Client() *s3.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	return s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.Region = os.Getenv("S3_REGION")
	})
}
