package repository

import (
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
	GetOneMovie(slug string) (models.Movies, error)
	UpdateMovie(payload *models.Movies, slug string) error
	DeleteMovie(slug string) error
}

func (r *MoviesRepository) CreateMovie(movie *models.Movies) (int, error) {
	err := r.db.Create(movie).Error
	if err != nil {
		return movie.ID, err
	}
	return movie.ID, nil
}

func (r *MoviesRepository) GetOneMovie(slug string) (models.Movies, error) {
	var movie models.Movies
	query := `SELECT id, title, slug, description, duration, image FROM movies WHERE slug = ?`

	err := r.db.Raw(query, slug).Scan(&movie).Error
	if err != nil {
		return movie, err
	}

	if movie.ID == 0 {
		return movie, gorm.ErrRecordNotFound
	}

	return movie, nil
}

func (r *MoviesRepository) UpdateMovie(payload *models.Movies, slug string) error {
	query := `UPDATE movies SET title = ?, description = ?, duration = ?, image = ?, slug = ? WHERE slug = ?`
	result := r.db.Exec(query, payload.Title, payload.Description, payload.Duration, payload.Image, payload.Slug, slug)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *MoviesRepository) DeleteMovie(slug string) error {
	var movie models.Movies
	result := r.db.Where("slug = ?", slug).Delete(&movie)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
