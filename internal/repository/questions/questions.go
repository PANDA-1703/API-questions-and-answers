package questions

import (
	"context"
	"fmt"
	"github.com/PANDA-1703/API-questions-and-answers/internal/entity"
)

func (r *repo) Create(ctx context.Context, question *entity.Question) (int64, error) {
	result := r.db.WithContext(ctx).Create(&question)
	if result.Error != nil {
		return 0, fmt.Errorf("QuestionsRepo.Create: %w", result.Error)
	}
	return question.ID, nil
}

func (r *repo) GetAll(ctx context.Context) ([]*entity.Question, error) {
	var questions []*entity.Question
	result := r.db.WithContext(ctx).Find(&questions)
	if result.Error != nil {
		return nil, fmt.Errorf("QuestionsRepo.GetAll: %w", result.Error)
	}
	return questions, nil
}

func (r *repo) GetByID(ctx context.Context, id int64) (*entity.Question, error) {
	var question *entity.Question
	result := r.db.WithContext(ctx).First(&question, id)
	if result.Error != nil {
		return nil, fmt.Errorf("QuestionsRepo.GetByID: %w", result.Error)
	}
	return question, nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	var question *entity.Question
	result := r.db.WithContext(ctx).Delete(&question, id)
	if result.Error != nil {
		return fmt.Errorf("QuestionsRepo.Delete: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("QuestionsRepo.Delete: questions not found or permission denied")
	}
	return nil
}
