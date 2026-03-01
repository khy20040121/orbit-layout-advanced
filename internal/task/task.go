package task

import (
	"github.com/khy20040121/orbit-layout-advanced/internal/repository"
	"github.com/khy20040121/orbit-layout-advanced/pkg/jwt"
	"github.com/khy20040121/orbit-layout-advanced/pkg/log"
	"github.com/khy20040121/orbit-layout-advanced/pkg/sid"
)

type Task struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *jwt.JWT
	tm     repository.Transaction
}

func NewTask(
	tm repository.Transaction,
	logger *log.Logger,
	sid *sid.Sid,
) *Task {
	return &Task{
		logger: logger,
		sid:    sid,
		tm:     tm,
	}
}
