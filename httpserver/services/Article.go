package services

import (
	"errors"
	"os"

	"github.com/XMatrixStudio/IceCream/httpserver/models"
)

type ArticleService interface {
	GetArticleById(id, userID string) (article models.Article)
	AddArticle(userID, name, url, text string, isComment bool) (err error)
}

type articleService struct {
	Model   models.ArticleModel
	Service *Service
}

func (s *articleService) GetArticleById(id, userID string) (article models.Article) {
	article = *s.Model.GetArticleByID(id)
	if article.WriterID != userID {
		article = models.Article{}
		return
	}
	return
}

func (s *articleService) AddArticle(userID, name, url, text string, isComment bool) (err error) {
	user, err := s.Service.Model.User.GetUserByID(userID)
	if err != nil {
		return
	}
	if user.Level == 0 {
		err = errors.New("invalid_level")
		return
	}
	file, err := os.Open("_post/" + url)
	if err != nil {
		return
	}
	_, err = file.WriteString(text)
	if err != nil {
		return
	}
	_, err = s.Model.AddArticle(name, url, "article", userID, text, isComment)
	return
}
