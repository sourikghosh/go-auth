package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"auth/implementation/auth"
	"auth/pkg"
	"auth/pkg/config"
	"auth/pkg/logger"
	"auth/transport/endpoints"
	transportHttp "auth/transport/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// loading the configuration
	cfg, err := config.LoadConfig(".")
	if err != nil {
		fmt.Printf("failed to load configuration: %s\n", err.Error())
		os.Exit(1)
	}

	// creating new logger from config
	l := logger.NewLogger(cfg)

	// creating mysql client
	conn, err := gorm.Open(mysql.Open(cfg.DB_DSN), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		l.Error(err, "db connection failed")
		os.Exit(1)
	}

	// initializing all services
	jwtSvc := pkg.NewJWTService(*cfg, l)
	repo := auth.NewRepository(conn, l)
	srv := auth.NewService(l, repo, *cfg, jwtSvc)

	end := endpoints.MakeEndpoints(srv, l)
	handler := transportHttp.NewHTTPService(end, jwtSvc)

	// creating server with timeout and assigning the routes
	server := &http.Server{
		Addr:         ":" + cfg.PORT,
		ReadTimeout:  config.HttpTimeOut,
		WriteTimeout: config.HttpTimeOut,
		IdleTimeout:  config.HttpTimeOut,
		Handler:      handler,
	}

	// start listening and serving http server
	go func() {
		l.Info("🚀 HTTP server running on PORT:" + cfg.PORT)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			l.Error(err, "err occurred:")
		}
	}()

	// listening for system events to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	l.Info("Signal received to Shutdown server...")

	ctxWithTimeOut, cancel := context.WithTimeout(context.Background(), config.ServerShutdownTimeOut)
	defer cancel()

	// gracefully shutdown http server
	if err := server.Shutdown(ctxWithTimeOut); err != nil {
		cancel()
		l.V(1).Error(err, "Server forced to shutdown:")
	}

	l.Info("application exited")
}
