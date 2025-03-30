package auth

import (
	"net/http"
	"try-graphql/internal/app/usecase/auth/login"

	"github.com/gin-gonic/gin"
	"github.com/madevara24/go-common/request"
	"github.com/madevara24/go-common/response"
)

func Login(inport login.Inport) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		req := login.InportRequest{}

		if err := request.UnmarshalJSON(c, &req); err != nil {
			return
		}

		token, err := inport.Execute(ctx, req)
		if err != nil {
			response.WriteError(c, err)
			return
		}

		c.JSON(http.StatusOK, response.BasePayload{
			Success: true,
			Data:    token,
		})
	}
}
