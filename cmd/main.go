package main

import (
	"net"
	"net/http"
	"os"

	"Alhanaqtah/habr-backend/internal/article"
	articleRepo "Alhanaqtah/habr-backend/internal/article/repository"
	"Alhanaqtah/habr-backend/internal/config"
	"Alhanaqtah/habr-backend/internal/user"
	userRepo "Alhanaqtah/habr-backend/internal/user/repository"
	"Alhanaqtah/habr-backend/pkg/client/postgres"
	"Alhanaqtah/habr-backend/pkg/env"
	"Alhanaqtah/habr-backend/pkg/logger"

	"github.com/go-chi/chi"
)

func main() {
	// Config & env-vars
	cfg := config.MustLoad()
	env.Load()

	// Logger
	log := logger.New(cfg.Env)

	log.Info("Initializing server")

	log.Debug("Initializing connection with database")

	// Database client
	pgClient, _ := postgres.NewClient(
		os.Getenv("PG_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	articleRepo := articleRepo.New(pgClient, log)
	userRepo := userRepo.New(pgClient, log)

	log.Debug("Initializing router")

	// Router
	router := chi.NewRouter()

	articles := article.NewHandler(articleRepo, log)
	articles.Register(router)

	users := user.NewHandler(userRepo, log)
	users.Register(router)

	// Server

	srv := &http.Server{
		Handler:      router,
		ReadTimeout:  cfg.HttpServer.Timeout,
		WriteTimeout: cfg.HttpServer.Timeout,
		IdleTimeout:  cfg.HttpServer.IdleTimeout,
	}

	listener, err := net.Listen("tcp", "localhost:"+cfg.HttpServer.Port)
	if err != nil {
		log.Error("failed to create listener: ", err.Error())
	}

	log.Info("Server is listening...")

	log.Error(srv.Serve(listener).Error())
}
