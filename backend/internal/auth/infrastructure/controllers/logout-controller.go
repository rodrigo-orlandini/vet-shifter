package controllers

import (
	"net/http"

	_ "rodrigoorlandini/vet-shifter/internal/_shared/api"

	"github.com/gin-gonic/gin"
)

type LogoutController struct{}

func NewLogoutController() *LogoutController {
	return &LogoutController{}
}

// Logout godoc
//
//	@Summary		Logout (client should discard token)
//	@Description	Stateless JWT: server does not invalidate. Client must clear stored token.
//	@Tags			auth
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	"OK"
//	@Failure		401	{object}	api.ApiErrorResponse	"Unauthorized"
//	@Router			/auth/logout [post]
func (c *LogoutController) Handle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "logged out"})
}
