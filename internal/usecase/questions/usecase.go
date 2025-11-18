package questions

import (
	"context"

	"github.com/PANDA-1703/API-questions-and-answers/internal/config"
	"github.com/PANDA-1703/API-questions-and-answers/internal/entity"
	"github.com/PANDA-1703/API-questions-and-answers/pkg/logger"
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
