package task

import (
	"net/http"
	"to-do-service/internal/app/usecase/task/getbyuuid"

	"github.com/gin-gonic/gin"
	"github.com/madevara24/go-common/response"
)

func GetByUUID(inport getbyuuid.Inport) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := getbyuuid.InportRequest{}
		req.UUID = c.Param("uuid")

		res, err := inport.Execute(c.Copy().Request.Context(), req)
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
