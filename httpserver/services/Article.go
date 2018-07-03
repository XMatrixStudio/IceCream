package services

import (
	"errors"
	"time"

	"github.com/XMatrixStudio/IceCream/generator"
	"github.com/XMatrixStudio/IceCream/httpserver/models"
)

type ArticleService interface {
	GetArticleByURL(userID, url string) (article models.Article, err error)
	AddArticle(userID, name, url, text string, isComment bool) (err error)
	UpdateArticle(userID, name, oldurl, url, text string, isComment bool) (err error)
	RemoveArticle(userID, url string) (err error)
	GetLikeInfo(userID, url string) (likeNum int64, isLike bool, err error)
	LikeArticle(userID, url string, isLike bool) (err error)
}

type articleService struct {
	Model   models.ArticleModel
	Service *Service
}

func (s *articleService) GetArticleByURL(userID, url string) (article models.Article, err error) {
	article = *s.Model.GetArticleByURL(url)
	if article.WriterID != userID {
		err = errors.New("invalid_user")
		return
	}
	return
}

func (s *articleService) AddArticle(userID, title, url, text string, isComment bool) (err error) {
	user, err := s.Service.Model.User.GetUserByID(userID)
	if err != nil {
		return
	}
	if user.Level == 0 || user.Level == -1 {
		return errors.New("invalid_level")
	}
	article := s.Model.GetArticleByURL(url)
	if article != nil {
		return errors.New("duplicate_url")
	}
	objectID, err := s.Model.AddArticle(title, url, "article", userID, text, isComment)
	if err != nil {
		return
	}
	err = s.Service.Model.Log.AddLogRecord(userID, "创建文章"+objectID.Hex())
	if err != nil {
		return
	}
	err = s.Service.Model.Website.IncOrDecArticlesNum(1)
	if err != nil {
		return
	}
	website, err := s.Service.Model.Website.GetWebsiteInfo()
	if err != nil {
		return
	}
	generator.GenerateArticle(website.Name, website.URL, title, url, text, user.Name, time.Now().Unix()*1000, isComment)
	s.Service.Model.BuildAllPages()
	return
}

func (s *articleService) UpdateArticle(userID, title, oldurl, url, text string, isComment bool) (err error) {
	user, err := s.Service.Model.User.GetUserByID(userID)
	if err != nil {
		return
	}
	if user.Level == 0 || user.Level == -1 {
		return errors.New("invalid_level")
	}
	article := s.Model.GetArticleByURL(oldurl)
	if article.WriterID != userID {
		return errors.New("invalid_user")
	}
	if oldurl != url {
		if s.Model.GetArticleByURL(url) != nil {
			return errors.New("duplicate_url")
		}
	}
	err = s.Model.UpdateArticle(article.ID.Hex(), title, url, text, isComment)
	if err != nil {
		return
	}
	err = s.Service.Model.Log.AddLogRecord(userID, "修改文章"+article.ID.Hex())
	if err != nil {
		return
	}
	website, err := s.Service.Model.Website.GetWebsiteInfo()
	if err != nil {
		return
	}
	generator.GenerateArticle(website.Name, website.URL, title, url, text, user.Name, time.Now().Unix()*1000, isComment)
	return
}

func (s *articleService) RemoveArticle(userID, url string) (err error) {
	article := s.Model.GetArticleByURL(url)
	if article == nil {
		return errors.New("invalid_article")
	} else if article.WriterID != userID {
		return errors.New("invalid_user")
	}
	err = s.Model.RemoveArticle(article.ID.Hex())
	if err != nil {
		return
	}
	err = s.Service.Model.Log.AddLogRecord(userID, "删除文章"+article.ID.Hex())
	if err != nil {
		return
	}
	err = s.Service.Model.Website.IncOrDecArticlesNum(-1)
	s.Service.Model.RemoveArticle(url)
	s.Service.Model.BuildAllPages()
	return
}

func (s *articleService) GetLikeInfo(userID, url string) (likeNum int64, isLike bool, err error) {
	article := s.Model.GetArticleByURL(url)
	if article == nil {
		return 0, false, errors.New("invalid_article")
	}
	if userID != "" {
		user, err := s.Service.Model.User.GetUserByID(userID)
		if err != nil {
			return 0, false, err
		}
		isLike, err = s.Service.Model.Like.IsLike(user.ID, article.ID)
		if err != nil {
			return 0, false, err
		}
		return article.LikeNum, isLike, err
	}
	return article.LikeNum, false, err
}

func (s *articleService) LikeArticle(userID, url string, isLike bool) (err error) {
	user, err := s.Service.Model.User.GetUserByID(userID)
	if err != nil {
		return
	}
	if user.Level == -1 {
		return errors.New("black_list_user")
	}
	article := s.Model.GetArticleByURL(url)
	if article == nil {
		return errors.New("invalid_article")
	}
	if isLike {
		s.Service.Model.Like.AddArticle(userID, article.ID.Hex())
		err = s.Model.AddNum(article.ID.Hex(), "likeNum", 1)
	} else {
		s.Service.Model.Like.RemoveArticle(userID, article.ID.Hex())
		err = s.Model.AddNum(article.ID.Hex(), "likeNum", -1)
	}
	return
}
