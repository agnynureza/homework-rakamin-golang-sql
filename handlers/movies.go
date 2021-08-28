package handlers

import (
	"github.com/agnynureza/homework-rakamin-golang-sql/common/utils"
	"github.com/agnynureza/homework-rakamin-golang-sql/models"
	service "github.com/agnynureza/homework-rakamin-golang-sql/services"
	"github.com/gofiber/fiber/v2"
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
	PostNewMovies(c *fiber.Ctx) error
	GetMovieByTitle(c *fiber.Ctx) error
}

func (m *MoviesHandler) PostNewMovies(c *fiber.Ctx) error {
	movie := &models.Movies{}

	if err := c.BodyParser(movie); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	validate := utils.NewValidator()
	if err := validate.Struct(movie); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	response, err := m.movieService.CreateNewMovie(movie)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":  false,
		"msg":    "success create data",
		"result": response,
	})
}

func (m *MoviesHandler) GetMovieByTitle(c *fiber.Ctx) error {
	title := c.Params("title")
	response, err := m.movieService.GetMovie(title)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if response.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "data not found",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":  false,
		"msg":    "success retrieve data",
		"result": response,
	})
}
