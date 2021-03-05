package ports

//go:generate mockgen -destination=../../../mocks/mock_services.go -package=mocks github.com/na7r1x/r8r/internal/core/ports UserService,PostService

import "github.com/na7r1x/r8r/api/internal/core/domain"

type UserService interface {
	Login(username string, password string) error
	Register(domain.User) error
}

type PostService interface {
	All() ([]domain.Post, error)
	Create(domain.Post) (string, error)
	One(postId string) (domain.Post, error)
	Update(domain.Post) error
	Delete(postId string) error
	Rate(username string, postId string, rating int) error

	AllImages(postId string) ([]string, error)
	OneImage(imageId string) (string, error)
	AddImage(postId string, image string) (string, error)
	RemoveImage(postId string, imageId string) error
}
