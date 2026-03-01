//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/khy20040121/orbit-layout-advanced/internal/repository"
	"github.com/khy20040121/orbit-layout-advanced/internal/server"
	"github.com/khy20040121/orbit-layout-advanced/internal/task"
	"github.com/khy20040121/orbit-layout-advanced/pkg/app"
	"github.com/khy20040121/orbit-layout-advanced/pkg/log"
	"github.com/khy20040121/orbit-layout-advanced/pkg/sid"
	"github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	//repository.NewRedis,
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewUserRepository,
)

var taskSet = wire.NewSet(
	task.NewTask,
	task.NewUserTask,
)
var serverSet = wire.NewSet(
	server.NewTaskServer,
)

// build App
func newApp(
	task *server.TaskServer,
) *app.App {
	return app.NewApp(
		app.WithServer(task),
		app.WithName("demo-task"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		repositorySet,
		taskSet,
		serverSet,
		newApp,
		sid.NewSid,
	))
}
