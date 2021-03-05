package userrepo

import (
	"fmt"
	"testing"

	_ "github.com/lib/pq"
	"github.com/na7r1x/r8r/api/internal/core/domain"
)

const (
	host     = "192.168.8.211"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

var psqlconn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
var pgrepo, err = NewPostgresRepo(psqlconn)

func Test_postgresRepo_GetAll(t *testing.T) {
	if err != nil {
		return
	}
	pgrepo.Init()
	tests := []struct {
		name    string
		pg      *postgresRepo
		want    []domain.User
		wantErr bool
	}{
		{"Get all users", pgrepo, []domain.User{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.pg.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresRepo.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}

func Test_postgresRepo_One(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		pg      *postgresRepo
		args    args
		want    domain.User
		wantErr bool
	}{
		{"Get a single user", pgrepo, args{"1"}, domain.User{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.pg.One(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresRepo.One() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Id != tt.args.id {
				t.Errorf("postgresRepo.One() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postgresRepo_Set(t *testing.T) {
	type args struct {
		user domain.User
	}

	testUser := domain.User{
		Username: "test",
		Password: "test",
		Email:    "test@test",
	}

	tests := []struct {
		name    string
		pg      *postgresRepo
		args    args
		wantErr bool
	}{
		{"Insert into users", pgrepo, args{testUser}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.pg.Set(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("postgresRepo.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_postgresRepo_Delete(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		pg      *postgresRepo
		args    args
		wantErr bool
	}{
		{"Delete from users", pgrepo, args{"2"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.pg.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("postgresRepo.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
