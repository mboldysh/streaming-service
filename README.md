# Streaming Service

[![Build Status](https://travis-ci.org/mboldysh/streaming-service.svg?branch=master)](https://travis-ci.org/mboldysh/streaming-service) 
[![Go Report Card](https://goreportcard.com/badge/github.com/mboldysh/streaming-service)](https://goreportcard.com/report/github.com/mboldysh/streaming-service)

## Development Setup

There are two options:

### 1. Run with minio

```console
# build streaming-service
make
# build streaming-service docker image
docker build -t streaming-service .
# up streming-service and minio in docker
docker-compose -f docker-compose.dev.yml up 
```

For correct presigned url work you should also edit /etc/hosts file and add minio like shown below

```console
127.0.0.1 localhost minio
```

### 2. Run with AWS

For run streaming-service with AWS you need AWS account with full acces to S3.

You need to export environment variables with your AWS credentials

```console
export AWS_ACCESS_KEY
export AWS_SECRET_ACCESS_KEY
export AWS_REGION
```

Build streaming-service

```console
make
```

Usage

```console
Usage of ./streaming-service:
  -b string
        bucket name (default "streaming-service-bucket")
  -dev
        development mode
  -p string
        server port (default ":8080")
  -t duration
        presigned url expire time (default 10m0s)
```
