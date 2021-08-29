package repository

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/agnynureza/homework-rakamin-golang-sql/models"
	"github.com/tj/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	repo *MoviesRepository
	mock sqlmock.Sqlmock
	db   *sql.DB
	err  error
)

func TestMain(m *testing.M) {
	db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("Error initialize DATA-DOG : %v", err)
	}
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{Conn: db, SkipInitializeWithVersion: true}), &gorm.Config{})
	repo = NewMoviesRepository(gormDB)

	exitVal := m.Run()

	defer db.Close()

	os.Exit(exitVal)
}

func TestMoviesRepository_CreateMovie(t *testing.T) {
	t.Run("should return success", func(t *testing.T) {
		inputMovie := models.Movies{
			Title:       "Titanic",
			Slug:        "titanic",
			Description: "lorem ipsum",
			Duration:    60,
			Image:       "image titanic URL",
		}
		outputMovieID := 1
		queryExpected := "INSERT INTO `movies` (`title`,`slug`,`description`,`duration`,`image`) VALUES (?,?,?,?,?)"

		mock.ExpectBegin()
		mock.ExpectExec(queryExpected).
			WillReturnResult(sqlmock.NewResult(1, 1)).
			WillReturnError(nil)
		mock.ExpectCommit()

		actualMovieID, err := repo.CreateMovie(&inputMovie)
		assert.NoError(t, err)
		assert.Equal(t, actualMovieID, outputMovieID)
	})
}

func TestMoviesRepository_GetOneMovie(t *testing.T) {
	t.Run("should return success", func(t *testing.T) {
		ouputMovieDetail := models.Movies{
			ID:          1,
			Title:       "Titanic",
			Slug:        "titanic",
			Description: "lorem ipsum",
			Duration:    60,
			Image:       "image titanic URL",
		}
		inputSlug := "titanic"
		queryExpected := `SELECT id, title, slug, description, duration, image FROM movies WHERE slug = ?`

		mock.ExpectQuery(queryExpected).
			WithArgs(inputSlug).
			WillReturnRows(sqlmock.NewRows([]string{"id", "title", "Slug", "Description", "Duration", "Image"}).
				AddRow(ouputMovieDetail.ID, ouputMovieDetail.Title, ouputMovieDetail.Slug, ouputMovieDetail.Description, ouputMovieDetail.Duration, ouputMovieDetail.Image)).
			WillReturnError(nil)

		actual, err := repo.GetOneMovie(inputSlug)
		assert.Nil(t, err)
		assert.Equal(t, actual, ouputMovieDetail)
	})
}

func TestMoviesRepository_UpdateMovie(t *testing.T) {
	t.Run("should return success", func(t *testing.T) {
		inputMovie := models.Movies{
			Title:       "Titanic version 2",
			Slug:        "boboboy",
			Description: "lorem ipsum",
			Duration:    60,
			Image:       "image titanic URL",
		}
		inputSlug := "titanic"
		queryExpected := "UPDATE movies SET title = ?, description = ?, duration = ?, image = ?, slug = ? WHERE slug = ?"

		mock.ExpectExec(queryExpected).
			WillReturnResult(sqlmock.NewResult(1, 1)).
			WillReturnError(nil)

		err := repo.UpdateMovie(&inputMovie, inputSlug)
		assert.NoError(t, err)
	})
}

func TestMoviesRepository_DeleteMovie(t *testing.T) {
	t.Run("should return success", func(t *testing.T) {
		inputSlug := "titanic"
		queryExpected := "DELETE FROM `movies` WHERE slug = ?"

		mock.ExpectBegin()
		mock.ExpectExec(queryExpected).
			WillReturnResult(sqlmock.NewResult(1, 1)).
			WillReturnError(nil)
		mock.ExpectCommit()

		err := repo.DeleteMovie(inputSlug)
		assert.NoError(t, err)
	})
}
