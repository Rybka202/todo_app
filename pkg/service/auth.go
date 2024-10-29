package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"
	"todo_app"
	"todo_app/pkg/repository"

	"github.com/golang-jwt/jwt"
)

const(
	salt = "kjbdiusopqap2685819jbxcb"
	signingKey = "q3ygr1k1873^%&#$%#@12756JGYF<r"
	tokenTTL = 12 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct{
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

func newAuthService(repo repository.Authorization) *AuthService{
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error){
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error){
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil{
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt: time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accesstoken string) (int, error){
	token, err := jwt.ParseWithClaims(accesstoken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil{
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok{
		return 0, errors.New("token claims are not of type")
	}

	return claims.UserID, nil
}

func generatePasswordHash(password string) string{
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}