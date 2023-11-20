package awsutils

import (
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

func (inst *AWSInstance) NewS3Client() *s3.Client {
	return s3.NewFromConfig(inst.awsConfig)
}
