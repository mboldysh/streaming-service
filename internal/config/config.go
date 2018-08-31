package config

import (
	"errors"
	"flag"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
)

var (
	listenAddr          string
	bucketName          string
	presignedExpireTime time.Duration
)

var (
	defaultListenAddr          = ":8080"
	defaultBucketName          = "streaming-service-bucket"
	defaultPresignedExpireTime = 10 * time.Minute
)

type Config struct {
	AWSCfg            aws.Config
	ListenAddr        string
	TrackBucketName   string
	PresignExpireTime time.Duration
}

func New() (*Config, error) {
	flag.StringVar(&listenAddr, "p", defaultListenAddr, "server port")
	flag.StringVar(&bucketName, "b", defaultBucketName, "bucket name")
	flag.DurationVar(&presignedExpireTime, "t", defaultPresignedExpireTime, "presigned url expire time")
	flag.Parse()

	cfg, err := external.LoadDefaultAWSConfig()

	if err != nil {
		return nil, err
	}

	_, err = cfg.Credentials.Retrieve()

	if err != nil {
		return nil, err
	}

	if cfg.Region == "" {
		return nil, errors.New("Set AWS_REGION env variable")
	}

	return &Config{
		AWSCfg:            cfg,
		ListenAddr:        listenAddr,
		TrackBucketName:   bucketName,
		PresignExpireTime: presignedExpireTime,
	}, nil

}
