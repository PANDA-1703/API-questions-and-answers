package answers

import (
	"API-quest-answ/internal/config"
	"API-quest-answ/internal/entity"
	"API-quest-answ/internal/usecase/questions"
	"API-quest-answ/pkg/logger"
	"context"
)

/*
Сгенерить моки:

cd internal/usecase/answers
mockgen -source=usecase.go -destination=./mocks/answers_repo_mock.go -package=mocks
*/

type AnswersRepo interface {
	Create(ctx context.Context, answer *entity.Answer) (int64, error)
	GetByID(ctx context.Context, id int64) (*entity.Answer, error)
	GetAllByQuestionID(ctx context.Context, questionID int64) ([]*entity.Answer, error)
	Delete(ctx context.Context, id int64, userID string) error
}

type AnswersUsecase struct {
	cfg           *config.ServiceConfig
	answersRepo   AnswersRepo
	questionsRepo questions.QuestionsRepo
	logger        logger.Logger
}

func NewAnswer(
	cfg *config.ServiceConfig,
	answersRepo AnswersRepo,
	questionsRepo questions.QuestionsRepo,
	logger logger.Logger,
) *AnswersUsecase {
	return &AnswersUsecase{
		cfg:           cfg,
		answersRepo:   answersRepo,
		questionsRepo: questionsRepo,
		logger:        logger,
	}
}
