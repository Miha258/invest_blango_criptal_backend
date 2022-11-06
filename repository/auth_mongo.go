package repository

import (
	"context"
	"errors"
	"invest_blango_criptal_backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


type AuthMongo struct {
	db *mongo.Client
	ctx *context.Context
}


func NewAuthMongo (db *mongo.Client, ctx *context.Context) *AuthMongo {
	return &AuthMongo{db: db, ctx: ctx}
}


func (r *AuthMongo) GetUser(user models.SingIn) (*models.User, error) {
	var target models.User
	users := r.db.Database("invest_site").Collection("users")

	if err := users.FindOne(*r.ctx, bson.D{{"login", user.Login}}).Decode(&target); err != nil {
		return nil, errors.New("User not found")
	}

	return &target, nil
}


func (r *AuthMongo) CreateUser(user models.User) (int, error) {
	users := r.db.Database("invest_site").Collection("users")
	user_id, _ := users.CountDocuments(*r.ctx, bson.D{})
	
	exists, _ := users.CountDocuments(*r.ctx, bson.D{{"login", user.Login}})
	
	if exists != 0 {
		return -1, errors.New("User already exists")
	}	
	
	_, err := users.InsertOne(*r.ctx, bson.D{{"_id", user_id + 1}, {"email", user.Email}, {"password", user.Password}, {"phone", user.Phone}, {"login", user.Login}, {"balance", 0}, {"crypto_balance", 0}, {"using_promo", user.UsingPromo}, {"documents", user.Documents}, {"promo_code", nil}})

	if err != nil {
		return -1, err	
	}

	return int(user_id), nil
}



