package models

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type LikeModel struct {
	DB *mgo.Collection
}

type Like struct {
	ID       bson.ObjectId   `bson:"_id"`
	UserID   bson.ObjectId   `bson:"uid"`
	Articles []bson.ObjectId `bson:"articles"`
}

func (m *LikeModel) AddArticle(userID, articleID string) error {
	return nil
}
