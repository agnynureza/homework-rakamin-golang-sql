package services

import (
	"errors"
	"testing"

	mocks "github.com/agnynureza/homework-rakamin-golang-sql/mocks/repository"
	"github.com/agnynureza/homework-rakamin-golang-sql/models"
	"github.com/golang/mock/gomock"
	"github.com/tj/assert"
)

var (
	mockCtrl          *gomock.Controller
	repoMock          *mocks.MockMovieRepoInterface
	service           *MoviesService
	errInternal       = errors.New("unexpected system error")
	defaulInt         = 0
	defaultMovieModel models.Movies
)

func InitMockRepo(t *testing.T) {
	mockCtrl = gomock.NewController(t)
	repoMock = mocks.NewMockMovieRepoInterface(mockCtrl)
	service = NewMoviesService(repoMock)
}

func TestMoviesService_CreateNewMovie(t *testing.T) {
	InitMockRepo(t)
	defer mockCtrl.Finish()

	expectedID := 1
	payload := &models.Movies{
		Title:       "Titanic",
		Slug:        "titanic",
		Description: "lorem ipsum",
		Duration:    60,
		Image:       "image titanic URL",
	}
	expectedRespon := payload
	expectedRespon.ID = expectedID

	t.Run("should return success", func(t *testing.T) {
		repoMock.EXPECT().CreateMovie(gomock.Any()).Return(expectedID, nil)

		actualRespon, err := service.CreateNewMovie(payload)
		assert.Nil(t, err)
		assert.Equal(t, actualRespon, expectedRespon)
	})

	t.Run("should return error", func(t *testing.T) {
		repoMock.EXPECT().CreateMovie(gomock.Any()).Return(defaulInt, errInternal)

		respon, actualErr := service.CreateNewMovie(payload)
		assert.Nil(t, respon)
		assert.Equal(t, actualErr, errInternal)
	})
}

func TestMoviesService_GetMovie(t *testing.T) {
	InitMockRepo(t)
	defer mockCtrl.Finish()

	slug := "titanic"
	expectedRespon := models.Movies{
		ID:          1,
		Title:       "Titanic",
		Slug:        "titanic",
		Description: "lorem ipsum",
		Duration:    60,
		Image:       "image titanic URL",
	}

	t.Run("should return success", func(t *testing.T) {
		repoMock.EXPECT().GetOneMovie(gomock.Any()).Return(expectedRespon, nil)

		actualRespon, err := service.GetMovie(slug)
		assert.Nil(t, err)
		assert.Equal(t, actualRespon, expectedRespon)
	})

	t.Run("should return error", func(t *testing.T) {
		repoMock.EXPECT().GetOneMovie(gomock.Any()).Return(defaultMovieModel, errInternal)

		response, actualErr := service.GetMovie(slug)
		assert.Equal(t, response, defaultMovieModel)
		assert.Equal(t, actualErr, errInternal)
	})
}

func TestMoviesService_UpdateMovie(t *testing.T) {
	InitMockRepo(t)
	defer mockCtrl.Finish()

	slug := "titanic"

	payload := models.Movies{
		Title:       "Titanic",
		Slug:        "titanic",
		Description: "lorem ipsum",
		Duration:    60,
		Image:       "image titanic URL",
	}
	expectedRespon := payload
	expectedRespon.ID = 1

	t.Run("should return success", func(t *testing.T) {
		repoMock.EXPECT().GetOneMovie(gomock.Any()).Return(expectedRespon, nil)
		repoMock.EXPECT().UpdateMovie(gomock.Any(), gomock.Any()).Return(nil)

		actualRespon, err := service.UpdateMovie(&payload, slug)
		assert.Nil(t, err)
		assert.Equal(t, actualRespon, expectedRespon)
	})

	t.Run("should return error get movies", func(t *testing.T) {
		repoMock.EXPECT().GetOneMovie(gomock.Any()).Return(defaultMovieModel, errInternal)
		repoMock.EXPECT().UpdateMovie(gomock.Any(), gomock.Any()).Return(nil)

		response, actualErr := service.UpdateMovie(&payload, slug)
		assert.Equal(t, response, defaultMovieModel)
		assert.Equal(t, actualErr, errInternal)
	})

}

func TestMoviesService_DeleteMovie(t *testing.T) {
	InitMockRepo(t)
	defer mockCtrl.Finish()

	slug := "titanic"

	t.Run("should return success", func(t *testing.T) {
		repoMock.EXPECT().DeleteMovie(gomock.Any()).Return(nil)

		err := service.DeleteMovie(slug)
		assert.Nil(t, err)
	})

	t.Run("should return error", func(t *testing.T) {
		repoMock.EXPECT().DeleteMovie(gomock.Any()).Return(errInternal)

		actualErr := service.DeleteMovie(slug)
		assert.Equal(t, actualErr, errInternal)
	})
}
