package models

import (
	"fmt"
	"log"
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
	log.Printf("MongoDB Connect Success!")
	return nil
}

func (m *Model) buildAllArticles() {
	for _, article := range m.Article.GetAllArticle() {
		user, err := m.User.GetUserByID(article.WriterID)
		if err != nil {
			fmt.Println("Loading fail: " + article.URL)
			continue
		}
		generator.GenerateArticle(article.Title, article.URL, article.Text, user.Name, article.PublishDate, article.Comment)
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
		articlesInfo = append(articlesInfo, generator.ArticleInfo{
			Title:      article.Title,
			URL:        article.URL,
			Date:       time.Unix(article.PublishDate/1000, 0).Format("2006-1-2"),
			WriterName: user.Name,
			Text:       article.Text,
		})
	}
	generator.GeneratePage(1, articlesInfo)
}

// NewModel 创建新的Mongo连接
func NewModel(c Mongo) (*Model, error) {
	model := new(Model)
	err := model.initMongo(c)
	model.buildAllArticles()
	model.BuildAllPages()
	return model, err
}
