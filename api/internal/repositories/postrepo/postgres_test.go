package postrepo

import (
	"fmt"
	"testing"

	_ "github.com/lib/pq"
	"github.com/na7r1x/r8r/api/internal/core/domain"
)

const (
	host     = "192.168.8.212"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

var psqlconn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
var pgrepo, err = NewPostgresRepo(psqlconn)

func Test_postgresRepo_All(t *testing.T) {
	type fields struct {
		DbLocation string
	}
	tests := []struct {
		name    string
		pg      *postgresRepo
		wantErr bool
	}{
		{"Fetch all posts", pgrepo, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.pg.All()
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresRepo.All() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_postgresRepo_One(t *testing.T) {
	type args struct {
		postId string
	}
	tests := []struct {
		name    string
		pg      *postgresRepo
		args    args
		wantErr bool
	}{
		{"Get a single post", pgrepo, args{"1"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.pg.One(tt.args.postId)
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresRepo.One() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_postgresRepo_Set(t *testing.T) {
	var thisPost = domain.Post{
		CreatedBy:   "admin",
		Type:        "test",
		Title:       "this is a test",
		Description: "test description",
		Rating:      5,
	}
	tests := []struct {
		name    string
		pg      *postgresRepo
		post    domain.Post
		wantErr bool
	}{
		{"Create a post", pgrepo, thisPost, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.pg.Set(tt.post); (err != nil) != tt.wantErr {
				t.Errorf("postgresRepo.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_postgresRepo_Delete(t *testing.T) {

	tests := []struct {
		name    string
		pg      *postgresRepo
		postId  string
		wantErr bool
	}{
		{"Delete a post", pgrepo, "2", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.pg.Delete(tt.postId); (err != nil) != tt.wantErr {
				t.Errorf("postgresRepo.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
