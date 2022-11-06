package service

import (
	"invest_blango_criptal_backend/models"
	"invest_blango_criptal_backend/repository"
)


type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(user models.SingIn) (*models.User, error)
	GenerateJWTToken(user models.SingIn) (string, error)
	ParseJWTToken(accessToken string) (int, string, error)
}


type Account interface {
	ChangePassword(userId int, newPassword string) error
	EditUserData(userId int, newUser models.UserDocs) error
	UpdateBalance(userId int, amount int64) error
	CreatePromocode(userId int, promocodeName string) error
	AcceptPromocodeUsage(userId int) error
}


type Service struct {
	Authorization
	Account
}


func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Account: NewAcountService(repos.Account),
	}
}