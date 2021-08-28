package cli

import (
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/agnynureza/homework-rakamin-golang-sql/app"
	"github.com/agnynureza/homework-rakamin-golang-sql/config"
	"github.com/agnynureza/homework-rakamin-golang-sql/config/database"
	"github.com/agnynureza/homework-rakamin-golang-sql/handlers"
	"github.com/agnynureza/homework-rakamin-golang-sql/middleware"
	"github.com/agnynureza/homework-rakamin-golang-sql/repository"
	route "github.com/agnynureza/homework-rakamin-golang-sql/routes"
	"github.com/agnynureza/homework-rakamin-golang-sql/services"
	"github.com/gofiber/fiber/v2"

	log "github.com/sirupsen/logrus"
)

type Cli struct {
	Args []string
}

func NewCli(args []string) *Cli {
	return &Cli{
		Args: args,
	}
}

func (c *Cli) Run(application *app.Application) {
	fiberConfig := config.FiberConfig()
	app := fiber.New(fiberConfig)

	//set up logger
	log.SetLevel(log.InfoLevel)
	log.StandardLogger()
	log.SetOutput(os.Stdout)
	if strings.ToLower(application.Config.LogLevel) == log.DebugLevel.String() {
		log.SetLevel(log.DebugLevel)
	}
	log.SetReportCaller(true)

	//middleware
	middleware.FiberMiddleware(app)

	//set up connection
	connDB := database.InitDb()

	// movies services
	moviesRepository := repository.NewMoviesRepository(connDB)
	moviesService := services.NewMoviesService(moviesRepository)
	moviesHandlers := handlers.NewMoviesHandler(moviesService)

	// token services
	var tokensHandlers = &handlers.TokenHandler{}

	// register handler to Routes
	routes := route.NewRoutes(moviesHandlers, tokensHandlers)
	routes.InitializeRoutes(app)

	//not found routes
	route.NotFoundRoute(app)

	log.Println(fmt.Sprintf("starting application { %v } on port :%s", application.Config.AppName, application.Config.AppPort))

	StartServerWithGracefulShutdown(app, application.Config.AppPort)

}

func StartServerWithGracefulShutdown(a *fiber.App, port string) {
	appPort := fmt.Sprintf(`:%s`, port)
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := a.Shutdown(); err != nil {
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	// Run server.
	if err := a.Listen(appPort); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}


