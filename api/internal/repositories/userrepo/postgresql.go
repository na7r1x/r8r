package userrepo

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
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
		return errors.New("could not close database; " + err.Error())
	}
	return nil
}

func (pg postgresRepo) GetAll() ([]domain.User, error) {
	records := []domain.User{}

	rows, err := database.Query("SELECT * FROM users")
	if err != nil {
		return nil, errors.New("failed retrieving users from storage; " + err.Error())
	}

	for rows.Next() {
		var record domain.User

		err := rows.Scan(&record.Id, &record.Username, &record.Password, &record.Email, &record.Created, &record.LastLogin)

		if err != nil {
			return nil, errors.New("failed mapping user from storage; " + err.Error())
		}

		records = append(records, record)
	}
	return records, nil
}

func (pg postgresRepo) One(id string) (domain.User, error) {
	row := database.QueryRow("SELECT * FROM users WHERE user_id = $1", id)

	var record domain.User

	err := row.Scan(&record.Id, &record.Username, &record.Password, &record.Email, &record.Created, &record.LastLogin)
	if err != nil {
		return record, errors.New("failed retrieving user from storage; " + err.Error())
	}

	return record, nil
}

func (pg postgresRepo) Set(user domain.User) error {
	statement, err := database.Prepare("INSERT INTO users (username,password,email,created_on) values($1,$2,$3,NOW()) ON CONFLICT (username) DO UPDATE SET password=$2, email=$3")
	if err != nil {
		return errors.New("failed database upsert; " + err.Error())
	}
	statement.Exec(user.Username, user.Password, user.Email)
	return nil
}

func (pg postgresRepo) Delete(id string) error {
	statement, err := database.Prepare("DELETE FROM users WHERE user_id = $1")
	if err != nil {
		return errors.New("failed to prepare delete statement; " + err.Error())
	}
	result, err := statement.Exec(id)
	if err != nil {
		return errors.New("failed to delete record [" + id + "]; " + err.Error())
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return errors.New("failed to check affected rows; " + err.Error())
	}
	if affected == 0 {
		return errors.New("record [" + id + "] does not exist; ")
	}
	return nil
}
