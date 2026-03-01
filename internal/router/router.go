package router

import (
	"github.com/khy20040121/orbit-layout-advanced/internal/handler"
	"github.com/khy20040121/orbit-layout-advanced/pkg/jwt"
	"github.com/khy20040121/orbit-layout-advanced/pkg/log"
	"github.com/spf13/viper"
)

type RouterDeps struct {
	Logger      *log.Logger
	Config      *viper.Viper
	JWT         *jwt.JWT
	UserHandler *handler.UserHandler
}
