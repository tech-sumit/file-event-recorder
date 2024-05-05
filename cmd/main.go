package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"os/signal"
	"syscall"
	"testCode/internal/repository"
	"testCode/internal/service"
	"testCode/internal/utils/taskqueue"
	"testCode/internal/worker"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"testCode/internal/repository/env"
	testCodeSql "testCode/internal/repository/sql"
)

var config *env.Config
var srv *service.Config
var wkr *worker.Config

func init() {
	var err error
	// Init Env to config struct
	if config, err = env.InitEnvConfig(); err != nil {
		panic(err)
	}

	// Setup worker queue
	taskqueue.Setup(config.MaxConcurrency)

	// Setup Database
	db, err := sql.Open(repository.DATABASE_DRIVER, config.DBPath)
	if err != nil {
		panic(err)
	}

	// Setup file processor service
	srv = service.New(testCodeSql.New(db))

	// Setup worker to collect file events
	wkr = worker.New(srv, config.DirPath)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start watcher
	if err := wkr.StartWatcher(ctx); err != nil {
		panic(err)
	}

	// Setup channel to listen for SIGINT signals.
	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan, syscall.SIGINT)

	shutdownInitiated := false

	// Start a goroutine to handle graceful shutdown.
	go func() {
		<-sigChan
		if !shutdownInitiated {
			shutdownInitiated = true
			fmt.Println("\nMAIN: SIGINT received, shutting down... Press CTRL+C again to force exit.")
			go func() {
				// Initiate graceful shutdown procedures
				cancel()
				taskqueue.Shutdown()
				time.Sleep(5 * time.Second) // Simulate delay in shutdown
				fmt.Println("MAIN: Shutdown complete.")
				os.Exit(0) // Exit cleanly after shutdown is complete
			}()
		} else {
			fmt.Println("\nMAIN: Forced shutdown initiated.")
			os.Exit(1) // Exit immediately with a non-zero status
		}
	}()

	// Print to indicate that the service is running
	fmt.Println("MAIN: Service is running. Press CTRL+C to stop.")

	// Prevent the main function from exiting immediately
	select {} // Block forever until a signal is received
}
