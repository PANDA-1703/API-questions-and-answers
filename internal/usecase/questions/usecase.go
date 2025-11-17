package questions

import (
	"API-quest-answ/internal/config"
	"API-quest-answ/internal/entity"
	"API-quest-answ/pkg/logger"
	"context"
)

/*
Сгенерить моки:

cd internal/usecase/answers
mockgen -source=usecase.go -destination=./mocks/questions_repo_mock.go -package=mocks
*/

type QuestionsRepo interface {
	Create(ctx context.Context, question *entity.Question) (int64, error)
	GetAll(ctx context.Context) ([]*entity.Question, error)
	GetByID(ctx context.Context, id int64) (*entity.Question, error)
	Delete(ctx context.Context, id int64, userID string) error
}

type Usecase struct {
	cfg           *config.ServiceConfig
	questionsRepo QuestionsRepo
	logger        logger.Logger
}

func New(
	cfg *config.ServiceConfig,
	questionsRepo QuestionsRepo,
	logger logger.Logger,
) *Usecase {
	return &Usecase{
		cfg:           cfg,
		questionsRepo: questionsRepo,
		logger:        logger,
	}
}
