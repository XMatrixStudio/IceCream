package services

import (
	"github.com/XMatrixStudio/IceCream/httpserver/models"
)

// Service 服务
type Service struct {
	Model *models.Model
}

// NewService 创建服务
func NewService(m *models.Model) *Service {
	service := new(Service)
	service.Model = m
	return service
}

// NewUserService 创建用户服务
func (s *Service) NewUserService() UserService {
	return &userService{
		Model:   s.Model.User,
		Service: s,
	}
}

func (s *Service) NewArticleService() ArticleService {
	return &articleService{
		Model:   s.Model.Article,
		Service: s,
	}
}
