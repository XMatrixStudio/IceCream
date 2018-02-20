package model

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User 用户基本信息
type Like struct {
	Id_ bson.ObjectId `bson:"_id"`
	ID  string        `bson:"id"`  // ID 【索引】(ID可以是文章，可以为评论，甚至可以是用户(关注系统))
	IDs []string      `bson:"ids"` // ID集合
}

var (
	ContentLikeDB *mgo.Collection // 点赞文章或页面的用户ID集合
	CommentLikeDB *mgo.Collection // 点赞评论的用户ID集合
	UserLikeDB    *mgo.Collection // 用户点赞文章或评论的ID的集合
)
