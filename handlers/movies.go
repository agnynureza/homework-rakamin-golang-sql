package handlers

import (
	service "github.com/agnynureza/homework-rakamin-golang-sql/services"
)

type MoviesHandler struct {
	movieService service.MovieServiceInterface
}

func NewMoviesHandler(movieService service.MovieServiceInterface) *MoviesHandler {
	return &MoviesHandler{
		movieService: movieService,
	}
}

type MovieHandlerInterface interface {
}
