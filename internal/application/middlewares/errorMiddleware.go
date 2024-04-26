package middlewares

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"kaichihcodeme.com/go-template/internal/domain/errors"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch e := err.(type) {
				case errors.BadRequestError:
					c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
				case errors.NotFoundError:
					c.JSON(http.StatusNotFound, gin.H{"error": e.Error()})
				case errors.ServerException:
					c.JSON(http.StatusInternalServerError, gin.H{"error": e.Error()})
				case errors.ApiServiceException:
					c.JSON(http.StatusInternalServerError, gin.H{"error": e.Error()})
				case errors.ApiServerError:
					bodyBytes, _ := io.ReadAll(e.Error().Body)
					e.Error().Body.Close()
					c.JSON(http.StatusInternalServerError, gin.H{"error": string(bodyBytes)})
				default:
					c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
				}
			}
		}()
		c.Next()
	}
}
