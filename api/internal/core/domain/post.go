package domain

import "time"

type Post struct {
	Id          string    `json:"id"`
	Created     time.Time `json:"created"`
	CreatedBy   string    `json:"createdBy"`
	Type        string    `json:"type"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Rating      int       `json:"rating"`
	Images      []string  `json:"images,omitempty"`
}

func NewPost(id string, created time.Time, createdBy string, postType string, title string, description string, rating int) Post {
	return Post{
		Id:          id,
		Created:     created,
		CreatedBy:   createdBy,
		Type:        postType,
		Title:       title,
		Description: description,
		Rating:      rating,
		Images:      make([]string, 0, 10),
	}
}
