package controllers

import (
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
	State string
	Msg   string
	Data  ArticleInfo
}

type ArticleInfo struct {
	Name    string
	URL     string
	Comment bool
	Text    string
}

func (c *ArticlesController) Get() (res ArticleRes) {
	id := c.Ctx.FormValue("id")
	userID := c.Session.GetString("userID")
	if userID != "" {
		res.State = "error"
		res.Msg = "not_login"
		return
	}
	article := c.Service.GetArticleById(id, userID)
	res.State = "success"
	res.Data = ArticleInfo{
		Name:    article.Name,
		URL:     article.URL,
		Comment: article.Comment,
		Text:    article.Text,
	}
	return
}

func (c *ArticlesController) Post() (res ArticleRes) {
	info := ArticleInfo{}
	c.Ctx.ReadForm(&info)
	userID := c.Session.GetString("userID")
	if userID != "" {
		res.State = "error"
		res.Msg = "not_login"
		return
	}
	err := c.Service.AddArticle(userID, info.Name, info.URL, info.Text, info.Comment)
	if err != nil {
		res.State = "error"
		res.Msg = "invalid_level_or_url"
	}
	c.Ctx.Redirect(c.Ctx.Host() + info.URL)
	return
}
