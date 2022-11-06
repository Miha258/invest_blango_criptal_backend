package service

import (
	"errors"
	"invest_blango_criptal_backend/models"
	"invest_blango_criptal_backend/repository"
	"time"
	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

const ( 
	salt = "dssidij12jididsakd12ivmcks"
	signingKey = "di78sdja12909sda0sddska021231sad"
	tokenTTL = 12 * time.Hour
)


type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"_id"`
	Password string `json:"password"`
}


type AuthService struct {
	repo repository.Authorization
}



func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}


func (s *AuthService) CreateUser(user models.User) (int, error) {
	hashedPassword, _ := hashPassword(user.Password)
	user.Password = hashedPassword
	return s.repo.CreateUser(user)
}


func (s *AuthService) GetUser(user models.SingIn) (*models.User, error) {
	targert, err := s.repo.GetUser(user)
	
	if err != nil {
		return nil, err
	}

	if !comparePassword(targert.Password, user.Password) {
		return nil, errors.New("Invalid password")
	}

	return targert, nil
}


func (s *AuthService) GenerateJWTToken(user models.SingIn) (string, error) {
	target, err := s.GetUser(user)

	if err != nil {
		logrus.Fatal(err)
		return "", err
	}
	
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt: time.Now().Unix(),
		},
		UserId: target.Id,
		Password: user.Password,
	})
	
	return jwtToken.SignedString([]byte(signingKey))
}	


func (s *AuthService) ParseJWTToken(accessToken string) (int, string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error)  {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid singing method")
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, "", err
	}
	
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, "", errors.New("Invalid token claims")
	}
	return claims.UserId, claims.Password, err
}


func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 9)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
} 


func comparePassword(hashedPassword string, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		logrus.Error(err)
		return false
	}
	return true
}
