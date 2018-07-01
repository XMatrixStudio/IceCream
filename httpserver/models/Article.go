package models

import (
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type ArticleModel struct {
	DB *mgo.Collection
}

// Article 内容
type Article struct {
	ID          bson.ObjectId `bson:"_id"`
	Title       string        `bson:"title"`       // 文章名字
	URL         string        `bson:"url"`         // 文章内容文件的url地址
	WriterID    string        `bson:"writerId"`    // 作者ID
	PublishDate int64         `bson:"publishDate"` // 发布日期
	EditDate    int64         `bson:"editDate"`    // 修改日期
	Text        string        `bson:"text"`        // 文本内容
	LikeNum     int64         `bson:"likeNum"`     // 点赞人数
	CommentNum  int64         `bson:"commentNum"`  // 评论次数
	ReadNum     int64         `bson:"readNum"`     // 阅读次数
	Top         bool          `bson:"top"`         // 是否置顶
	Lock        bool          `bson:"lock"`        // 是否锁定
	Comment     bool          `bson:"comment"`     // 是否开放评论
	Type        string        `bson:"type"`        // 类型， 页面还是文章 'page' - 页面 'article' - 文章
}

// AddArticle 增加内容
func (m *ArticleModel) AddArticle(title, url, ArticleType, userID, text string, isComment bool) (bson.ObjectId, error) {
	newArticle := bson.NewObjectId()
	err := m.DB.Insert(&Article{
		ID:          newArticle,
		Title:       title,
		URL:         url,
		WriterID:    userID,
		Comment:     isComment,
		Text:        text,
		Type:        ArticleType,
		PublishDate: time.Now().Unix() * 1000,
		EditDate:    time.Now().Unix() * 1000,
	})
	if err != nil {
		return "", err
	}
	return newArticle, nil
}

func (m *ArticleModel) UpdateArticle(id, title, url, text string, isComment bool) (err error) {
	err = m.DB.UpdateId(bson.ObjectIdHex(id), bson.M{
		"$set": bson.M{
			"title":   title,
			"url":     url,
			"text":    text,
			"comment": isComment,
		},
	})
	return
}

// EditArticle 更新修改日期
func (m *ArticleModel) EditArticle(id string) error {
	err := m.DB.UpdateId(bson.ObjectIdHex(id), bson.M{"$set": bson.M{"editDate": time.Now().Unix() * 1000}})
	return err
}

// AddNum 增加一个或减少一个阅读("readNum")/点赞("likeNum")/评论数("commentNum")
func (m *ArticleModel) AddNum(id, name string, num int) error {
	err := m.DB.UpdateId(bson.ObjectIdHex(id), bson.M{"$inc": bson.M{name: num}})
	return err
}

// SetStatus 设置置顶("top")/评论("comment")/锁定("lock")状态
func (m *ArticleModel) SetStatus(id, name string, status bool) error {
	err := m.DB.UpdateId(bson.ObjectIdHex(id), bson.M{"$set": bson.M{name: status}})
	return err
}

// RemoveArticle 删除内容
func (m *ArticleModel) RemoveArticle(id string) error {
	err := m.DB.RemoveId(bson.ObjectIdHex(id))
	return err
}

// GetArticleByID 根据ID查询内容
func (m *ArticleModel) GetArticleByID(id string) *Article {
	Article := new(Article)
	err := m.DB.FindId(id).One(&Article)
	if err != nil {
		return nil
	}
	return Article
}

// GetArticleByWriter 根据作者ID查询内容
func (m *ArticleModel) GetArticleByWriter(id string) []Article {
	var Article []Article
	err := m.DB.Find(bson.M{"writerId": id}).All(&Article)
	if err != nil {
		return nil
	}
	return Article
}

func (m *ArticleModel) GetAllArticle() []Article {
	var article []Article
	err := m.DB.Find(nil).All(&article)
	if err != nil {
		return nil
	}
	return article
}

func (m *ArticleModel) GetArticleByURL(url string) *Article {
	article := new(Article)
	err := m.DB.Find(bson.M{"url": url}).One(&article)
	if err != nil {
		return nil
	}
	return article
}

// GetPageArticle 获取内容指定分页内容集合
func (m *ArticleModel) GetPageArticle(eachNum, pageNum int) []Article {
	var Article []Article
	err := m.DB.Find(nil).Sort("-editDate").Skip(eachNum * (pageNum - 1)).Limit(eachNum).All(&Article)
	if err != nil {
		return nil
	}
	return Article
}
