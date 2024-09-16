package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"tender-service/internal/config"
	"tender-service/internal/controller"
	"tender-service/internal/logger"
	"tender-service/internal/repository"
	"tender-service/internal/repository/repo"
	"tender-service/internal/routers"
)

type App struct {
	Router *chi.Mux
	repo   *repo.Repo
	server *http.Server
}

func New(ctx context.Context) (*App, error) {
	cfg, err := config.MustLoad()
	if err != nil {
		return nil, err
	}

	logger.Init(cfg.LogLevel)

	rp, err := repository.New(ctx, *cfg, cfg.StorageType)
	if err != nil {
		return nil, err
	}

	ctrl := controller.New(rp)

	r := routers.New(cfg.APIVersion, ctrl, cfg.ServerAddress)

	server := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: r,
	}

	a := &App{
		Router: r,
		repo:   rp,
		server: server,
	}

	return a, nil
}

func (a *App) Run() error {
	logger.Info("Starting app", "message", fmt.Sprintf("Listening on %s", a.server.Addr))

	err := a.server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		logger.Info("App stopped")
	}

	return err
}

func (a *App) Close(ctx context.Context, t time.Duration) {
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(ctx, t)
	defer cancel()

	if err := a.server.Shutdown(ctx); err != nil {
		logger.Error("failed to shutdown app", "error", err)
	}

	logger.Info("Shutting down app")
}
