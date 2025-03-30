package user

import (
	"net/http"
	"try-graphql/internal/app/usecase/user/register"

	"github.com/gin-gonic/gin"
	"github.com/madevara24/go-common/request"
	"github.com/madevara24/go-common/response"
)

func Register(inport register.Inport) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		req := register.InportRequest{}

		if err := request.UnmarshalJSON(c, &req); err != nil {
			return
		}

		err := inport.Execute(ctx, req)
		if err != nil {
			response.WriteError(c, err)
			return
		}

		c.JSON(http.StatusOK, response.BasePayload{
			Success: true,
			Message: "successfully register user",
		})
	}
}
