package entity

import (
	"API-quest-answ/api/gen/swagger/models"
	"API-quest-answ/pkg/utils"
	"time"

	"github.com/go-openapi/strfmt"
)

type Question struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	Text      string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
}

func (Question) TableName() string {
	return "repo"
}

func FromHTTPQuestion(q *models.QuestionCreate) *Question {
	return &Question{
		Text:      utils.FromPtr(q.Text),
		CreatedAt: time.Now(),
	}
}

func (q *Question) ToHTTPQuestion() *models.Question {
	return &models.Question{
		ID:        utils.ToPtr(q.ID),
		Text:      utils.ToPtr(q.Text),
		CreatedAt: (*strfmt.DateTime)(utils.ToPtr(q.CreatedAt)),
	}
}
