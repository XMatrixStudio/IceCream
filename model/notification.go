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
	Id_     bson.ObjectId `bson:"_id"`
	Content string        `bson:"content"` // 通知内容
	Read    bool          `bson:"read"`    // 是否以读
	Type    string        `bson:"type"`    // 类型： "system", "like", "reply"...
}

// NotificationDB 通知数据库
var NotificationDB *mgo.Collection

// InitNotification 初始化用户通知
func InitNotification(user string) error {
	newNotification := bson.NewObjectId()
	err := NotificationDB.Insert(&Comment{
		Id_:    newNotification,
		UserID: user,
	})
	return err
}

// AddNotification 添加一条通知 类型:"system", "like", "reply" ...
func AddNotification(content, user, notificationType string) error {
	newNotification := &NotificationDetail{
		Id_:     bson.NewObjectId(),
		Content: content,
		Read:    false,
		Type:    notificationType,
	}
	err := NotificationDB.Update(
		bson.M{"userId": bson.ObjectIdHex(user)},
		bson.M{"$push": bson.M{"notifications": &newNotification}})
	return err
}

// ReadANotification 标记通知
func ReadANotification(user, id string, status bool) error {
	err := NotificationDB.Update(
		bson.M{"userId": bson.ObjectIdHex(user), "notifications._id": id},
		bson.M{"$set": bson.M{"notifications.$.read": status}})
	return err
}

// RemoveANotification 删除通知
func RemoveANotification(user, id string) error {
	err := NotificationDB.Update(
		bson.M{"userId": bson.ObjectIdHex(user), "notifications._id": id},
		bson.M{"$pull": bson.M{"notifications._id": id}})
	return err
}

// GetNotificationsByUser 获取用户所有通知
func GetNotificationByUser(user string) ([]NotificationDetail, error) {
	var notifications []NotificationDetail
	err := NotificationDB.Find(nil).All(notifications)
	if err != nil {
		return nil, err
	}
	return notifications, nil
}
