package model

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Content 内容
type Content struct {
	Id_        bson.ObjectId `bson:"_id"`
	Name       string        `bson:"name"`       // 文章名字
	URL        string        `bson:"url"`        // 文章内容文件的url地址
	LikeNum    int64         `bson:"likeNum"`    // 点赞人数
	CommentNum int64         `bson:"commentNum"` // 评论次数
	ReadNum    int64         `bson:"readNum"`    // 阅读次数
	Top        bool          `bson:"top"`        // 是否置顶
	Lock       bool          `bson:"lock"`       // 是否锁定
	Comment    bool          `bson:"comment"`    // 是否开放评论
	Type       string        `bson:"type"`       // 类型， 页面还是文章 'page' - 页面 'article' - 文章
}

// ContentDB 内容数据库
var ContentDB *mgo.Collection

// AddContent 增加内容
func AddContent(name, url, contentType string, isComment bool) (bson.ObjectId, error) {
	newContent := bson.NewObjectId()
	err := ContentDB.Insert(&Content{
		Id_:     newContent,
		Name:    name,
		URL:     url,
		Comment: isComment,
		Type:    contentType,
	})
	if err != nil {
		return "", err
	}
	return newContent, nil
}

// AddNum 增加一个阅读("readNum")/点赞("likeNum")/评论数("commentNum")
func AddNum(id, name string) bool {
	_, err := ContentDB.UpsertId(bson.ObjectIdHex(id), bson.M{"$inc": bson.M{name: 1}})
	if err != nil {
		return false
	}
	return true
}

// SetStatus 设置置顶("top")/评论("comment")/锁定("lock")状态
func SetStatus(id, name string, status bool) bool {
	_, err := ContentDB.UpsertId(bson.ObjectIdHex(id), bson.M{"$set": bson.M{name: status}})
	if err != nil {
		return false
	}
	return true
}
