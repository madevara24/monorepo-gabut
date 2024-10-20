package task

import (
	"net/http"
	"to-do-app/internal/app/usecase/task/create"

	"github.com/gin-gonic/gin"
	"github.com/madevara24/go-common/request"
	"github.com/madevara24/go-common/response"
)

func Create(inport create.Inport) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := create.InportRequest{}

		if err := request.UnmarshalJSON(c, &req); err != nil {
			return
		}

		err := inport.Execute(c.Copy().Request.Context(), req)
		if err != nil {
			response.WriteError(c, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "successfully create task",
		})
	}
}
