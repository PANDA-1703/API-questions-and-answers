package entity

import (
	"time"

	"github.com/PANDA-1703/API-questions-and-answers/api/gen/swagger/models"
	"github.com/PANDA-1703/API-questions-and-answers/pkg/utils"

	"github.com/go-openapi/strfmt"
)

type Answer struct {
	ID         int64     `gorm:"primaryKey;autoIncrement"`
	QuestionID int64     `gorm:"not null;index"`
	UserID     string    `gorm:"type:varchar(255);notnull"`
	Text       string    `gorm:"type:text;not null"`
	CreatedAt  time.Time `gorm:"not null;default:now()"`

	// связь с Question
	Question Question `gorm:"fireignKey:QuestionID;constraint:OnDelete:CASCADE"`
}

func (Answer) TableName() string {
	return "answers"
}

func FromHTTPAnswer(a *models.AnswerCreate, questionID int64) *Answer {
	return &Answer{
		QuestionID: questionID,
		UserID:     utils.FromPtr(a.UserID),
		Text:       utils.FromPtr(a.Text),
		CreatedAt:  time.Now(),
	}
}

func (a *Answer) ToHTTPAnswer() *models.Answer {
	return &models.Answer{
		ID:         utils.ToPtr(a.ID),
		QuestionID: utils.ToPtr(a.QuestionID),
		UserID:     utils.ToPtr(a.UserID),
		Text:       utils.ToPtr(a.Text),
		CreatedAt:  (*strfmt.DateTime)(utils.ToPtr(a.CreatedAt)),
	}
}
