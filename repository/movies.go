package repository

import (
	"gorm.io/gorm"
)

type MoviesRepository struct {
	db *gorm.DB
}

func NewMoviesRepository(db *gorm.DB) *MoviesRepository {
	return &MoviesRepository{
		db: db,
	}
}

type MovieRepoInterface interface {
}
