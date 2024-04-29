package init

import (
	"fmt"
	"net/http"
	"os"

	"kaichihcodeme.com/go-template/internal/application/middlewares"
	"kaichihcodeme.com/go-template/internal/application/routes"
	"kaichihcodeme.com/go-template/internal/infra/config"
	logger "kaichihcodeme.com/go-template/pkg/zap-logger"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	router *gin.Engine
	sync   func() error
)

func New() *http.Server {
	envPrefix := "Go-Template"
	configAddr := "configs/go-template.json"

	// Load configuration via Viper
	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()
	viper.SetConfigFile(configAddr)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error when getting from config file: %w", err))
	}

	// set default config values
	viper.SetDefault("mysql.app_name", "Go-Template")
	viper.SetDefault("mysql.conn_timeout", 30)
	viper.SetDefault("mysql.max_idle_conns", 3)
	viper.SetDefault("mysql.max_open_conns", 10)

	// set the version
	// I only give Go-Template_ENV in the docker-compose file, so in the local environment,
	// env should be empty.
	var version string
	env := viper.GetString("env")
	if env == "" {
		env = "local"
	} else {
		version = viper.GetString("version")
	}

	// get db setting
	// If env is local, use the connection string defined in the config file
	// Otherwise, use the connection string defined in the environment variable
	if env != "local" {
		mysqlHost := viper.GetString("MYSQL_HOST")
		mysqlDatabase := viper.GetString("MYSQL_DATABASE")
		mysqlUser := viper.GetString("MYSQL_USER")
		mysqlPassword := viper.GetString("MYSQL_PASSWORD")
		// form a connection string to mysql
		viper.Set("mysql.connection_string", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", mysqlUser, mysqlPassword, mysqlHost, mysqlDatabase))
	}

	// set to config global vars model
	config.MySQL.ConnectionString = viper.GetString("mysql.connection_string")
	config.MySQL.AppName = viper.GetString("mysql.app_name")
	config.MySQL.ConnTimeout = viper.GetInt("mysql.conn_timeout")
	config.MySQL.MaxIdleConns = viper.GetInt("mysql.max_idle_conns")
	config.MySQL.MaxOpenConns = viper.GetInt("mysql.max_open_conns")

	config.Env = env

	// initialize the log
	sync = logger.InitLogger(logger.InfoLevel,
		logger.String("version", version),
		logger.String("process_name", envPrefix))

	// init Gin and routes
	gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.New()
	// setup Middleware here
	router.Use(middlewares.ErrorHandler())

	// regist the routes
	routes.RegisterApiRoutes(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
}

func Close() {
	// DB connection close
	// mysql close

	// log sync over
	sync()
}
