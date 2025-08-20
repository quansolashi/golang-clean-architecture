package cmd

import (
	"clean-architecture/internal/application/usecase"
	"clean-architecture/internal/config"
	"clean-architecture/internal/infrastructure/auth/jwt"
	database "clean-architecture/internal/infrastructure/database/gorm"
	"clean-architecture/internal/infrastructure/logger"
	"clean-architecture/internal/infrastructure/repository"
	"clean-architecture/internal/interfaces/http/controller"
	"clean-architecture/internal/interfaces/http/middleware"
	"context"

	"gorm.io/gorm"
)

func (a *app) inject(ctx context.Context) error {
	/* env */
	env, err := a.loadEnv()
	if err != nil {
		return err
	}
	a.env = env

	/* logger */
	logger, err := logger.NewClient(logger.WithLevel(env.LogLevel))
	if err != nil {
		return err
	}
	a.logger = logger

	/* db */
	db, err := a.newDBClient()
	if err != nil {
		return err
	}

	/* use case dependencies */
	repo := repository.NewRepository(db)
	tokenService := jwt.NewJWT("secret")

	ucParams := &usecase.Params{
		Repository:   repo,
		TokenService: tokenService,
	}
	usecase := usecase.NewUsecase(ucParams)

	/* controller */
	controller := controller.NewController(usecase)
	a.controller = controller

	/* auth middleware */
	a.authentication = middleware.NewAuthMiddleware(tokenService)

	return nil
}

func (a *app) loadEnv() (*config.Environment, error) {
	config := config.NewClient()
	env, err := config.Load()
	if err != nil {
		return nil, err
	}
	return env, nil
}

func (a *app) newDBClient() (*gorm.DB, error) {
	params := &database.Params{
		Socket:   a.env.DBSocket,
		Host:     a.env.DBHost,
		Port:     a.env.DBPort,
		Database: a.env.DBDatabase,
		Username: a.env.DBUsername,
		Password: a.env.DBPassword,
	}
	db, err := database.NewDatabase(params)
	if err != nil {
		return nil, err
	}
	return db, nil
}
