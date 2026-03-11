package repository

import (
	"gin-web-api-go/internal/model"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (r UserRepository) CreateUser(u model.User) error {
	_, err := r.db.Exec(`INSERT INTO users 
	(username, email, password_hash, role, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6)`,
		u.Username, u.Email, u.PasswordHash, u.Role, u.CreatedAt, u.UpdatedAt,
	)
	return err
}

func (r UserRepository) GetByEmail(email string) (model.User, error) {
	user := model.User{}
	err := r.db.Get(&user, "SELECT * FROM users WHERE email=$1", email)
	return user, err
}

func (r UserRepository) GetByID(id uuid.UUID) (model.User, error) {
	user := model.User{}
	err := r.db.Get(&user, "SELECT * FROM users WHERE id=$1", id)
	return user, err
}

func (r UserRepository) GetAll(page int, limit int, role string) ([]model.User, error) {
	var users []model.User
	var err error
	page = (page - 1) * limit
	if role != "" {
		err = r.db.Select(&users, "SELECT * FROM users WHERE role=$1 LIMIT $2 OFFSET $3;", role, limit, page)
	} else {
		err = r.db.Select(&users, "SELECT * FROM users LIMIT $1 OFFSET $2;", limit, page)
	}
	return users, err
}
