package controllers

import (
	"regexp"
	"strings"

	"github.com/XMatrixStudio/IceCream/httpserver/services"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

type ArticlesController struct {
	Ctx     iris.Context
	Service services.ArticleService
	Session *sessions.Session
}

type ArticleRes struct {
	State   string      `json:"state"`
	Msg     string      `json:"msg"`
	Article ArticleInfo `json:"article"`
}

type ArticleInfo struct {
	Title   string `json:"title"`
	OldURL  string `json:"oldurl"`
	URL     string `json:"url"`
	Comment bool   `json:"comment"`
	Text    string `json:"text"`
}

func (c *ArticlesController) Get() (res ArticleRes) {
	url := c.Ctx.FormValue("url")
	userID := c.Session.GetString("userID")
	if userID == "" {
		res.State = "error"
		res.Msg = "not_login"
		return
	}
	article, err := c.Service.GetArticleByURL(userID, url)
	if err != nil {
		res.State = "error"
		res.Msg = err.Error()
		return
	}
	res.State = "success"
	res.Article = ArticleInfo{
		Title:   article.Title,
		URL:     article.URL,
		Comment: article.Comment,
		Text:    article.Text,
	}
	return
}

var reservedPath = [...]string{
	"about",
	"archives",
	"editor",
	"page",
	"settings",
}

func (c *ArticlesController) Post() (res ArticleRes) {
	info := ArticleInfo{}
	c.Ctx.ReadJSON(&info)
	userID := c.Session.GetString("userID")
	if userID == "" {
		res.State = "error"
		res.Msg = "not_login"
		return
	}
	flag, err := regexp.MatchString(`^([A-Za-z0-9_-]+/{0,1})+$`, info.URL)
	if err != nil || !flag || info.Title == "" || info.Text == "" {
		res.State = "error"
		res.Msg = "invalid_params"
		return
	}
	basePath := strings.Split(info.URL, "/")[0]
	for _, str := range reservedPath {
		if str == basePath {
			res.State = "error"
			res.Msg = "reserved_url"
			return
		}
	}
	if info.URL[len(info.URL)-1] != '/' {
		info.URL += "/"
	}
	err = c.Service.AddArticle(userID, info.Title, info.URL, info.Text, info.Comment)
	if err != nil {
		res.State = "error"
		res.Msg = err.Error()
		return
	}
	res.State = "success"
	return
}

func (c *ArticlesController) Put() (res ArticleRes) {
	info := ArticleInfo{}
	c.Ctx.ReadJSON(&info)
	userID := c.Session.GetString("userID")
	if userID == "" {
		res.State = "error"
		res.Msg = "not_login"
		return
	}
	flag, err := regexp.MatchString(`^([A-Za-z0-9_-]+/{0,1})+$`, info.URL)
	if err != nil || !flag || info.Title == "" || info.Text == "" {
		res.State = "error"
		res.Msg = "invalid_params"
		return
	}
	basePath := strings.Split(info.URL, "/")[0]
	for _, str := range reservedPath {
		if str == basePath {
			res.State = "error"
			res.Msg = "reserved_url"
			return
		}
	}
	if info.URL[len(info.URL)-1] != '/' {
		info.URL += "/"
	}
	err = c.Service.UpdateArticle(userID, info.Title, info.OldURL, info.URL, info.Text, info.Comment)
	if err != nil {
		res.State = "error"
		res.Msg = err.Error()
		return
	}
	res.State = "success"
	return
}

func (c *ArticlesController) Delete() (res ArticleRes) {
	url := c.Ctx.FormValue("url")
	userID := c.Session.GetString("userID")
	if userID == "" {
		res.State = "error"
		res.Msg = "not_login"
		return
	}
	err := c.Service.RemoveArticle(userID, url)
	if err != nil {
		res.State = "error"
		res.Msg = err.Error()
	}
	res.State = "success"
	return
}

type ArticleCommentReq struct {
	URL    string `json:"url"`
	Text   string `json:"text"`
	Father string `json:"father"`
}

type ArticleCommentRes struct {
	State    string                        `json:"state"`
	Msg      string                        `json:"msg"`
	Comments []services.ArticleCommentInfo `json:"comments"`
}

func (c *ArticlesController) GetComments() (res ArticleCommentRes) {
	url := c.Ctx.FormValue("url")
	comments, err := c.Service.GetComments(url)
	if err != nil {
		res.State = "error"
		res.Msg = err.Error()
		return
	}
	res.State = "success"
	res.Comments = comments
	return
}

func (c *ArticlesController) PostComments() (res ArticleCommentRes) {
	comment := ArticleCommentReq{}
	c.Ctx.ReadJSON(&comment)
	userID := c.Session.GetString("userID")
	if userID == "" {
		res.State = "error"
		res.Msg = "not_login"
		return
	}
	if comment.URL == "" || comment.Text == "" {
		res.State = "error"
		res.Msg = "invalid_params"
		return
	}
	err := c.Service.AddComment(userID, comment.URL, comment.Text, comment.Father)
	if err != nil {
		res.State = "error"
		res.Msg = err.Error()
		return
	}
	res.State = "success"
	return
}

type ArticleLikeRes struct {
	State       string          `json:"state"`
	Msg         string          `json:"msg"`
	ArticleLike ArticleLikeInfo `json:"articleLike"`
}
type ArticleLikeInfo struct {
	Num  int64 `json:"num"`
	Like bool  `json:"like"`
}

func (c *ArticlesController) GetLike() (res ArticleLikeRes) {
	url := c.Ctx.FormValue("url")
	userID := c.Session.GetString("userID")
	likeNum, isLike, err := c.Service.GetLikeInfo(userID, url)
	if err != nil {
		res.State = "error"
		res.Msg = err.Error()
		return
	}
	res.State = "success"
	res.ArticleLike = ArticleLikeInfo{
		Num:  likeNum,
		Like: isLike,
	}
	return
}

type ArticleLikeReq struct {
	URL string `json:"url"`
}

func (c *ArticlesController) PostLike() (res ArticleLikeRes) {
	req := ArticleLikeReq{}
	c.Ctx.ReadJSON(&req)
	userID := c.Session.GetString("userID")
	if userID == "" {
		res.State = "error"
		res.Msg = "not_login"
		return
	}
	err := c.Service.LikeArticle(userID, req.URL, true)
	if err != nil {
		res.State = "error"
		res.Msg = err.Error()
		return
	}
	res.State = "success"
	return
}

func (c *ArticlesController) DeleteLike() (res ArticleLikeRes) {
	url := c.Ctx.FormValue("url")
	userID := c.Session.GetString("userID")
	if userID == "" {
		res.State = "error"
		res.Msg = "not_login"
		return
	}
	err := c.Service.LikeArticle(userID, url, false)
	if err != nil {
		res.State = "error"
		res.Msg = err.Error()
		return
	}
	res.State = "success"
	return
}
