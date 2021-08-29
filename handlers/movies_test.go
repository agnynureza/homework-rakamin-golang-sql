package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/agnynureza/homework-rakamin-golang-sql/common/utils"
	"github.com/agnynureza/homework-rakamin-golang-sql/middleware"
	mocks "github.com/agnynureza/homework-rakamin-golang-sql/mocks/services"
	"github.com/agnynureza/homework-rakamin-golang-sql/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var (
	handlerMovie      *MoviesHandler
	servicesMock      *mocks.MockMovieServiceInterface
	JWToken           string
	authHeader        = "Authorization"
	contentType       = "Content-Type"
	app               *fiber.App
	mockCtrl          *gomock.Controller
	errInternal       = errors.New("unexpected system error")
	errRecordNotFound = errors.New("record not found")
	defaultMovies     models.Movies
)

type ResponseJson struct {
	Error   bool          `json:"error"`
	Message string        `json:"msg"`
	Result  models.Movies `json:"result"`
}

func TestMain(m *testing.M) {
	app = fiber.New(fiber.Config{
		Immutable:     true,
		CaseSensitive: true,
		AppName:       "Movies",
	})

	os.Setenv("JWT_SECRET_KEY", "supersecret")
	os.Setenv("JWT_SECRET_KEY_EXPIRE_DAYS_COUNT", "7")

	JWToken, _ = utils.GenerateNewAccessToken()

	exitVal := m.Run()

	defer mockCtrl.Finish()

	os.Exit(exitVal)
}

func InitGomock(t *testing.T) {
	mockCtrl = gomock.NewController(t)
	servicesMock = mocks.NewMockMovieServiceInterface(mockCtrl)
	handlerMovie = NewMoviesHandler(servicesMock)
}

func TestMoviesHandler_PostNewMovies(t *testing.T) {
	InitGomock(t)
	app.Post("/movie", middleware.JWTProtected(), handlerMovie.PostNewMovies)
	expectedResponse := models.Movies{
		ID:          1,
		Title:       "Titanic",
		Slug:        "titanic",
		Description: "lorem ipsum",
		Duration:    60,
		Image:       "image titanic URL",
	}

	t.Run("should return success", func(t *testing.T) {
		servicesMock.EXPECT().CreateNewMovie(gomock.Any()).Return(&expectedResponse, nil)
		url := "http://example.com/movie"

		payload := map[string]interface{}{
			"title":       "golang",
			"slug":        "golang",
			"description": "lorem ipsum",
			"duration":    60,
			"image":       "image titanic URL",
		}
		jsonValue, _ := json.Marshal(payload)
		req := httptest.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
		req.Header.Set(authHeader, "Bearer "+JWToken)
		req.Header.Set(contentType, "application/json")

		resp, err := app.Test(req)
		require.Equal(t, err, nil)
		require.Equal(t, resp.StatusCode, http.StatusCreated)

		body, _ := ioutil.ReadAll(resp.Body)
		var movieActual ResponseJson
		err = json.Unmarshal(body, &movieActual)
		require.Equal(t, err, nil)
		require.Equal(t, movieActual.Result, expectedResponse)
	})

	t.Run("should return error validation", func(t *testing.T) {
		url := "http://example.com/movie"

		payload := map[string]interface{}{
			"slug":        "golang",
			"description": "lorem ipsum",
			"duration":    60,
			"image":       "image titanic URL",
		}

		jsonValue, _ := json.Marshal(payload)
		req := httptest.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
		req.Header.Set(authHeader, "Bearer "+JWToken)
		req.Header.Set(contentType, "application/json")

		resp, err := app.Test(req)
		require.Equal(t, err, nil)
		require.Equal(t, resp.StatusCode, http.StatusBadRequest)
	})

	t.Run("should return error from service", func(t *testing.T) {
		servicesMock.EXPECT().CreateNewMovie(gomock.Any()).Return(nil, errInternal)

		url := "http://example.com/movie"

		payload := map[string]interface{}{
			"title":       "golang",
			"slug":        "golang",
			"description": "lorem ipsum",
			"duration":    60,
			"image":       "image titanic URL",
		}

		jsonValue, _ := json.Marshal(payload)
		req := httptest.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
		req.Header.Set(authHeader, "Bearer "+JWToken)
		req.Header.Set(contentType, "application/json")

		resp, err := app.Test(req)
		require.Equal(t, err, nil)
		require.Equal(t, resp.StatusCode, http.StatusInternalServerError)
	})
}

func TestMoviesHandler_GetMovieBySlug(t *testing.T) {
	InitGomock(t)
	app.Get("/movie/:slug", middleware.JWTProtected(), handlerMovie.GetMovieBySlug)
	// expectedResponse := models.Movies{
	// 	ID:          1,
	// 	Title:       "Titanic",
	// 	Slug:        "titanic",
	// 	Description: "lorem ipsum",
	// 	Duration:    60,
	// 	Image:       "image titanic URL",
	// }

	// t.Run("should return success", func(t *testing.T) {
	// 	servicesMock.EXPECT().GetMovie(gomock.Any()).Return(expectedResponse, nil)

	// 	url := "http://example.com/movie/titanic"

	// 	req := httptest.NewRequest("GET", url, nil)
	// 	req.Header.Set(authHeader, "Bearer "+JWToken)
	// 	req.Header.Set(contentType, "application/json")

	// 	resp, err := app.Test(req)
	// 	require.Equal(t, err, nil)
	// 	require.Equal(t, resp.StatusCode, http.StatusOK)

	// 	body, _ := ioutil.ReadAll(resp.Body)
	// 	var movieActual ResponseJson
	// 	err = json.Unmarshal(body, &movieActual)
	// 	require.Equal(t, err, nil)
	// 	require.Equal(t, movieActual.Result, expectedResponse)
	// })

	t.Run("should return data not found", func(t *testing.T) {
		servicesMock.EXPECT().GetMovie(gomock.Any()).Return(defaultMovies, errRecordNotFound)

		url := "http://example.com/movie/titanic"

		req := httptest.NewRequest("GET", url, nil)
		req.Header.Set(authHeader, "Bearer "+JWToken)
		req.Header.Set(contentType, "application/json")

		resp, err := app.Test(req)
		require.Equal(t, err, nil)
		require.Equal(t, resp.StatusCode, http.StatusNotFound)

		body, _ := ioutil.ReadAll(resp.Body)
		var movieActual ResponseJson
		err = json.Unmarshal(body, &movieActual)
		require.Equal(t, err, nil)
		require.Equal(t, movieActual.Message, "data not found")
	})
}
