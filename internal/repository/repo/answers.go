package repo

import (
	"API-quest-answ/internal/entity"
	"context"
	"fmt"
)

func (r Repo) CreateAnswer(ctx context.Context, answer *entity.Answer) (int64, error) {
	result := r.db.WithContext(ctx).Create(answer)
	if result.Error != nil {
		return 0, fmt.Errorf("AnswerRepo.CreateAnswer: %w", result.Error)
	}
	return answer.ID, nil
}

func (r Repo) GetAnswerByID(ctx context.Context, id int64) (*entity.Answer, error) {
	var answer *entity.Answer
	result := r.db.WithContext(ctx).First(&answer, id)
	if result.Error != nil {
		return nil, fmt.Errorf("AnswersRepo.GetByID: %w", result.Error)
	}
	return answer, nil
}

func (r Repo) GetAllByQuestionID(ctx context.Context, questionID int64) ([]*entity.Answer, error) {
	var answers []*entity.Answer
	result := r.db.WithContext(ctx).Where("question_id = ?", questionID).Find(&answers)
	if result.Error != nil {
		return nil, fmt.Errorf("AnswersRepo.GetAllByQuestionID: %w", result.Error)
	}
	return answers, nil
}

func (r Repo) DeleteAnswer(ctx context.Context, id int64, userID string) error {
	result := r.db.WithContext(ctx).Where("id = ? AND user_id = ?", id, userID).Delete(&entity.Answer{})
	if result.Error != nil {
		return fmt.Errorf("AnswersRepo.DeleteAnswer: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("AnswersRepo.DeleteAnswer: answer not found or permission denied")
	}
	return nil
}
