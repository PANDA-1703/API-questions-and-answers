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
	Delete(ctx context.Context, id int64) error
}

type QuestionUsecase struct {
	cfg           *config.ServiceConfig
	questionsRepo QuestionsRepo
	logger        logger.Logger
}

func NewQuestion(
	cfg *config.ServiceConfig,
	questionsRepo QuestionsRepo,
	logger logger.Logger,
) *QuestionUsecase {
	return &QuestionUsecase{
		cfg:           cfg,
		questionsRepo: questionsRepo,
		logger:        logger,
	}
}
