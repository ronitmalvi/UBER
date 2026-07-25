package main

import (
	"context"      //Used to control cancellation and timeout.
	"log"          //Prints logs.
	"net/http"     //Provides the HTTP server.
	"os"           //Provides access to operating system features.
	"os/signal"    //Listens for OS signals. Ctrl+C  kill command  Docker stop
	"syscall"      //Contains signal constants like SIGINT SIGTERM
	"time"         //Used for timeout.
	"github.com/ronitmalvi/uber-backend/internal/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"github.com/ronitmalvi/uber-backend/internal/config"        //Loads environment variables.
	"github.com/ronitmalvi/uber-backend/internal/routes"
)

func main() {

	utils.InitLogger()
	defer utils.Logger.Sync()                //Logger.Sync() flushes buffered log entries before the application exits.

	cfg := config.Load()

	router := gin.Default()
	routes.Register(router,cfg)

	server :=&http.Server{
		Addr: ":"+cfg.ServerPort,
		Handler: router,
	}

	go func() {
		utils.Logger.Info(
			"Server started",
			zap.String("application", cfg.AppName),
			zap.String("port", cfg.ServerPort),
		)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)   //Creates a channel. A channel is used for communication between goroutines. Here it stores operating system signals.

	signal.Notify(
		quit,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-quit                           //This blocks the main goroutine until a signal arrives.

	utils.Logger.Info("Shutdown signal received")

	ctx, cancel := context.WithTimeout(                          //This creates a context with a 5-second deadline.
		context.Background(),                                    //The server gets up to 5 seconds to finish any in-progress requests.
		5*time.Second,
	)
	defer cancel()                                                //to release resources associated with the context.

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")

}

