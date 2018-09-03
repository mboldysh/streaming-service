package main

import (
	"log"

	"github.com/mboldysh/streaming-service/internal/config"
	"github.com/mboldysh/streaming-service/internal/router"
	"github.com/mboldysh/streaming-service/internal/router/trackhandler"
	"github.com/mboldysh/streaming-service/internal/server"
	"github.com/mboldysh/streaming-service/internal/service/trackservice"
	"github.com/mboldysh/streaming-service/internal/store/s3"
)

func main() {
	cfg, err := config.New()

	if err != nil {
		log.Fatal(err.Error())
	}

	trackStore, err := s3.NewTrackStore(cfg.S3Clinet, cfg.S3uploader, cfg.TrackBucketName, cfg.PresignExpireTime)

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
