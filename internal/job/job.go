package job

import (
	"github.com/khy20040121/orbit-layout-advanced/internal/repository"
	"github.com/khy20040121/orbit-layout-advanced/pkg/jwt"
	"github.com/khy20040121/orbit-layout-advanced/pkg/log"
	"github.com/khy20040121/orbit-layout-advanced/pkg/sid"
)

type Job struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *jwt.JWT
	tm     repository.Transaction
}

func NewJob(
	tm repository.Transaction,
	logger *log.Logger,
	sid *sid.Sid,
) *Job {
	return &Job{
		logger: logger,
		sid:    sid,
		tm:     tm,
	}
}
