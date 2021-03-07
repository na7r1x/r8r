package postrepo

import (
	"database/sql"
	"errors"

	"github.com/na7r1x/r8r/api/internal/core/domain"
)

type postgresRepo struct {
	DbLocation string
}

var database *sql.DB

func NewPostgresRepo(loc string) (*postgresRepo, error) {
	pg := new(postgresRepo)
	pg.DbLocation = loc
	err := pg.Init()
	if err != nil {
		return nil, err
	}
	return pg, nil
}

func (pg postgresRepo) Init() error {
	var err error
	database, err = sql.Open("postgres", pg.DbLocation)
	if err != nil {
		return err
	}
	return nil
}

func (pg postgresRepo) Destroy() error {
	err := database.Close()
	if err != nil {
		return err
	}
	return nil
}

func (pg postgresRepo) All() ([]domain.Post, error) {
	records := []domain.Post{}
	rows, err := database.Query("SELECT * FROM posts")
	if err != nil {
		return nil, errors.New("error retrieving posts from storage; reason: " + err.Error())
	}
	for rows.Next() {
		var record domain.Post
		err := rows.Scan(&record.Id, &record.Created, &record.CreatedBy, &record.Type, &record.Title, &record.Description)
		if err != nil {
			return nil, errors.New("failed mapping db record to object domain.Post; reason: " + err.Error())
		}
		records = append(records, record)
	}
	return records, nil
}

func (pg postgresRepo) One(postId string) (domain.Post, error) {
	row := database.QueryRow("SELECT * FROM posts WHERE post_id = $1", postId)
	var record domain.Post
	err := row.Scan(&record.Id, &record.Created, &record.CreatedBy, &record.Type, &record.Title, &record.Description)
	if err != nil {
		return record, errors.New("error retrieving post [" + postId + "] from storate; reason: " + err.Error())
	}
	return record, nil
}

func (pg postgresRepo) Set(post domain.Post) error {
	return errors.New("not implemented")
}

func (pg postgresRepo) Delete(postId string) error {
	return errors.New("not implemented")
}
