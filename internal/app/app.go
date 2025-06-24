package app

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/doganarif/govisual"
	"github.com/elibr-edu/gateway/internal/auth"
	"github.com/gin-gonic/gin"
)

type App struct {
	server *http.Server
	done   chan error
}

func gracefulShutdown(apiServer *http.Server, done chan error) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")
	stop() // Allow Ctrl+C to force shutdown

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
		done <- err
		return
	}

	log.Println("Server exiting")

	// Notify the main goroutine that the shutdown is complete
	done <- nil
}

func NewApp() *App {
	// Create a new Gin router
	router := gin.Default()

	// Register the routes
	RegisterRoutes(router)

	// Create a new HTTP server
	server := &http.Server{
		Addr: ":8080",
		Handler: govisual.Wrap(
			router,
			govisual.WithRequestBodyLogging(true),
			govisual.WithResponseBodyLogging(true),
			// govisual.WithOpenTelemetry(true),
		),
	}

	return &App{
		server: server,
		done:   make(chan error),
	}
}

func RegisterRoutes(router *gin.Engine) {
	// Register the Auth routes
	auth.RegisterRoutes(router.Group("/auth"))
}

func (a *App) Run() error {
	// Start the graceful shutdown
	go gracefulShutdown(a.server, a.done)

	// Start the HTTP server
	if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("listen: %s\n", err)
		return err
	}

	// Wait for the shutdown to complete
	return <-a.done
}
