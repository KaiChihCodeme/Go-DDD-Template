package init

import (
	"fmt"
	"net/http"
	"os"

	logger "kaichihcodeme.com/go-template/pkg/zap-logger"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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
	viper.SetDefault("mysql.cmd_timeout", 180)
	viper.SetDefault("mysql.max_idle_conns", 3)
	viper.SetDefault("mysql.max_open_conns", 10)

	// set the version
	var version string
	env := viper.GetString("env")
	if env == "" {
		env = "local"
	} else {
		version = viper.GetString("version")
	}

	// set to config global vars model

	// initialize the log
	sync = logger.InitLogger(logger.InfoLevel,
		logger.String("version", version),
		logger.String("process_name", envPrefix))

	// init Gin and routes
	gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.New()
	// setup Middleware here
	// router.Use(Middleware)

	// regist the routes
	// routes.RegisterApiRoutes(router)

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
