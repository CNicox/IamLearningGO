package service

import (
	"errors"
	"gin-web-api-go/config"
	"gin-web-api-go/internal/model"
	"gin-web-api-go/internal/repository"
	"gin-web-api-go/pkg/jwt"

	"github.com/alexedwards/argon2id"
)

type AuthService struct {
	userRepo repository.UserRepository
	config   *config.Config
}

func NewAuthService(r repository.UserRepository, c *config.Config) AuthService {
	return AuthService{
		userRepo: r,
		config:   c,
	}
}

func (r AuthService) RegisterUser(unHashedPassword string, u *model.User) error {
	hash, err := argon2id.CreateHash(unHashedPassword, argon2id.DefaultParams)
	u.PasswordHash = hash
	createUserErr := r.userRepo.CreateUser(*u)
	if createUserErr != nil {
		panic(createUserErr)
	}
	return err
}

func (r AuthService) LoginUser(email, unHashedPassword string) (string, error) {
	user, err := r.userRepo.GetByEmail(email)
	if err != nil {
		return "", err
	}
	match, err := argon2id.ComparePasswordAndHash(unHashedPassword, user.PasswordHash)
	if err != nil {
		return "", err
	}
	token := ""
	if match {
		token, err = jwt.GenerateJWTToken(r.config.JWT.Secret, user.Role, r.config.JWT.TTL, user.Id)
	} else {
		return "", errors.New("invalid credentials")
	}
	if err != nil {
		return "", err
	}
	return token, err
}
