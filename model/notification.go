package model

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Notification 用户通知
type Notification struct {
	Id_           bson.ObjectId        `bson:"_id"`
	UserID        string               `bson:"userId"`        // 用户ID 【索引】
	Notifications []NotificationDetail `bson:"notifications"` // 通知集合
}

// NotificationDetail 通知详情
type NotificationDetail struct {
	Content string `bson:"content"` // 通知内容
	Read    bool   `bson:"read"`    // 是否以读
	Type    string `bson:"type"`    // 类型： "system", "like", "reply"...
}

// NotificationDB 通知数据库
var NotificationDB *mgo.Collection
