package dbInfra

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"

	serviceProvidersDb "kaichihcodeme.com/go-template/internal/domain/serviceProviders/db"
	"kaichihcodeme.com/go-template/internal/infra/config"
	logger "kaichihcodeme.com/go-template/pkg/zap-logger"
)

type MysqlConnectionService struct {
	config *config.MySQLConfig
	env    *string
}

func NewMysqlConnectionService(config *config.MySQLConfig, env *string) serviceProvidersDb.IDbConnector {
	return &MysqlConnectionService{
		config: config,
		env:    env,
	}
}

func (m *MysqlConnectionService) GetConnection() (*sql.DB, error) {
	connectionString, err := m.GetConnectionString()
	if err != nil {
		logger.ErrorStacks("Error",
			logger.String("Error", err.Error()))

		return nil, err
	}

	db, err := sql.Open("mysql", *connectionString)
	if err != nil {
		logger.ErrorStacks("Error",
			logger.String("Error", err.Error()))

		return nil, err
	}

	// Setting the connection
	db.SetMaxIdleConns(config.MySQL.MaxIdleConns)
	db.SetMaxOpenConns(config.MySQL.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(config.MySQL.ConnTimeout) * time.Minute)

	return db, nil
}

func (m *MysqlConnectionService) GetConnectionString() (*string, error) {
	// structure the connection string
	if *m.env == "dev" || *m.env == "local" {
		// no need to decrypt the connection string
		return &m.config.ConnectionString, nil
	} else {
		// need to decrypt the connection string (TODO: Need to implement decrypt algorithm)
		// return m.config.ConnectionString, nil
	}

	return nil, nil

}
