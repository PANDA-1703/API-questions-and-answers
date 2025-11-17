package questions

import (
	"API-quest-answ/internal/entity"
	"context"
	"fmt"
)

func (u QuestionUsecase) Create(ctx context.Context, question *entity.Question) (int64, error) {
	id, err := u.questionsRepo.Create(ctx, question)
	if err != nil {
		return 0, fmt.Errorf("uc questions: create failed: %w", err)
	}
	return id, nil
}

func (u QuestionUsecase) GetAll(ctx context.Context) ([]*entity.Question, error) {
	questions, err := u.questionsRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("uc questions: get all failed: %w", err)
	}
	return questions, nil
}

func (u QuestionUsecase) GetByID(ctx context.Context, id int64) (*entity.Question, error) {
	question, err := u.questionsRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("uc questions: get by id failed: %w", err)
	}
	return question, nil
}

func (u QuestionUsecase) Delete(ctx context.Context, id int64) error {
	err := u.questionsRepo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("uc questions: delete failed %w", err)
	}
	return nil
}
