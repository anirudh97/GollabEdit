package awsutils

import (
	"bytes"
	"io"
	"sync"

	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	once           sync.Once
	instance       *AWSInstance
	awsInstanceErr error
)

type AWSInstance struct {
	awsConfig aws.Config
}

type S3Client struct {
	client *s3.Client
}

// Initializes the AWS instance only once (Singleton)
func GetInstance() (*AWSInstance, error) {
	once.Do(func() {
		var cfg aws.Config
		cfg, awsInstanceErr = config.LoadDefaultConfig(context.TODO())
		instance = &AWSInstance{awsConfig: cfg}
	})
	return instance, awsInstanceErr
}

// Create S3 Client
func NewS3Client() (*S3Client, error) {
	awsInst, err := GetInstance()
	if err != nil {
		return nil, err
	}
	return &S3Client{client: s3.NewFromConfig(awsInst.awsConfig)}, nil
}

func (inst *S3Client) Upload(bucket string, objectKey string, fileContent []byte) error {
	input := &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &objectKey,
		Body:   bytes.NewReader(fileContent),
	}

	_, err := inst.client.PutObject(context.TODO(), input)

	return err
}

func (client *S3Client) Download(bucket string, objectKey string) ([]byte, error) {
	input := &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &objectKey,
	}

	result, err := client.client.GetObject(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()

	fileContent, err := io.ReadAll(result.Body)

	return fileContent, err
}
