package model

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User 用户基本信息
type User struct {
	Id_   bson.ObjectId `bson:"_id"`   // 用户ID
	Name  string        `bson:"name"`  // 用户唯一名字
	Class string        `bson:"class"` // 用户类型 "admin", "reader", "writer"
	Info  UserInfo      `bson:"info"`  // 用户个性信息
}

// UserInfo 用户个性信息
type UserInfo struct {
	Avatar    string `bson:"avatar"`    // 头像URL
	Bio       string `bson:"bio"`       // 个人简介
	BirthDate string `bson:"birthDate"` // 生日
	Email     string `bson:"email"`     // 邮箱
	Gender    int    `bson:"gender"`    // 性别
	Location  string `bson:"location"`  // 所在地
	NikeName  string `bson:"nikeName"`  // 昵称
	Phone     string `bson:"phone"`     // 手机号码
	URL       string `bson:"url"`       // 个人主页
}

// UserDB 用户数据库
var UserDB *mgo.Collection

// AddUser 添加用户
func AddUser() (bson.ObjectId, error) {
	newUser := bson.NewObjectId()
	err := UserDB.Insert(&User{
		Id_:   newUser,
		Name:  "user_" + string(newUser),
		Class: "reader",
	})
	if err != nil {
		return "", err
	}
	return newUser, nil
}

// SetUserInfo 设置用户信息
func SetUserInfo(id string, info UserInfo) bool {
	data := bson.M{"$set": info}
	_, err := UserDB.UpsertId(bson.ObjectIdHex(id), data)
	if err != nil {
		return false
	}
	return true
}

// SetUserName 设置用户名
func SetUserName(id, name string) bool {
	_, err := UserDB.UpsertId(bson.ObjectIdHex(id), bson.M{"$set": bson.M{"name": name}})
	if err != nil {
		return false
	}
	return true
}

// SetUserClass 设置用户类型
func SetUserClass(id, class string) bool {
	_, err := UserDB.UpsertId(bson.ObjectIdHex(id), bson.M{"$set": bson.M{"class": class}})
	if err != nil {
		return false
	}
	return true
}

// GetByID 根据ID查询用户
func GetByID(id string) *User {
	user := new(User)
	err := UserDB.FindId(id).One(&user)
	if err != nil {
		return nil
	}
	return user
}
