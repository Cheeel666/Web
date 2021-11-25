package auth

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"
	"web/internal/database"
	"web/internal/model"

	"github.com/dgrijalva/jwt-go/v4"
)

// Authorizer struct
type Authorizer struct {
	repo database.DbConnection

	hashSalt       string
	SigningKey     []byte
	expireDuration time.Duration
}

// NewAuthorizer object
func NewAuthorizer(repo database.DbConnection, hashSalt string, signingKey []byte, expireDuration time.Duration) *Authorizer {
	return &Authorizer{
		repo:           repo,
		hashSalt:       hashSalt,
		SigningKey:     signingKey,
		expireDuration: expireDuration,
	}
}

// SignUp to serice
func (a *Authorizer) SignUp(ctx context.Context, user model.User) error {
	// Create password hash
	pwd := sha1.New()
	pwd.Write([]byte(user.Password))
	pwd.Write([]byte(a.hashSalt))
	user.Password = fmt.Sprintf("%x", pwd.Sum(nil))
	a.repo.AddUsrInfo(user)
	return nil
}

// SignIn to service
func (a *Authorizer) SignIn(ctx context.Context, user model.User) (model.User, string, error) {
	pwd := sha1.New()
	pwd.Write([]byte(user.Password))
	pwd.Write([]byte(a.hashSalt))
	user.Password = fmt.Sprintf("%x", pwd.Sum(nil))

	user, err := a.repo.GetUserByEmail(user)
	if err != nil {
		return user, "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(a.expireDuration)),
			IssuedAt:  jwt.At(time.Now()),
		},
		Username: user.Name,
		Role:     user.Role,
	})

	str, err := token.SignedString(a.SigningKey)
	return user, str, err
}

func ParseToken(accessToken string, signingKey []byte) (bool, interface{}) {
	fmt.Println(accessToken)
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	if err != nil {
		return false, nil
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return true, claims
	}
	// fmt.Println("Invalid token")
	return false, nil
}
