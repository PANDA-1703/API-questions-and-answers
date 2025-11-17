package answers

import (
	"API-quest-answ/internal/entity"
	"context"
	"fmt"
)

//type AnswersRepo interface {
//	CreateAnswer(ctx context.Context, answer entity.Answer) (int64, error)
//	GetAnswerByID(ctx context.Context, id int64) (entity.Answer, error)
//	GetAllByQuestionID(ctx context.Context, questionID int64) ([]entity.Answer, error)
//	DeleteAnswer(ctx context.Context, id int64, userID string) error
//}

func (u Usecase) CreateAnswers(ctx context.Context, answer *entity.Answer) (int64, error) {
	id, err := u.answersRepo.CreateAnswer(ctx, answer)
	if err != nil {
		return 0, fmt.Errorf("uc answers: create failed: %w", err)
	}
	return id, nil
}

func (u Usecase) GetAnswerByID(ctx context.Context, id int64) (*entity.Answer, error) {
	answer, err := u.answersRepo.GetAnswerByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("uc answers: get by id failed: %w", err)
	}
	return answer, nil
}

func (u Usecase) GetAllByQuestionID(ctx context.Context, questionID int64) ([]*entity.Answer, error) {
	answers, err := u.answersRepo.GetAllByQuestionID(ctx, questionID)
	if err != nil {
		return nil, fmt.Errorf("uc answers: get all by question ID failed: %w", err)
	}
	return answers, nil
}

func (u Usecase) DeleteAnswer(ctx context.Context, id int64, userID string) error {
	err := u.answersRepo.DeleteAnswer(ctx, id, userID)
	if err != nil {
		return fmt.Errorf("uc answers: delete failed: %w", err)
	}
	return nil
}
