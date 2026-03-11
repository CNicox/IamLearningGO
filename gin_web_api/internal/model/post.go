package model

import (
	"github.com/google/uuid"
	"time"
)

type PostStatus string

const (
	Draft     PostStatus = "draft"
	Published PostStatus = "published"
)

type Post struct {
	Id        uuid.UUID  `db:"id"`
	Title     string     `db:"title"`
	Content   string     `db:"content"`
	AuthorId  uuid.UUID  `db:"author_id"`
	Status    PostStatus `db:"status"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
}
