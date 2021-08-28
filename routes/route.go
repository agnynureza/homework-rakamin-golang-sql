package routes

import (
	"github.com/agnynureza/homework-rakamin-golang-sql/handlers"
	"github.com/agnynureza/homework-rakamin-golang-sql/middleware"
	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	movieHandler handlers.MovieHandlerInterface
	tokenHandler handlers.TokenHandlerInterface
}

func NewRoutes(movieHandler handlers.MovieHandlerInterface, tokenHandler handlers.TokenHandlerInterface) *Routes {
	return &Routes{
		movieHandler: movieHandler,
		tokenHandler: tokenHandler,
	}
}

func (r *Routes) InitializeRoutes(a *fiber.App) {
	moviesRoute := a.Group("/movie")

	// movies route
	moviesRoute.Post("/", middleware.JWTProtected(), r.movieHandler.PostNewMovies)
	moviesRoute.Get("/:slug", middleware.JWTProtected(), r.movieHandler.GetMovieBySlug)
	moviesRoute.Put("/:slug", middleware.JWTProtected(), r.movieHandler.PutMovie)
	moviesRoute.Delete("/:slug", middleware.JWTProtected(), r.movieHandler.DeleteMovieBySlug)

	// login
	a.Get("/login", r.tokenHandler.GetNewAccessToken)
}
