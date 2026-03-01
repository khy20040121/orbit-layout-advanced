//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/khy20040121/orbit-layout-advanced/internal/handler"
	"github.com/khy20040121/orbit-layout-advanced/internal/job"
	"github.com/khy20040121/orbit-layout-advanced/internal/repository"
	"github.com/khy20040121/orbit-layout-advanced/internal/router"
	"github.com/khy20040121/orbit-layout-advanced/internal/server"
	"github.com/khy20040121/orbit-layout-advanced/internal/service"
	"github.com/khy20040121/orbit-layout-advanced/pkg/app"
	"github.com/khy20040121/orbit-layout-advanced/pkg/jwt"
	"github.com/khy20040121/orbit-layout-advanced/pkg/log"
	"github.com/khy20040121/orbit-layout-advanced/pkg/server/http"
	"github.com/khy20040121/orbit-layout-advanced/pkg/sid"
	"github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	//repository.NewRedis,
	//repository.NewMongo,
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewUserRepository,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)

var jobSet = wire.NewSet(
	job.NewJob,
	job.NewUserJob,
)
var serverSet = wire.NewSet(
	server.NewHTTPServer,
	server.NewJobServer,
)

// build App
func newApp(
	httpServer *http.Server,
	jobServer *server.JobServer,
	// task *server.Task,
) *app.App {
	return app.NewApp(
		app.WithServer(httpServer, jobServer),
		app.WithName("demo-server"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		repositorySet,
		serviceSet,
		handlerSet,
		jobSet,
		serverSet,
		wire.Struct(new(router.RouterDeps), "*"),
		sid.NewSid,
		jwt.NewJwt,
		newApp,
	))
}
