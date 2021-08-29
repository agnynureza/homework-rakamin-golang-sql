package handlers

import (
	"errors"

	"github.com/agnynureza/homework-rakamin-golang-sql/common/utils"
	"github.com/agnynureza/homework-rakamin-golang-sql/models"
	service "github.com/agnynureza/homework-rakamin-golang-sql/services"
	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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
	GetMovieBySlug(c *fiber.Ctx) error
	PutMovie(c *fiber.Ctx) error
	DeleteMovieBySlug(c *fiber.Ctx) error
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

func (m *MoviesHandler) GetMovieBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	response, err := m.movieService.GetMovie(slug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   "data not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  false,
		"msg":    "success retrieve data",
		"result": response,
	})
}

func (m *MoviesHandler) PutMovie(c *fiber.Ctx) error {
	movie := &models.Movies{}
	slug := c.Params("slug")
	var mysqlErr *mysql.MySQLError

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

	response, err := m.movieService.UpdateMovie(movie, slug)
	if err != nil {
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  false,
		"msg":    "success update data",
		"result": response,
	})
}

func (m *MoviesHandler) DeleteMovieBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	err := m.movieService.DeleteMovie(slug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   "data not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  false,
		"result": "success",
	})
}
