package models

import (
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type LogModel struct {
	DB *mgo.Collection
}

// Log 记录
type Log struct {
	ID     bson.ObjectId `bson:"_id"`  //记录ID
	UserID bson.ObjectId `bson:"uid"`  //用户ID
	Logs   []LogRecord   `bson:"logs"` //记录内容数组
}

// LogRecord 记录内容
type LogRecord struct {
	Date  int64  `bson:"date"`  //记录时间
	State bool   `bson:"state"` //记录状态
	Msg   string `bson:"msg"`   //记录内容文本
}

func (m *LogModel) AddLogDocument(userID string) (err error) {
	newLog := bson.NewObjectId()
	err = m.DB.Insert(&Log{
		ID:     newLog,
		UserID: bson.ObjectIdHex(userID),
		Logs:   []LogRecord{},
	})
	return
}

func (m *LogModel) AddLogRecord(userID, msg string) (err error) {
	err = m.DB.Update(
		bson.M{"uid": bson.ObjectIdHex(userID)},
		bson.M{"$push": bson.M{"logs": LogRecord{
			Date:  time.Now().Unix() * 1000,
			State: true,
			Msg:   msg,
		}}},
	)
	return
}
