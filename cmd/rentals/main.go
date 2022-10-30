package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/nvasilev98/rentals/cmd/rentals/env"
	"github.com/nvasilev98/rentals/cmd/rentals/internal/rentals"
	"github.com/nvasilev98/rentals/pkg/repository/postgres"
	r "github.com/nvasilev98/rentals/pkg/repository/postgres/rentals"
	"github.com/sirupsen/logrus"
)

const timeout = 1000000

func main() {
	appConfig, err := env.LoadAppConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	dbConfig, err := postgres.LoadDBConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	dbClient, err := postgres.Connect(dbConfig)
	if err != nil {
		logrus.Fatal(err)
	}

	rentalsRepository, err := r.NewRepository(dbClient)
	if err != nil {
		logrus.Fatal(err)
	}

	handler := gin.Default()
	presenter := rentals.NewPresenter(rentalsRepository)

	handler.GET("/rentals/:id", presenter.RetrieveRentalByID)
	handler.GET("/rentals", presenter.RetrieveRentals)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", appConfig.Host, appConfig.Port),
		Handler: handler,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			if err != nil {
				logrus.Fatal(err)
			}
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	<-sigChan
	signal.Stop(sigChan)

	shutdownCtx, shutdownCancelFunc := context.WithTimeout(context.Background(), timeout)
	defer shutdownCancelFunc()
	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		logrus.Fatal("failed to gracefully shutdown http server")
	}

	if err := rentalsRepository.Close(); err != nil {
		logrus.Fatal(err)
	}

}
