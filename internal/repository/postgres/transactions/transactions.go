package postgres

import (
	"context"

	"gorm.io/gorm"
)

type TxManager struct {
	db *gorm.DB
}

func New(db *gorm.DB) *TxManager {
	return &TxManager{
		db: db,
	}
}

func (r *TxManager) Do(ctx context.Context, fn func(ctx context.Context, tx *gorm.DB) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(ctx, tx)
	})
}
