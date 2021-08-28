package services

import (
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
}
