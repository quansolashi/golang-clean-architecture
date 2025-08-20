package cmd

import (
	"clean-architecture/internal/config"
	dlogger "clean-architecture/internal/domain/service/logger"
	"clean-architecture/internal/infrastructure/logger"
	"clean-architecture/internal/interfaces/http"
	"clean-architecture/internal/interfaces/http/controller"
	"clean-architecture/internal/interfaces/http/middleware"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

type app struct {
	authentication *middleware.AuthMiddleware
	controller     *controller.Controller
	env            *config.Environment
	logger         logger.Client
}

func Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app := &app{}
	err := app.inject(ctx)
	if err != nil {
		return err
	}

	router := app.newRouter()
	server := http.NewHTTPServer(router, app.env.Port)

	eg, ectx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		if err := server.Serve(); err != nil {
			return nil
		}
		return nil
	})

	app.logger.Info("Server started", dlogger.Int("port", int64(app.env.Port)))

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ectx.Done():
		app.logger.Error("Server stopped", dlogger.Error(ectx.Err()))
	case signal := <-quit:
		app.logger.Info("Shutdown Server ...", dlogger.String("signal", signal.String()))
		delay := time.Duration(5) * time.Second
		time.Sleep(delay)
	}

	if err := server.Stop(ctx); err != nil {
		app.logger.Error("Failed to stopped http server", dlogger.Error(err))
		return err
	}
	return eg.Wait()
}
