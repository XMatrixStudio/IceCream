package model

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User 用户基本信息
type Content struct {
	Id_        bson.ObjectId `bson:"_id"`        // 文章名字
	Name       string        `bson:"name"`       // 文章名字
	URL        string        `bson:"url"`        // 文章内容MD文件的url地址
	LikeNum    int64         `bson:"likeNum"`    // 点赞人数
	CommentNum int64         `bson:"commentNum"` // 评论次数
	ReadNum    int64         `bson:"readNum"`    // 阅读次数
	Top        bool          `bson:"top"`        // 是否置顶
	Lock       bool          `bson:"lock"`       // 是否锁定
	Comment    bool          `bson:"comment"`    // 是否开放评论
	Type       string        `bson:"type"`       // 类型， 页面还是文章 'page' - 页面 'article' - 文章
}

var ContentDB *mgo.Collection
