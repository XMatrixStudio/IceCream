package services

import (
	"errors"

	"github.com/XMatrixStudio/IceCream/httpserver/models"
)

type WebsiteService interface {
	UpdateWebsiteInfo(userID, name, url, text string) (err error)
}

type websiteService struct {
	Model   models.WebsiteModel
	Service *Service
}

func (s *websiteService) UpdateWebsiteInfo(userID, name, url, text string) (err error) {
	user, err := s.Service.Model.User.GetUserByID(userID)
	if err != nil {
		return
	}
	if user.Level != 99 {
		return errors.New("invalid_level")
	}
	s.Model.EditWebsiteInfo(name, url, text)
	s.Service.Model.InitWebsite()
	s.Service.Model.BuildAllArticles()
	s.Service.Model.BuildAllPages()
	return
}
