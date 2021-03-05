package ports

//go:generate mockgen -destination=../../../mocks/mock_repositories.go -package=mocks github.com/na7r1x/r8r/api/internal/core/ports UserRepository,PostRepository

import "github.com/na7r1x/r8r/api/internal/core/domain"

type UserRepository interface {
	One(string) (domain.User, error)
	Set(domain.User) error
	Delete(string) error
	All() ([]domain.User, error)
}

type PostRepository interface {
	One(string) (domain.Post, error)
	Set(domain.Post) error
	Delete(string) error
	All() ([]domain.Post, error)
	Query(interface{}) ([]domain.Post, error)
}
