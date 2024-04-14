package config

type MySQLConfig struct {
	AppName          string
	ConnTimeout      int
	MaxIdleConns     int
	MaxOpenConns     int
	ConnectionString string
}

var (
	MySQL MySQLConfig
	Env   string
)
