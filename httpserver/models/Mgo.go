package models

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/XMatrixStudio/IceCream/generator"
	"github.com/globalsign/mgo"
)

// Mongo 数据库配置
type Mongo struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	Name     string `yaml:"Name"`
}

// Model 模型
type Model struct {
	Config  Mongo
	DB      *mgo.Database
	User    UserModel
	Article ArticleModel
	Log     LogModel
	Website WebsiteModel
	Like    LikeModel
}

// initMongo 初始化数据库
func (m *Model) initMongo(conf Mongo) error {
	m.Config = conf
	if m.DB != nil {
		m.DB.Session.Close()
	}
	session, err := mgo.Dial(
		"mongodb://" +
			conf.User +
			":" + conf.Password +
			"@" + conf.Host +
			":" + conf.Port +
			"/" + conf.Name)
	if err != nil {
		return err
	}
	m.DB = session.DB(conf.Name)
	m.User.DB = m.DB.C("users")
	m.Log.DB = m.DB.C("logs")
	m.Article.DB = m.DB.C("articles")
	m.Website.DB = m.DB.C("website")
	m.Like.DB = m.DB.C("likes")
	log.Printf("MongoDB Connect Success!")
	return nil
}

func (m *Model) initWebsite() {
	err := m.Website.InitWebsiteInfo("冰淇淋", "https://xmatrix.studio/")
	if err != nil {
		return
	}
	website, err := m.Website.GetWebsiteInfo()
	if err != nil {
		return
	}
	generator.Generate("default", website.Name, website.URL)
}

func (m *Model) buildAllArticles() {
	website, err := m.Website.GetWebsiteInfo()
	if err != nil {
		return
	}
	for _, article := range m.Article.GetAllArticle() {
		user, err := m.User.GetUserByID(article.WriterID)
		if err != nil {
			fmt.Println("Loading fail: " + article.URL)
			continue
		}
		generator.GenerateArticle(website.Name, website.URL, article.Title, article.URL, article.Text, user.Name, article.PublishDate, article.Comment)
	}
}

func (m *Model) BuildAllPages() {
	articles := m.Article.GetPageArticle(100, 1)
	articlesInfo := []generator.ArticleInfo{}
	for _, article := range articles {
		user, err := m.User.GetUserByID(article.WriterID)
		if err != nil {
			continue
		}
		index := strings.Index(article.Text, "#")
		var text string
		if index != -1 && index <= 200 {
			text = article.Text[:index]
		} else if (index == -1 && len(article.Text) > 200) || index > 200 {
			text = article.Text[:200]
		} else {
			text = article.Text
		}
		articlesInfo = append(articlesInfo, generator.ArticleInfo{
			Title:      article.Title,
			URL:        article.URL,
			Date:       time.Unix(article.PublishDate/1000, 0).Format("2006-1-2"),
			WriterName: user.Name,
			Text:       text,
		})
	}
	website, err := m.Website.GetWebsiteInfo()
	if err != nil {
		return
	}
	for i := 1; (i-1)*10 <= website.Articles; i++ {
		if i*10 <= website.Articles {
			generator.GeneratePage(website.Name, website.URL, website.Articles, i, articlesInfo[(i-1)*10:i*10-1])
			generator.GenerateArchive(website.Name, website.URL, website.Articles, i, articlesInfo[(i-1)*10:i*10-1])
		} else {
			generator.GeneratePage(website.Name, website.URL, website.Articles, i, articlesInfo[(i-1)*10:])
			generator.GenerateArchive(website.Name, website.URL, website.Articles, i, articlesInfo[(i-1)*10:])
		}
	}
}

func (m *Model) RemoveArticle(url string) {
	generator.RemoveArticle(url)
}

// NewModel 创建新的Mongo连接
func NewModel(c Mongo) (*Model, error) {
	model := new(Model)
	err := model.initMongo(c)
	model.initWebsite()
	model.buildAllArticles()
	model.BuildAllPages()
	return model, err
}
