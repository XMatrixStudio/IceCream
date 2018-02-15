package model

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User 用户基本信息
type User struct {
	Id_   bson.ObjectId `bson:"_id"`
	Name  string        `bson:"name"`
	Class int           `bson:"class"`
	Token string        `bson:"token"`
	Info  UserInfo      `bson:"info"`
}

// UserInfo 用户个性信息
type UserInfo struct {
	Avatar    string `bson:"avatar"`
	Bio       string `bson:"bio"`
	BirthDate string `bson:"birthDate"`
	Email     string `bson:"email"`
	Gender    int    `bson:"gender"`
	Location  string `bson:"location"`
	NikeName  string `bson:"nikeName"`
	Phone     string `bson:"phone"`
	URL       string `bson:"url"`
}

// UserDB 用户数据库
var UserDB *mgo.Collection

// AddUser 添加用户
func AddUser() (bson.ObjectId, error) {
	newUser := bson.NewObjectId()
	err := UserDB.Insert(&User{
		Id_: bson.NewObjectId(),
	})
	if err != nil {
		return "", err
	}
	return newUser, nil
}
