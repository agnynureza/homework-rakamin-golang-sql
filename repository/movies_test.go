package repository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/agnynureza/homework-rakamin-golang-sql/models"
	"github.com/tj/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMockDB(t *testing.T) (*MoviesRepository, sqlmock.Sqlmock, *sql.DB) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Error initialize DATA-DOG : %v", err)
	}
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
	repo := NewMoviesRepository(gormDB)

	return repo, mock, db
}

func TestMoviesRepository_CreateMovie(t *testing.T) {
	repo, mock, db := InitMockDB(t)
	defer db.Close()

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

		mock.ExpectExec(queryExpected).
			WillReturnResult(sqlmock.NewResult(1, 1)).
			WillReturnError(nil)

		actualMovieID, err := repo.CreateMovie(&inputMovie)
		assert.NoError(t, err)
		assert.Equal(t, actualMovieID, outputMovieID)
	})
}

func TestMoviesRepository_GetOneMovie(t *testing.T) {
	repo, mock, db := InitMockDB(t)
	defer db.Close()

	t.Run("should return success", func(t *testing.T) {
		ouputMovieDetail := models.Movies{
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
			WillReturnRows(sqlmock.NewRows([]string{"title", "Slug", "Description", "Duration", "Image"}).
				AddRow(ouputMovieDetail.Title, ouputMovieDetail.Slug, ouputMovieDetail.Description, ouputMovieDetail.Duration, ouputMovieDetail.Image)).
			WillReturnError(nil)

		actual, err := repo.GetOneMovie(inputSlug)
		assert.Nil(t, err)
		assert.Equal(t, actual, ouputMovieDetail)
	})
}
