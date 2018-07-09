package controllers

import (
	"regexp"

	"github.com/XMatrixStudio/IceCream/httpserver/services"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

type WebsiteController struct {
	Ctx     iris.Context
	Service services.WebsiteService
	Session *sessions.Session
}

type WebsiteReq struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Text string `json:"text"`
}

type WebsiteRes struct {
	State string `json:"state"`
	Msg   string `json:"msg"`
}

func (c *WebsiteController) Put() (res WebsiteRes) {
	req := WebsiteReq{}
	c.Ctx.ReadJSON(&req)
	userID := c.Session.GetString("userID")
	if userID == "" {
		res.State = "error"
		res.Msg = "not_login"
		return
	}
	flag, err := regexp.MatchString(`^http[s]{0,1}:\/\/([A-Za-z0-9_.-]+\/)+$`, req.URL)
	if err != nil || !flag || req.Name == "" || req.URL == "" {
		res.State = "error"
		res.Msg = "invalid_params"
		return
	}
	err = c.Service.UpdateWebsiteInfo(userID, req.Name, req.URL, req.Text)
	if err != nil {
		res.State = "error"
		res.Msg = err.Error()
		return
	}
	res.State = "success"
	return
}
