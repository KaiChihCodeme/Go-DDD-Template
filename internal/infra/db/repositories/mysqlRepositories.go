package dbInfra

import (
	"github.com/KaiChihCodeme/Go-DDD-Template/internal/domain/models"
	serviceProvidersDb "github.com/KaiChihCodeme/Go-DDD-Template/internal/domain/serviceProviders/db"
	sqlConst "github.com/KaiChihCodeme/Go-DDD-Template/internal/infra/sql"
	logger "github.com/KaiChihCodeme/Go-DDD-Template/pkg/zap-logger"

	"errors"
)

type MysqlRepositories struct {
	mysqlClient serviceProvidersDb.IDbConnector
}

func NewMysqlRepositories(mysqlClient serviceProvidersDb.IDbConnector) serviceProvidersDb.IMysqlRepositories {
	return &MysqlRepositories{
		mysqlClient: mysqlClient,
	}
}

func (m *MysqlRepositories) GetCafe(cafeRequest models.GetCafeRequest) (*models.Cafe, error) {
	dbClient, err := m.mysqlClient.GetConnection()
	if err != nil {
		return nil, err
	}

	defer dbClient.Close()

	queryString := sqlConst.GetCafeByName

	rows, err := dbClient.Query(queryString, cafeRequest.Name)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	defer func() {
		if err := rows.Close(); err != nil {
			logger.ErrorStacks("Error",
				logger.String("reason", "Failed to close DB connection"),
				logger.String("Error", err.Error()))
		}
	}()

	var cafes []models.Cafe
	for rows.Next() {
		var cafe models.Cafe
		if err := rows.Scan(&cafe.Uid, &cafe.Name, &cafe.Address); err != nil {
			logger.ErrorStacks("Error",
				logger.String("reason", "failed to scan DB results"),
				logger.String("Error", err.Error()))
			return nil, err
		}

		cafes = append(cafes, cafe)
	}

	if err := rows.Err(); err != nil {
		logger.ErrorStacks("Error",
			logger.String("reason", "failed to scan DB results"),
			logger.String("Error", err.Error()))
		return nil, err
	}

	if len(cafes) == 0 {
		logger.ErrorStacks("Error",
			logger.String("reason", "failed to scan DB results"))

		return nil, errors.New("cafe not found")
	}

	return &cafes[0], nil
}
