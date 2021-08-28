package repository

import (
	"errors"

	"github.com/agnynureza/homework-rakamin-golang-sql/models"
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
	CreateMovie(movie *models.Movies) (int, error)
	GetOneMovie(title string) (models.Movies, error)
}

func (r *MoviesRepository) CreateMovie(movie *models.Movies) (int, error) {
	err := r.db.Create(movie).Error
	if err != nil {
		return movie.ID, err
	}
	return movie.ID, nil
}

func (r *MoviesRepository) GetOneMovie(title string) (models.Movies, error) {
	var movie models.Movies
	query := `SELECT id, title, slug, description, duration, image FROM movies WHERE title = ?`

	err := r.db.Raw(query, title).Scan(&movie).Error
	if err != nil {
		return movie, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return movie, nil
	}

	return movie, nil
}
