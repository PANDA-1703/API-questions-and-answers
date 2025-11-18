package questions

import (
	"context"
	"github.com/PANDA-1703/API-questions-and-answers/internal/entity"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, question *entity.Question) (int64, error)
	GetAll(ctx context.Context) ([]*entity.Question, error)
	GetByID(ctx context.Context, id int64) (*entity.Question, error)
	Delete(ctx context.Context, id int64) error
}

type repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repo{
		db: db,
	}
}
