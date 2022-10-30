package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/nvasilev98/rentals/cmd/rentals/internal/rentals"
	"github.com/nvasilev98/rentals/pkg/repository/postgres"
	r "github.com/nvasilev98/rentals/pkg/repository/postgres/rentals"
)

func main() {
	config, err := postgres.LoadDBConfig()
	if err != nil {
		log.Fatal(err)
	}

	dbClient, err := postgres.Connect(config)
	if err != nil {
		log.Fatal(err)
	}

	rentalsRepository, err := r.NewRepository(dbClient)
	if err != nil {
		log.Fatal(err)
	}

	handler := gin.Default()
	presenter := rentals.NewPresenter(rentalsRepository)

	handler.GET("/rentals/:id", presenter.RetrieveRentalByID)
	handler.GET("/rentals", presenter.RetrieveRentals)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: handler,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			if err != nil {
				log.Fatal(err)
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

	shutdownCtx, shutdownCancelFunc := context.WithTimeout(context.Background(), 1000000)
	defer shutdownCancelFunc()
	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Fatal("failed to gracefully shutdown http server")
	}

	if err := rentalsRepository.Close(); err != nil {
		log.Fatal(err)
	}

}
