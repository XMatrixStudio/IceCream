package models

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type WebsiteModel struct {
	DB *mgo.Collection
}

type Website struct {
	ID       bson.ObjectId `bson:"_id"`
	Name     string        `bson:"name"`
	URL      string        `bson:"url"`
	Text     string        `bson:"text"`
	Articles int           `bson:"articles"`
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
