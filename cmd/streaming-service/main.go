package main

import (
	"log"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/mboldysh/streaming-service/internal/router/userrouter"
	"github.com/mboldysh/streaming-service/pkg/metrics"

	"github.com/mboldysh/streaming-service/internal/config"
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

	userrouter := userrouter.New(trackService)

	server := server.New(cfg.ListenAddr)

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	server.InitMiddleware(middleware.Logger, cors.Handler)

	server.InitRoutes(
		userrouter,
		metrics.HealthCheck(),
	)

	server.Run()
}
