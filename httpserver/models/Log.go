package models

import (
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type LogModel struct {
	DB *mgo.Collection
}

type Log struct {
	ID     bson.ObjectId `bson:"_id"`
	UserID bson.ObjectId `bson:"uid"`
	Logs   []LogRecord   `bson:"logs"`
}

type LogRecord struct {
	Date  int64  `bson:"date"`
	State bool   `bson:"state"`
	Msg   string `bson:"msg"`
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
