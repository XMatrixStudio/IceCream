package controllers

import (
	"regexp"

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
