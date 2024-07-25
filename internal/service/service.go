package service

import (
	"gin-casbin-admin/internal/repository"
	"gin-casbin-admin/pkg/jwt"
	"gin-casbin-admin/pkg/log"
	"gin-casbin-admin/pkg/sid"
)

type Service struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *jwt.JWT
	tm     repository.Transaction
}

func NewService(
	tm repository.Transaction,
	logger *log.Logger,
	sid *sid.Sid,
	jwt *jwt.JWT,
) *Service {
	return &Service{
		logger: logger,
		sid:    sid,
		jwt:    jwt,
		tm:     tm,
	}
}
