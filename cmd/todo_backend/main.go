package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/satoshi-tahara-st/todo_backend/pkg/application"
)

func main() {
	stage := os.Getenv("STAGE")
	product := os.Getenv(("PRODUCT"))
	conf, err := application.NewConfig(fmt.Sprintf("configs/%s/%s/setting.yaml", stage, product))
	if err != nil {
		log.Fatalf("configのロード中にエラーが発生しました: %v\n", err)
	}

	e := NewRouter(*conf)
	port := "1325"

	// Graceful shutdown のため、interrupt signalをキャッチして即時終了しないようにする
	// https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/notify-with-context/server.go
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := e.Start(port); err != nil && err != http.ErrServerClosed {
			// logger.Log.Fatal("Failed to listen and serve", "err", err)
		}
	}()

	<-ctx.Done()
	stop()
	// logger.Log.Debug("Shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		// logger.Log.Fatal("Server forced to shutdown", "err", err)
	}

	// logger.Log.Debug("Server exiting")
}
