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
}

func (m *MoviesService) CreateNewMovie(movie *models.Movies) (*models.Movies, error) {
	id, err := m.movieRepo.CreateMovie(movie)
	if err != nil {
		return nil, err
	}
	movie.ID = id

	return movie, nil
}

func (m *MoviesService) GetMovie(title string) (models.Movies, error) {
	movie, err := m.movieRepo.GetOneMovie(title)
	if err != nil {
		return movie, err
	}

	return movie, nil
}
