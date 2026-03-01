package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/khy20040121/orbit-layout-advanced/pkg/jwt"
	"github.com/khy20040121/orbit-layout-advanced/pkg/log"
)

type Handler struct {
	logger *log.Logger
}

func NewHandler(
	logger *log.Logger,
) *Handler {
	return &Handler{
		logger: logger,
	}
}
func GetUserIdFromCtx(ctx *gin.Context) string {
	v, exists := ctx.Get("claims")
	if !exists {
		return ""
	}
	return v.(*jwt.MyCustomClaims).UserId
}
