package routes

import (
	"github.com/gin-gonic/gin"

	"kaichihcodeme.com/go-template/internal/application/controllers"
	"kaichihcodeme.com/go-template/internal/domain/services"
	"kaichihcodeme.com/go-template/internal/infra/config"
	dbInfraConn "kaichihcodeme.com/go-template/internal/infra/db/connections"
	dbInfra "kaichihcodeme.com/go-template/internal/infra/db/repositories"
	logger "kaichihcodeme.com/go-template/pkg/zap-logger"
)

type Provider struct {
	cafeController *controllers.CafeController
}

func newDependencyProviders() *Provider {
	// new all dependencies provider/services/repositories
	// then return Provider struct
	mysqlConnectionService := dbInfraConn.NewMysqlConnectionService(&config.MySQL, &config.Env)
	mysqlRepository := dbInfra.NewMysqlRepositories(mysqlConnectionService)
	cafeService := services.NewCafeService(mysqlRepository)
	cafeController := controllers.NewCafeController(cafeService)

	return &Provider{
		cafeController: cafeController,
	}
}

func RegisterApiRoutes(router *gin.Engine) {
	providers := newDependencyProviders()

	// add routes of the following router
	router.GET("/api/v1/cafe", func(ctx *gin.Context) {
		logger.Info("Info",
			logger.String("url", ctx.Request.URL.Path))

		providers.cafeController.GetCafe(ctx)
	})
}
