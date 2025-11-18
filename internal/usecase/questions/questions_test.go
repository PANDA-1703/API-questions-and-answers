package questions

import (
	"context"
	"errors"
	"testing"

	"github.com/PANDA-1703/API-questions-and-answers/internal/entity"
	"github.com/PANDA-1703/API-questions-and-answers/internal/usecase/questions/mocks"
	"github.com/PANDA-1703/API-questions-and-answers/pkg/logger"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestQuestionUsecase_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockQuestionsRepo(ctrl)
	uc := NewQuestion(nil, mockRepo, logger.New())

	q := &entity.Question{Text: "test question"}

	mockRepo.EXPECT().
		Create(gomock.Any(), q).
		Return(int64(42), nil)

	id, err := uc.Create(context.Background(), q)

	assert.NoError(t, err)
	assert.Equal(t, int64(42), id)
}

func TestQuestionUsecase_GetByID_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockQuestionsRepo(ctrl)
	uc := NewQuestion(nil, mockRepo, logger.New())

	mockRepo.EXPECT().
		GetByID(gomock.Any(), int64(1)).
		Return(nil, errors.New("db error"))

	res, err := uc.GetByID(context.Background(), 1)

	assert.Error(t, err)
	assert.Nil(t, res)
}
