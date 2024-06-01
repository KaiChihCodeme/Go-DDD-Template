package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/KaiChihCodeme/Go-DDD-Template/internal/application/models/requests"
	"github.com/KaiChihCodeme/Go-DDD-Template/internal/application/models/responses"
	"github.com/KaiChihCodeme/Go-DDD-Template/internal/domain/services"
	logger "github.com/KaiChihCodeme/Go-DDD-Template/pkg/zap-logger"
)

// service injection
type CafeController struct {
	CafeService services.CafeService
}

func NewCafeController(cafeService *services.CafeService) *CafeController {
	return &CafeController{
		CafeService: *cafeService,
	}
}

//	@Summary	Get cafe information
//	@Tags		Cafe
//	@version	1.0
//	@produce	application/json
//	@Param		name	query		string						true	"name"
//	@Success	200		{string} string "ok"
//	@Router		/api/v1/cafe [get]
func (f *CafeController) GetCafe(ctx *gin.Context) {
	var request requests.CafeRequest

	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.ErrorStacks("Error",
			logger.String("url", ctx.Request.URL.Path),
			logger.String("Error", err.Error()))

		return
	}

	cafe, err := f.CafeService.GetCafe(request.ToDomain())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.ErrorStacks("Error",
			logger.String("url", ctx.Request.URL.Path),
			logger.String("Error", err.Error()))

		return
	}

	// response model transform
	response := responses.GetCafeResponse{
		Uid:     cafe.Uid,
		Name:    cafe.Name,
		Address: cafe.Address,
	}

	ctx.JSON(http.StatusOK, response)
}

func (f *CafeController) CreateCafe(c *gin.Context) {

}
