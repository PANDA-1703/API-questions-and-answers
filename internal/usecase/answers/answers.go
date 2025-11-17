package answers

import (
	"API-quest-answ/internal/entity"
	"context"
	"fmt"
)

func (u AnswersUsecase) Create(ctx context.Context, answer *entity.Answer) (int64, error) {
	_, err := u.questionsRepo.GetByID(ctx, answer.QuestionID)
	if err != nil {
		return 0, fmt.Errorf("question does not exist: %w", err)
	}

	id, err := u.answersRepo.Create(ctx, answer)
	if err != nil {
		return 0, fmt.Errorf("uc answers: create failed: %w", err)
	}
	return id, nil
}

func (u AnswersUsecase) GetByID(ctx context.Context, id int64) (*entity.Answer, error) {
	answer, err := u.answersRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("uc answers: get by id failed: %w", err)
	}
	return answer, nil
}

func (u AnswersUsecase) GetAllByQuestionID(ctx context.Context, questionID int64) ([]*entity.Answer, error) {
	answers, err := u.answersRepo.GetAllByQuestionID(ctx, questionID)
	if err != nil {
		return nil, fmt.Errorf("uc answers: get all by question ID failed: %w", err)
	}
	return answers, nil
}

func (u AnswersUsecase) Delete(ctx context.Context, id int64, userID string) error {
	err := u.answersRepo.Delete(ctx, id, userID)
	if err != nil {
		return fmt.Errorf("uc answers: delete failed: %w", err)
	}
	return nil
}
