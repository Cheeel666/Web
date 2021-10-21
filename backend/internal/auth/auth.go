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
	signingKey     []byte
	expireDuration time.Duration
}

// NewAuthorizer object
func NewAuthorizer(repo database.DbConnection, hashSalt string, signingKey []byte, expireDuration time.Duration) *Authorizer {
	return &Authorizer{
		repo:           repo,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
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
	})
	str, err := token.SignedString(a.signingKey)
	return user, str, err
}
