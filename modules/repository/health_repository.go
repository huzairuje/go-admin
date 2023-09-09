package repository

import (
	"context"
	"gorm.io/gorm"
)

type HealthRepository struct {
	db *gorm.DB
}

func NewHealthRepository(db *gorm.DB) HealthRepository {
	return HealthRepository{
		db: db,
	}
}

func (r HealthRepository) CheckUpTimeDB(ctx context.Context) (err error) {
	db, err := r.db.WithContext(ctx).DB()
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return
	}

	return
}
