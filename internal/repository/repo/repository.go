package repo

import (
	"API-quest-answ/internal/repository"

	"gorm.io/gorm"
)

type Repo struct {
	db     *gorm.DB
	getter repository.ContextGetter
}

func New(
	db *gorm.DB,
	getter repository.ContextGetter,
) *Repo {
	return &Repo{
		db:     db,
		getter: getter,
	}
}
