package usersrv

import (
	"fmt"
	"testing"

	"github.com/na7r1x/r8r/api/internal/core/domain"
	"github.com/na7r1x/r8r/api/internal/repositories/userrepo"
)

const (
	host     = "192.168.8.211"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

var psqlconn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
var pgrepo, err = userrepo.NewPostgresRepo(psqlconn)

func Test_service_Register(t *testing.T) {
	type args struct {
		user domain.User
	}
	tests := []struct {
		name    string
		srv     service
		args    args
		wantErr bool
	}{
		{"Attempt to register empty user", New(pgrepo), args{domain.User{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.srv.Register(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("service.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
