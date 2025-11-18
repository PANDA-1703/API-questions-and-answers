package answers

import (
	"context"
	"github.com/PANDA-1703/API-questions-and-answers/internal/entity"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, answer *entity.Answer) (int64, error)
	GetByID(ctx context.Context, id int64) (*entity.Answer, error)
	GetAllByQuestionID(ctx context.Context, questionID int64) ([]*entity.Answer, error)
	Delete(ctx context.Context, id int64, userID string) error
}

type repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repo{
		db: db,
	}
}
