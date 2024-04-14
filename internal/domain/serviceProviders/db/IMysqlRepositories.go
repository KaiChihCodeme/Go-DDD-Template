package serviceProvidersDb

import "kaichihcodeme.com/go-template/internal/domain/models"

type IMysqlRepositories interface {
	GetCafe(cafeRequst models.GetCafeRequest) (*models.Cafe, error)
}
