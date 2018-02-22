package model

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Comment 评论
type Comment struct {
	Id_       bson.ObjectId `bson:"_id"`
	ArticleID string        `bson:"articleId"` // 文章ID 【索引】
	UserID    string        `bson:"userId"`    // 评论用户ID 【索引】
	Content   string        `bson:"content"`   // 评论内容
	FatherID  string        `bson:"fatherId"`  // 父评论ID
	LikeNum   int64         `bson:"likeNum"`   // 点赞数
	Top       bool          `bson:"top"`       // 是否置顶
}

// CommentDB 评论数据库
var CommentDB *mgo.Collection

// AddComment 增加评论
func AddComment(article, user, content, fatherID string) (bson.ObjectId, error) {
	newComment := bson.NewObjectId()
	err := CommentDB.Insert(&Comment{
		Id_:       newComment,
		ArticleID: article,
		UserID:    user,
		Content:   content,
		FatherID:  fatherID,
	})
	if err != nil {
		return "", err
	}
	return newComment, nil
}
