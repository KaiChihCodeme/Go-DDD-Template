package requests

import "kaichihcodeme.com/go-template/internal/domain/models"

type CafeRequest struct {
	Name string `form:"name" binding:"required"`
}

func (c CafeRequest) ToDomain() models.GetCafeRequest {
	return models.GetCafeRequest{
		Name: c.Name,
	}
}
