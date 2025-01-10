package task

import (
	"net/http"
	"to-do-service/internal/app/usecase/task/getall"

	"github.com/gin-gonic/gin"
	"github.com/madevara24/go-common/response"
)

func GetAll(inport getall.Inport) gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := inport.Execute(c.Copy().Request.Context())
		if err != nil {
			response.WriteError(c, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "successfully get task",
			"data":    res,
		})
	}
}
