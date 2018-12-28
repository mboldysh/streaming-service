# Streaming Service

[![Build Status](https://travis-ci.org/mboldysh/streaming-service.svg?branch=master)](https://travis-ci.org/mboldysh/streaming-service) 
[![Go Report Card](https://goreportcard.com/badge/github.com/mboldysh/streaming-service)](https://goreportcard.com/report/github.com/mboldysh/streaming-service)

Streaming Service is a service which allows easily share music across devices by storing it 
in a cloud and stream it. Developed as an alternative to traditional streaming 
services one of the main advantages of which besides a huge music library 
is ability to share music across devices. Ideal for those who have a local music 
collection and don't want to buy a subscription to the streaming service. Also, as service 
deployed decentralized, it's easy to share your server with your friends. Just send a link 
to your service to friends you want to invite and here it is! Use the same username to listen and 
modify one playlist or use different usernames to have individual playlists. This repository contains
the source code for the Streaming Service backend.

## Development Setup

There are two options:

### 1. Run with minio

```console
# build streaming-service
make
# build streaming-service docker image
docker build -t streaming-service .
# up streaming-service and minio
docker-compose -f docker-compose.dev.yml up 
```

Edit /etc/hosts file and add minio like shown below for correct presigned url work 

```console
127.0.0.1 localhost minio
```

### 2. Run with AWS S3

Prerequisites:

- AWS account with full acces to S3.
- Next environment variables are exported

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

## Deploy to AWS

1. To deploy Streaming Service you need to setup AWS CLI. Detailed information can be found
here: https://aws.amazon.com/cli/

2. Build docker image

```console
docker build -t streaming-service:latest .
```
3. Push docker image to container registry(e.g. Docker Hub or AWS ECR)

4. Configure deployment parameters. Open .env file in /cloudformation directory.
As S3 bucket name is globally unique, buckets name should be changed to unique ones.
IMAGE_ADDRESS should be set to the adress of the image which was pushed in previous step.
Optionally, REGION can be changed to preferable region

5. Deploy stack:

```console
# Open /cloudformation folder
cd cloudformation
# Deploy stack
make full-deploy
```
Upon completing these steps cloudformation stack will start to creating. Cloudformation creation
can be tracked in AWS Cloudformation console. After stack creation will be completed, server adress can be found
in stack outputs.