package handlers

import (
	service "github.com/agnynureza/homework-rakamin-golang-sql/services"
)

type TokenHandler struct {
	movieService service.MovieServiceInterface
}

func NewTokenHandler(movieService service.MovieServiceInterface) *TokenHandler {
	return &TokenHandler{
		movieService: movieService,
	}
}

type TokenHandlerInterface interface {
}
