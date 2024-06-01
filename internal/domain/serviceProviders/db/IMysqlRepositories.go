package serviceProvidersDb

import "github.com/KaiChihCodeme/Go-DDD-Template/internal/domain/models"

type IMysqlRepositories interface {
	GetCafe(cafeRequst models.GetCafeRequest) (*models.Cafe, error)
}
