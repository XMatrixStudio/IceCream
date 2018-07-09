package models

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type LikeModel struct {
	DB *mgo.Collection
}

// Like 点赞
type Like struct {
	ID       bson.ObjectId   `bson:"_id"`      // 点赞ID
	UserID   bson.ObjectId   `bson:"uid"`      // 用户ID
	Articles []bson.ObjectId `bson:"articles"` // 文章ID数组
}

func (m *LikeModel) AddDocument(userID string) (err error) {
	count, err := m.DB.Find(bson.M{"uid": bson.ObjectIdHex(userID)}).Count()
	if err != nil || count != 0 {
		return
	}
	newLike := bson.NewObjectId()
	m.DB.Insert(Like{
		ID:       newLike,
		UserID:   bson.ObjectIdHex(userID),
		Articles: []bson.ObjectId{},
	})
	return
}

func (m *LikeModel) IsLike(userID, articleID bson.ObjectId) (isLike bool, err error) {
	count, err := m.DB.Find(bson.M{
		"uid":      userID,
		"articles": bson.M{"$all": []bson.ObjectId{articleID}},
	}).Count()
	if err != nil {
		return false, err
	}
	return count == 1, nil
}

func (m *LikeModel) AddArticle(userID, articleID string) (err error) {
	err = m.AddDocument(userID)
	if err != nil {
		return
	}
	m.DB.Update(
		bson.M{"uid": bson.ObjectIdHex(userID)},
		bson.M{"$push": bson.M{"articles": bson.ObjectIdHex(articleID)}},
	)
	return nil
}

func (m *LikeModel) RemoveArticle(userID, articleID string) (err error) {
	return m.DB.Update(
		bson.M{"uid": bson.ObjectIdHex(userID)},
		bson.M{"$pull": bson.M{"articles": bson.ObjectIdHex(articleID)}},
	)
}
