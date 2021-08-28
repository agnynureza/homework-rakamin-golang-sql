package routes

import (
	"github.com/agnynureza/homework-rakamin-golang-sql/handlers"
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
	// movies route
	// route.Post("/book", middleware.JWTProtected(), b.BookController.CreateBook)
	// route.Post("/book/:id/upload-image", middleware.JWTProtected(), b.BookController.UploadBookImage)
	// route.Get("/books", middleware.JWTProtected(), b.BookController.GetBooks)
	// route.Get("/book/:id", middleware.JWTProtected(), b.BookController.GetBookByID)
	// route.Put("/book", middleware.JWTProtected(), b.BookController.UpdateBook)
	// route.Delete("/book", middleware.JWTProtected(), b.BookController.DeleteBook)

	// login
	//a.Get("/login")
	//route.Get("/login", b.BookController.GetNewAccessToken)
}
