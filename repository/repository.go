package repository

import (
	"context"
	"invest_blango_criptal_backend/models"

	"go.mongodb.org/mongo-driver/mongo"
)


type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(user models.SingIn) (*models.User, error)
}


type Account interface {
	ChangePassword(userId int, newPassword string) error
	EditUserData(userId int, newDocs models.UserDocs) error
	UpdateBalance(userId int, amount int64) error
	CreatePromocode(userId int, promocodeName string) error
	AcceptPromocodeUsage(userId int) error
}


type Repository struct {
	Authorization
	Account
}


func NewRepository(db *mongo.Client, ctx *context.Context) *Repository {
	return &Repository{
		Authorization: NewAuthMongo(db, ctx),
		Account: NewAccountMongo(db, ctx),
	}
}