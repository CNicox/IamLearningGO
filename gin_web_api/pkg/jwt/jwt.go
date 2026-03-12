package jwt

import (
	"time"

	jwtlib "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type MyCustomClaims struct {
	Role string `json:"role"`
	jwtlib.RegisteredClaims
}

func GenerateJWTToken(mySigningKey string, role string, ttl time.Duration, sub uuid.UUID) (string, error) {
	claims := MyCustomClaims{
		Role: role,
		RegisteredClaims: jwtlib.RegisteredClaims{
			ExpiresAt: jwtlib.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt:  jwtlib.NewNumericDate(time.Now()),
			Subject:   sub.String(),
		},
	}
	token := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(mySigningKey))
	return ss, err
}
