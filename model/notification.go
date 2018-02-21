package model

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Notication 用户通知
type Notification struct {
	Id_     bson.ObjectId `bson:"_id"`
	UserID  string        `bson:"userId"`  // 用户ID 【索引】
	Content string        `bson:"content"` // 通知内容
	Read    bool          `bson:"read"`    // 是否以读
	Type    string        `bson:"type"`    // 类型： "system", "like", "reply"...
}

var NotificationDB *mgo.Collection
