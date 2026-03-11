package repository

import (
	"gin-web-api-go/internal/model"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	sq "github.com/Masterminds/squirrel"
)

type PostRepository struct {
	db *sqlx.DB
}

func NewPostRepository(db *sqlx.DB) PostRepository {
	return PostRepository{
		db: db,
	}
}

func (r PostRepository) CreatePost(p model.Post) error {
	_, err := r.db.Exec(`INSERT INTO posts 
	(title, content, author_id, status, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6)`,
		p.Title, p.Content, p.AuthorId, p.Status, p.CreatedAt, p.UpdatedAt,
	)
	return err
}

func (r PostRepository) EditPost(p model.Post, id uuid.UUID) error {
	_, err := r.db.Exec(`
	UPDATE posts 
	SET title=$1, content=$2, status=$3, updated_at=$4
	WHERE id=$5`,
		p.Title, p.Content, p.Status, p.UpdatedAt, id,
	)
	return err
}

func (r PostRepository) ChangePostStatus(id uuid.UUID, status string, updatedAt time.Time) error {
	_, err := r.db.Exec(`
	UPDATE posts 
	SET status=$1, updated_at=$2
	WHERE id=$3`,
		status, updatedAt, id,
	)
	return err
}

func (r PostRepository) DeletePost(id uuid.UUID) error {
	_, err := r.db.Exec("DELETE FROM posts WHERE id=$1", id)
	return err
}

func (r PostRepository) GetByID(id uuid.UUID) (model.Post, error) {
	post := model.Post{}
	err := r.db.Get(&post, "SELECT * FROM posts WHERE id=$1", id)
	return post, err
}

func (r PostRepository) GetAll(page int, limit int, status string, author_id uuid.UUID, searchContent string) ([]model.Post, error) {
	var posts []model.Post
	page = (page - 1) * limit
	query := sq.Select("*").
		From("posts").
		Limit(uint64(limit)).
		Offset(uint64(page)).
		PlaceholderFormat(sq.Dollar)

	if status != "" {
		query = query.Where(sq.Eq{"status": status})
	}
	if len(author_id) != 0 {
		query = query.Where(sq.Eq{"author_id": author_id})
	}
	if searchContent != "" {
		query = query.Where(sq.Or{
			sq.ILike{"title": "%" + searchContent + "%"},
			sq.ILike{"content": "%" + searchContent + "%"},
		})
	}
	sql, args, err := query.ToSql()
	err = r.db.Select(&posts, sql, args...)

	return posts, err
}
