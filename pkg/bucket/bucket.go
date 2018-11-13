package bucket

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

//CreateIfDoesntExists creates new bucket if the bucket with specified in
//parameter name doesn't exist and returns true. Return false if bucket
//already exists.
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
	}
	return false, nil
}

//Exists returns true if the bucket with specified name
//is exists. False otherwise.
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

//ObjectExists returns true if the object exists in specified bucket. False otherwise.
func ObjectExists(client *s3.S3, bucketName, key string) (bool, error) {
	req := client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})

	_, err := req.Send()

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchKey:
				return false, nil
			default:
				return false, err
			}
		}
	}

	return true, nil
}

//Create creates bucket with specified name
func Create(client *s3.S3, bucketName string) error {
	req := client.CreateBucketRequest(&s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})

	_, err := req.Send()

	return err
}
