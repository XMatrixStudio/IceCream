package model

import (
	"fmt"

	"github.com/XMatrixStudio/icecream/config"
	"gopkg.in/mgo.v2"
)

// DB 数据库实例
var DB *mgo.Database

// InitMongo 初始化连接
func InitMongo() error {
	session, err := mgo.Dial(
		"mongodb://" +
			config.Mongo.User +
			":" + config.Mongo.Password +
			"@" + config.Mongo.Host +
			":" + config.Mongo.Port +
			"/" + config.Mongo.Name)
	if err != nil {
		return err
	}
	fmt.Println("Mongodb Connect Success")
	DB = session.DB(config.Mongo.Name)
	UserDB = DB.C("users")
	ContentDB = DB.C("content")
	CommentDB = DB.C("comment")
	ContentLikeDB = DB.C("like")
	CommentLikeDB = DB.C("commentLike")
	UserLikeDB = DB.C("userLike")
	NotificationDB = DB.C("notification")
	return nil
}
