package bukcet

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func CreateIfDoesntExists(client *s3.S3, bucketName string) (bool, error) {
	exists, err := Exists(client, bucketName)
	if err != nil {
		return false, err
	}
	if !exists {
		err = Create(client, bucketName)
		if err != nil {
			return false, err
		}
		return true, nil
	} else {
		return false, nil
	}
}

func Exists(client *s3.S3, bucketName string) (bool, error) {
	req := client.HeadBucketRequest(&s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	_, err := req.Send()

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				return false, nil
			case "NotFound":
				return false, nil
			default:
				return false, err
			}
		} else {
			return false, err
		}
	}

	return true, nil
}

func Create(client *s3.S3, bucketName string) error {
	req := client.CreateBucketRequest(&s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})

	_, err := req.Send()

	return err
}
