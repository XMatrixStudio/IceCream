package model

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User 用户基本信息
type Comment struct {
	Id_       bson.ObjectId `bson:"_id"`
	ArticleID string        `bson:"articleId"` // 文章ID 【索引】
	UserID    string        `bson:"userId"`    // 评论用户ID 【索引】
	Content   string        `bson:"content"`   // 评论内容
	FatherID  string        `bson:"fatherId"`  // 父评论ID
	LikeNum   int64         `bson:"likeNum"`   // 点赞数
	Top       bool          `bson:"top"`       // 是否置顶
}

var CommentDB *mgo.Collection
