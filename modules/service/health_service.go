package service

import (
	"context"
	"go-admin/infrastructure/logger"
	"go-admin/modules/repository"
)

type HealthService struct {
	healthRepository repository.HealthRepository
}

func NewHealthService(healthRepository repository.HealthRepository) HealthService {
	return HealthService{
		healthRepository: healthRepository,
	}
}

func (u *HealthService) CheckUpTime(ctx context.Context) error {
	ctxName := "CheckUpTime"
	errCheckDb := u.healthRepository.CheckUpTimeDB(ctx)
	if errCheckDb != nil {
		logger.Error(ctx, ctxName, "got error when %s : %v", ctxName, errCheckDb)
		return errCheckDb
	}

	return nil
}
