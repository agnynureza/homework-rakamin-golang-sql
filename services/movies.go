package services

import (
	"github.com/agnynureza/homework-rakamin-golang-sql/models"
	"github.com/agnynureza/homework-rakamin-golang-sql/repository"
)

type MoviesService struct {
	movieRepo repository.MovieRepoInterface
}

func NewMoviesService(movieRepo repository.MovieRepoInterface) *MoviesService {
	return &MoviesService{
		movieRepo: movieRepo,
	}
}

type MovieServiceInterface interface {
	CreateNewMovie(movie *models.Movies) (*models.Movies, error)
	GetMovie(title string) (models.Movies, error)
	UpdateMovie(movie *models.Movies, slug string) (models.Movies, error)
	DeleteMovie(slug string) error
}

func (m *MoviesService) CreateNewMovie(movie *models.Movies) (*models.Movies, error) {
	id, err := m.movieRepo.CreateMovie(movie)
	if err != nil {
		return nil, err
	}
	movie.ID = id

	return movie, nil
}

func (m *MoviesService) GetMovie(slug string) (models.Movies, error) {
	movie, err := m.movieRepo.GetOneMovie(slug)
	if err != nil {
		return movie, err
	}

	return movie, nil
}

func (m *MoviesService) UpdateMovie(payload *models.Movies, slug string) (movie models.Movies, err error) {
	// update movie data
	err = m.movieRepo.UpdateMovie(payload, slug)
	if err != nil {
		return movie, err
	}

	// select movie
	movie, err = m.movieRepo.GetOneMovie(payload.Slug)
	if err != nil {
		return movie, err
	}

	return movie, nil
}

func (m *MoviesService) DeleteMovie(slug string) error {
	err := m.movieRepo.DeleteMovie(slug)
	if err != nil {
		return err
	}

	return nil
}
