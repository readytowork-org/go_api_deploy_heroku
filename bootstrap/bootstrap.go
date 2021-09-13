package bootstrap

import (
	"context"

	"go_api_deploy_heroku/api/controllers"
	"go_api_deploy_heroku/api/middlewares"
	"go_api_deploy_heroku/api/routes"
	"go_api_deploy_heroku/lib"
	"go_api_deploy_heroku/repository"
	"go_api_deploy_heroku/services"

	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	controllers.Module,
	routes.Module,
	lib.Module,
	services.Module,
	middlewares.Module,
	repository.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler lib.RequestHandler,
	routes routes.Routes,
	env lib.Env,
	logger lib.Logger,
	middlewares middlewares.Middlewares,
	database lib.Database,
) {
	conn, _ := database.DB.DB()

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("Starting Application")
			logger.Info("---------------------")
			logger.Info("------- CLEAN -------")
			logger.Info("---------------------")

			conn.SetMaxOpenConns(10)
			go func() {
				middlewares.Setup()
				routes.Setup()
				handler.Gin.Run(":" + env.ServerPort)
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			logger.Info("Stopping Application")
			conn.Close()
			return nil
		},
	})
}
