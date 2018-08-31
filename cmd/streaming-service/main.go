package main

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/s3manager"
	"github.com/mboldysh/streaming-service/internal/config"
	"github.com/mboldysh/streaming-service/internal/router"
	"github.com/mboldysh/streaming-service/internal/router/trackhandler"
	"github.com/mboldysh/streaming-service/internal/server"
	"github.com/mboldysh/streaming-service/internal/service/trackservice"
	store "github.com/mboldysh/streaming-service/internal/store/s3"
)

func main() {
	cfg, err := config.New()

	if err != nil {
		log.Fatal(err.Error())
	}

	s3Client := s3.New(cfg.AWSCfg)
	s3Uploader := s3manager.NewUploaderWithClient(s3Client)

	trackStore, err := store.NewTrackStore(s3Client, s3Uploader, cfg.TrackBucketName, cfg.PresignExpireTime)

	if err != nil {
		log.Fatal(err.Error())
	}

	trackService := trackservice.New(trackStore)

	trackHandler := trackhandler.New(trackService)

	r := router.Init(
		trackHandler,
	)

	server := server.NewServer(cfg.ListenAddr, r)
	server.Run()
}
