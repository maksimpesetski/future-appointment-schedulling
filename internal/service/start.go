package service

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func (s *Service) Start() {
	// listen to shutdown from the listen thread, before exiting the main thread
	shutDownChan := make(chan bool, 2)

	// listen to the appropriate signals, and notify a channel
	stopChan := make(chan os.Signal, 10)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	server := &http.Server{
		Addr:    ":8080",
		Handler: s.router,
	}

	server.RegisterOnShutdown(func() {
		transport, ok := http.DefaultTransport.(*http.Transport)
		if !ok {
			panic("Cannot cast http.DefaultTransport to *http.Transport")
		}
		transport.DisableKeepAlives = true
		transport.CloseIdleConnections()
		server.SetKeepAlivesEnabled(false)
		s.logger.Info("register on shutdown completed")
	})

	go func() {
		s.logger.Info("serving HTTP", zap.Int("Port:", 8080))

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatal("serverError", zap.Error(err))
		}

		s.logger.Info("server exited")
		s.logger.Sync() // flushes the log buffer
		shutDownChan <- true
	}()

	<-stopChan // wait for a signal to exit
	s.logger.Info("shutting down HTTP server")

	defer s.logger.Sync() // flushes the log buffer

	// shutdown the server by gracefully draining connections
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		shutDownChan <- true
		s.logger.Fatal("shutdown error", zap.Error(err))
	}

	<-shutDownChan
	s.logger.Info("shutdown complete")
}
