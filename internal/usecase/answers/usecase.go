package answers

import (
	"API-quest-answ/internal/config"
	"API-quest-answ/internal/entity"
	"API-quest-answ/pkg/logger"
	"context"
)

/*
Сгенерить моки:

cd internal/usecase/answers
mockgen -source=usecase.go -destination=./mocks/answers_repo_mock.go -package=mocks
*/

type AnswersRepo interface {
	CreateAnswer(ctx context.Context, answer *entity.Answer) (int64, error)
	GetAnswerByID(ctx context.Context, id int64) (*entity.Answer, error)
	GetAllByQuestionID(ctx context.Context, questionID int64) ([]*entity.Answer, error)
	DeleteAnswer(ctx context.Context, id int64, userID string) error
}

type Usecase struct {
	cfg         *config.ServiceConfig
	answersRepo AnswersRepo
	logger      logger.Logger
}

func New(
	cfg *config.ServiceConfig,
	answersRepo AnswersRepo,
	logger logger.Logger,
) *Usecase {
	return &Usecase{
		cfg:         cfg,
		answersRepo: answersRepo,
		logger:      logger,
	}
}
