package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	sysinit "kaichihcodeme.com/go-template/init"
	logger "kaichihcodeme.com/go-template/pkg/zap-logger"
)

func main() {
	srvr := sysinit.New()

	defer func() {
		if r := recover(); r != nil {
			logger.ErrorStacks("backend server got error from recove when initializing", logger.String("Error", logger.TransformToString(r)))
		}

		sysinit.Close()
	}()

	logger.Info("Stating the server...")

	go func() {
		if err := srvr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.ErrorStacks("backend server got error from start", logger.String("Error", err.Error()))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	c := <-quit
	logger.Info("Shutting down server...", logger.String("signal", fmt.Sprintf("signal: %v", c)))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srvr.Shutdown(ctx); err != nil {
		logger.ErrorStacks("backend server got error from shutdown", logger.String("Error", err.Error()))
	}

	select {
	case <-ctx.Done():
		logger.Info("timeout of 10 seconds.")
	}
	logger.Info("Server exiting")
}
