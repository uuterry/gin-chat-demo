package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type UserBasic struct {
	Identity  string `bson:"identity"`
	Account   string `bson:"account"`
	Password  string `bson:"password"`
	Nickname  string `bson:"nickname"`
	Sex       int    `bson:"sex"`
	Email     string `bson:"email"`
	Avatar    string `bson:"avatar"`
	CreatedAt int64  `bson:"create_at"`
	UpdateAt  int64  `bson:"update_at"`
}

func (UserBasic) CollectionName() string {
	return "user_basic"
}

func InsertOneUserBasic(ub *UserBasic) error {
	_, err := MongoDB.Collection(UserBasic{}.CollectionName()).
		InsertOne(context.Background(), ub)
	return err
}

func GetUserBasicByAccountPassword(account, password string) (*UserBasic, error) {
	ub := &UserBasic{}
	err := MongoDB.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(), bson.D{
			{"account", account},
			{"password", password},
		}).Decode(ub)
	return ub, err
}

func GetUserBasicByIdentity(identity string) (*UserBasic, error) {
	ub := &UserBasic{}
	err := MongoDB.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(), bson.D{
			{"identity", identity},
		}).Decode(ub)
	return ub, err
}

func GetUserBasicByAccount(account string) (*UserBasic, error) {
	ub := &UserBasic{}
	err := MongoDB.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(), bson.D{
			{"account", account},
		}).Decode(ub)
	return ub, err
}

func GetUserBasicCountByEmail(email string) (int64, error) {
	return MongoDB.Collection(UserBasic{}.CollectionName()).
		CountDocuments(context.Background(), bson.D{
			{"email", email},
		})
}

func GetUserBasicCountByAccount(account string) (int64, error) {
	return MongoDB.Collection(UserBasic{}.CollectionName()).
		CountDocuments(context.Background(), bson.D{
			{"account", account},
		})
}
