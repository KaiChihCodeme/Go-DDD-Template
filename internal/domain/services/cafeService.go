package services

import (
	"errors"

	"kaichihcodeme.com/go-template/internal/domain/models"
	serviceProvidersDb "kaichihcodeme.com/go-template/internal/domain/serviceProviders/db"
)

type CafeService struct {
	MysqlRepository serviceProvidersDb.IMysqlRepositories
}

func NewCafeService(mysqlRepository serviceProvidersDb.IMysqlRepositories) *CafeService {
	return &CafeService{
		MysqlRepository: mysqlRepository,
	}
}

func (s *CafeService) GetCafe(cafeRequest models.GetCafeRequest) (*models.Cafe, error) {
	if cafeRequest.Name == "" {
		return nil, errors.New("cafe name is required")
	}

	cafe, err := s.MysqlRepository.GetCafe(cafeRequest)
	if err != nil {
		return nil, err
	}

	return cafe, nil
}
