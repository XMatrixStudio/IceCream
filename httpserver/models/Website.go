package models

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type WebsiteModel struct {
	DB *mgo.Collection
}

// Website 网站信息
type Website struct {
	ID       bson.ObjectId `bson:"_id"`      // ID
	Name     string        `bson:"name"`     // 网站显示名字
	URL      string        `bson:"url"`      // 网站显示URL
	Text     string        `bson:"text"`     // 网站关于内容
	Articles int           `bson:"articles"` // 文章数量
}

func (m *WebsiteModel) InitWebsiteInfo(name, url string) error {
	count, err := m.DB.Find(nil).Count()
	if err != nil || count != 0 {
		return err
	}
	newWebsite := bson.NewObjectId()
	return m.DB.Insert(&Website{
		ID:       newWebsite,
		Name:     name,
		URL:      url,
		Articles: 0,
	})
}

func (m *WebsiteModel) GetWebsiteInfo() (website Website, err error) {
	err = m.DB.Find(nil).One(&website)
	return
}

func (m *WebsiteModel) IncOrDecArticlesNum(num int) error {
	return m.DB.Update(nil, bson.M{"$inc": bson.M{"articles": num}})
}

func (m *WebsiteModel) EditWebsiteInfo(name, url, text string) error {
	return m.DB.Update(nil, bson.M{
		"$set": bson.M{
			"name": name,
			"url":  url,
			"text": text,
		},
	})
}
