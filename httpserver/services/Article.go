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
	GetComments(url string) (infos []ArticleCommentInfo, err error)
	AddComment(userID, url, text, father string) (err error)
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
	s.Service.Model.BuildAllPages()
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

type ArticleCommentInfo struct {
	ID         string `json:"id"`
	User       string `json:"user"`
	Avatar     string `json:"avatar"`
	Text       string `json:"text"`
	Date       int64  `json:"date"`
	FatherID   string `json:"fatherId"`
	FatherUser string `json:"fatherUser"`
}

func (s *articleService) GetComments(url string) (infos []ArticleCommentInfo, err error) {
	article := s.Model.GetArticleByURL(url)
	if article == nil {
		return nil, errors.New("invalid_article")
	}
	comments, err := s.Service.Model.Comment.GetCommentByArticleID(article.ID.Hex())
	if err != nil {
		return nil, err
	}
	for _, comment := range comments {
		user, err := s.Service.Model.User.GetUserByID(comment.UserID.Hex())
		if err != nil {
			return nil, err
		}
		var fatherUser string
		if comment.Father != "" {
			fatherComment, err := s.Service.Model.Comment.GetCommentByID(comment.Father)
			if err != nil {
				return nil, err
			}
			father, err := s.Service.Model.User.GetUserByID(fatherComment.UserID.Hex())
			if err != nil {
				return nil, err
			}
			fatherUser = father.Name
		}
		infos = append(infos, ArticleCommentInfo{
			ID:         comment.ID.Hex(),
			User:       user.Name,
			Avatar:     user.Info.Avatar,
			Text:       comment.Content,
			Date:       comment.Date,
			FatherID:   comment.Father,
			FatherUser: fatherUser,
		})
	}
	return
}

func (s *articleService) AddComment(userID, url, text, father string) (err error) {
	user, err := s.Service.Model.User.GetUserByID(userID)
	if err != nil {
		return
	}
	if user.Level == -1 {
		return errors.New("invalid_level")
	}
	article := s.Model.GetArticleByURL(url)
	if article == nil {
		return errors.New("invalid_article")
	}
	if article.Comment == false {
		return errors.New("disabled_comment")
	}
	if father != "" {
		comment, err := s.Service.Model.Comment.GetCommentByID(father)
		if err != nil {
			return err
		}
		if comment == nil {
			return errors.New("invalid_father")
		}
		err = s.Service.Model.Log.AddLogRecord(userID, "评论文章"+url+"中的评论")
		if err != nil {
			return err
		}
	} else {
		err = s.Service.Model.Log.AddLogRecord(userID, "评论文章"+url)
		if err != nil {
			return
		}
	}
	err = s.Service.Model.Comment.AddComment(article.ID.Hex(), userID, father, text)
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
