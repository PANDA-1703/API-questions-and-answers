package repo

import (
	"API-quest-answ/internal/entity"
	"context"
	"fmt"
)

func (r Repo) Create(ctx context.Context, question *entity.Question) (int64, error) {
	result := r.db.WithContext(ctx).Create(&question)
	if result.Error != nil {
		return 0, fmt.Errorf("QuestionsRepo.Create: %w", result.Error)
	}
	return question.ID, nil
}

func (r Repo) GetAll(ctx context.Context) ([]*entity.Question, error) {
	var questions []*entity.Question
	result := r.db.WithContext(ctx).Find(&questions)
	if result.Error != nil {
		return nil, fmt.Errorf("QuestionsRepo.GetAll: %w", result.Error)
	}
	return questions, nil
}

func (r Repo) GetByID(ctx context.Context, id int64) (*entity.Question, error) {
	var question *entity.Question
	result := r.db.WithContext(ctx).First(&question, id)
	if result.Error != nil {
		return nil, fmt.Errorf("QuestionsRepo.GetByID: %w", result.Error)
	}
	return question, nil
}

func (r Repo) Delete(ctx context.Context, id int64) error {
	result := r.db.WithContext(ctx).Delete(&entity.Question{}, id)
	if result.Error != nil {
		return fmt.Errorf("QuestionsRepo.Delete: %w", result.Error)
	}
	return nil
}
