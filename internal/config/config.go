package config

import (
	"errors"
	"flag"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/s3manager"
)

var (
	listenAddr          string
	bucketName          string
	presignedExpireTime time.Duration
	dev                 bool
)

var (
	defaultListenAddr          = ":8080"
	defaultBucketName          = "streaming-service-data"
	defaultPresignedExpireTime = 10 * time.Minute
	defaultDev                 = false
)

//Config contains project configuration
type Config struct {
	AWSCfg            aws.Config
	S3Clinet          *s3.S3
	S3uploader        *s3manager.Uploader
	ListenAddr        string
	TrackBucketName   string
	PresignExpireTime time.Duration
}

//New returns new config
func New() (*Config, error) {
	flag.StringVar(&listenAddr, "p", defaultListenAddr, "server port")
	flag.StringVar(&bucketName, "b", defaultBucketName, "bucket name")
	flag.DurationVar(&presignedExpireTime, "t", defaultPresignedExpireTime, "presigned url expire time")
	flag.BoolVar(&dev, "dev", defaultDev, "development mode")
	flag.Parse()

	cfg, err := external.LoadDefaultAWSConfig()

	if err != nil {
		return nil, err
	}

	s3Client, err := createS3Client(cfg, dev)

	if err != nil {
		return nil, err
	}

	s3Uploader := s3manager.NewUploaderWithClient(s3Client)

	return &Config{
		AWSCfg:            cfg,
		S3Clinet:          s3Client,
		S3uploader:        s3Uploader,
		ListenAddr:        listenAddr,
		TrackBucketName:   bucketName,
		PresignExpireTime: presignedExpireTime,
	}, nil

}

//createS3Client returns S3 client. In case if dev flag set to true it returns
//client configured to use with Minio with address 'http://minio:9000'. See Minio configuraton
//in docker-compose.dev.yml file. Otherwise client configured to work with aws s3 wiil be returned
func createS3Client(cfg aws.Config, dev bool) (*s3.S3, error) {
	var s3Client *s3.S3

	if dev {
		cfg.EndpointResolver = createCustomS3Endpoint("http://minio:9000")

		s3Client = s3.New(cfg)
		s3Client.ForcePathStyle = true
	} else {
		_, err := cfg.Credentials.Retrieve()

		if err != nil {
			return nil, err
		}

		if cfg.Region == "" {
			return nil, errors.New("Set AWS_REGION env variable")
		}

		s3Client = s3.New(cfg)
	}

	return s3Client, nil
}

//createCustomS3Endpoint returns aws.EndpointResolverFunc configured to work with
//specified in parameter url
func createCustomS3Endpoint(url string) aws.EndpointResolverFunc {
	defaultResolver := endpoints.NewDefaultResolver()
	s3CustomResolver := func(serice, region string) (aws.Endpoint, error) {
		if serice == endpoints.S3ServiceID {
			return aws.Endpoint{
				URL:           url,
				SigningRegion: "ua-east-1",
			}, nil
		}

		return defaultResolver.ResolveEndpoint(serice, region)
	}

	return s3CustomResolver
}
