package models

import (
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type CommentModel struct {
	DB *mgo.Collection
}

// Comment 评论
type Comment struct {
	ID        bson.ObjectId `bson:"_id"`       // 评论ID 【索引】
	ArticleID bson.ObjectId `bson:"articleId"` // 文章ID 【索引】
	UserID    bson.ObjectId `bson:"userId"`    // 评论用户ID 【索引】
	Date      int64         `bson:"date"`      // 发布时间
	Content   string        `bson:"content"`   // 评论内容
	FatherID  bson.ObjectId `bson:"fatherId"`  // 父评论ID
	LikeNum   int64         `bson:"likeNum"`   // 点赞数
	Top       bool          `bson:"top"`       // 是否置顶
}

func (m *CommentModel) AddComment(articeID, userID, fatherID bson.ObjectId, content string) (err error) {
	newComment := bson.NewObjectId()
	return m.DB.Insert(Comment{
		ID:        newComment,
		ArticleID: articeID,
		UserID:    userID,
		FatherID:  fatherID,
		Date:      time.Now().Unix() * 1000,
		Content:   content,
		LikeNum:   0,
		Top:       false,
	})
}

func (m *CommentModel) GetCommentByArticleID(articeID bson.ObjectId) (comments []Comment, err error) {
	err = m.DB.Find(bson.M{"articleId": articeID}).All(&comments)
	return
}
