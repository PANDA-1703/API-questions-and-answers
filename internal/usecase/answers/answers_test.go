package answers

import (
	"context"
	"errors"
	"testing"

	"github.com/PANDA-1703/API-questions-and-answers/internal/entity"
	amocks "github.com/PANDA-1703/API-questions-and-answers/internal/usecase/answers/mocks"
	qmocks "github.com/PANDA-1703/API-questions-and-answers/internal/usecase/questions/mocks"
	"github.com/PANDA-1703/API-questions-and-answers/pkg/logger"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAnswersUsecase_Create_Ok(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	aRepo := amocks.NewMockAnswersRepo(ctrl)
	qRepo := qmocks.NewMockQuestionsRepo(ctrl)

	uc := NewAnswer(nil, aRepo, qRepo, logger.New())

	answer := &entity.Answer{QuestionID: 10, Text: "answer text"}

	qRepo.EXPECT().
		GetByID(gomock.Any(), int64(10)).
		Return(&entity.Question{}, nil)

	aRepo.EXPECT().
		Create(gomock.Any(), answer).
		Return(int64(55), nil)

	id, err := uc.Create(context.Background(), answer)

	assert.NoError(t, err)
	assert.Equal(t, int64(55), id)
}

func TestAnswersUsecase_Create_QuestionDoesNotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	aRepo := amocks.NewMockAnswersRepo(ctrl)
	qRepo := qmocks.NewMockQuestionsRepo(ctrl)

	uc := NewAnswer(nil, aRepo, qRepo, logger.New())

	answer := &entity.Answer{QuestionID: 10}

	qRepo.EXPECT().
		GetByID(gomock.Any(), int64(10)).
		Return(nil, errors.New("not found"))

	id, err := uc.Create(context.Background(), answer)

	assert.Error(t, err)
	assert.Equal(t, int64(0), id)
}
