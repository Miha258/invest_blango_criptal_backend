package repository

import (
	"context"
	"errors"
	"invest_blango_criptal_backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


type AccountMongo struct {
	db *mongo.Client
	ctx *context.Context
}


func NewAccountMongo (db *mongo.Client, ctx *context.Context) *AccountMongo {
	return &AccountMongo{db: db, ctx: ctx}
}


func (r *AccountMongo) ChangePassword(userId int, newPassword string) error {
	users := r.db.Database("invest_site").Collection("users")
	_, err := users.UpdateOne(*r.ctx, bson.D{{"_id", userId}}, bson.M{"$set": bson.D{{"password", newPassword}}})

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}



func (r *AccountMongo) EditUserData(userId int, newDocs models.UserDocs) error {
	users := r.db.Database("invest_site").Collection("users")
	_, err := users.UpdateOne(*r.ctx, bson.D{{"_id", userId}}, bson.M{"$set": bson.D{{"documents", &newDocs}}})

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}


func (r *AccountMongo) UpdateBalance(userId int, amount int64) error {
	users := r.db.Database("invest_site").Collection("users")
	_, err := users.UpdateOne(*r.ctx, bson.D{{"_id", userId}}, bson.M{"$inc": bson.D{{"balance", amount}}})

	if err != nil {
		return errors.New(err.Error())
	}

	return nil 
} 


func (r *AccountMongo) CreatePromocode(userId int, promocodeName string) error {
	users := r.db.Database("invest_site").Collection("users")

	_, err := users.UpdateOne(*r.ctx, bson.D{{"_id", userId}}, bson.M{"$set": bson.D{{"promo_code", models.UserPromo{Name: promocodeName, Count: 0}}}})

	if err != nil {
		return errors.New(err.Error())
	}

	return nil 
} 


func (r *AccountMongo) AcceptPromocodeUsage(userId int) error {
	var targert models.User
	users := r.db.Database("invest_site").Collection("users")

	if err := users.FindOne(*r.ctx, bson.D{{"_id", userId}}).Decode(&targert); err != nil {
		return errors.New(err.Error())
	}

	if targert.UsingPromo != "" {
		newUserPromocodeData := models.UserPromo{Name: targert.PromoCode.Name, Count:  targert.PromoCode.Count + 1}
		_, err := users.UpdateOne(*r.ctx, bson.D{{"_id", userId}}, bson.M{"$set": bson.D{{"promo_code", newUserPromocodeData}}})

		if err != nil {
			return errors.New(err.Error())
		}

		return nil 
	}
	return errors.New("User haven`t used promocode")
} 