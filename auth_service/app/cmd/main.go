package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/GermanBogatov/auth_service/internal/config"
	"github.com/GermanBogatov/auth_service/internal/handler"
	"github.com/GermanBogatov/auth_service/internal/service"
	"github.com/GermanBogatov/auth_service/internal/storage"
	"github.com/GermanBogatov/auth_service/pkg/jwt"
	"github.com/GermanBogatov/auth_service/pkg/logging"
	"github.com/GermanBogatov/auth_service/pkg/postgresql"
	"github.com/GermanBogatov/auth_service/pkg/redis"
	"github.com/GermanBogatov/auth_service/pkg/shutdown"
	"net"
	"net/http"
	"os"
	"syscall"
	"time"
)

func main() {

	logging.Init()
	logger := logging.GetLogger()
	logger.Println("logger initialized...")

	logger.Println(" config initializing...")
	cfg := config.GetConfig()

	logger.Println("Redis-client initializing...")
	RedisClient, err := redis.NewClient(cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.Password, cfg.Redis.DB)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Println("JWT Helper initializing...")
	NewHelper := jwt.NewHelper(logger, RedisClient)

	logger.Println("Postgresql-client initializing...")
	PostgresqlClient, err := postgresql.NewClient(context.Background(), 5, cfg.PostgresqlDB.Username, cfg.PostgresqlDB.Password,
		cfg.PostgresqlDB.Host, cfg.PostgresqlDB.Port, cfg.PostgresqlDB.Database)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Println("Storage initializing...")
	Storage := storage.NewStorage(PostgresqlClient, logger)
	if err != nil {
		panic(err)
	}

	logger.Println("Service initializing...")
	Service, err := service.NewService(Storage, logger)
	if err != nil {
		panic(err)
	}

	logger.Println("Handler initializing...")

	Handler, err := handler.NewHandler(Service, logger, NewHelper)
	if err != nil {
		panic(err)
	}

	logger.Println("start application...")
	start(Handler.InitRoutes(), logger, cfg)
}

func start(router http.Handler, logger logging.Logger, cfg *config.Config) {
	var server *http.Server
	var listener net.Listener

	logger.Infof("bind application to host: %s and port: %s", cfg.HTTP.IP, cfg.HTTP.Port)

	var err error

	listener, err = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.HTTP.IP, cfg.HTTP.Port))
	if err != nil {
		logger.Fatal(err)
	}

	server = &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go shutdown.Graceful([]os.Signal{syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGTERM}, server)

	logger.Println("Application initialized and started...")

	if err := server.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			logger.Warn("server shutdown")
		default:
			logger.Fatal(err)
		}
	}
}
