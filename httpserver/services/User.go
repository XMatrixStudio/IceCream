package services

import (
	"github.com/XMatrixStudio/IceCream/httpserver/models"
	"github.com/XMatrixStudio/Violet.SDK.Go"
)

type UserService interface {
	InitViolet(c violetSdk.Config)
	// 登陆部分API
	GetLoginURL(redirectURL string) (url, state string)
	InitUserOrNot(code string) (user models.User, err error)
	Logout(userID string) (err error)
	//Login(name, password string) (valid bool, email string, err error)
	//GetUser(code string) (ID, name string, err error)
	//Register(name, email, password string) (err error)
	//GetEmailCode(email string) error
	//ValidEmail(email, vCode string) error
	//GetUserInfo(id string) (user models.Users, err error)
	//SetUserName(id, name string) error
}

type userService struct {
	Model   models.UserModel
	Service *Service
	Violet  violetSdk.Violet
}

type LoginRes struct {
	Valid bool
	Email string
	Code  string
}

func (s *userService) GetLoginURL(redirectURL string) (url, state string) {
	url, state = s.Violet.GetLoginURL(redirectURL)
	return
}

func (s *userService) Logout(userID string) (err error) {
	return s.Service.Model.Log.AddLogRecord(userID, "注销")
}

func (s *userService) InitUserOrNot(code string) (user models.User, err error) {
	tRes, err := s.Violet.GetToken(code)
	if err != nil {
		return
	}
	token, userVID := tRes.Token, tRes.UserID
	uRes, err := s.Violet.GetUserBaseInfo(userVID, token)
	if err != nil {
		return
	}
	user, err = s.Model.GetUserByVID(userVID)
	if err != nil {
		_, err = s.Model.AddUser(userVID, uRes.Name, uRes.Email, token, uRes.Info.Avatar, uRes.Info.Gender)
		if err != nil {
			return
		}
		user, err = s.Model.GetUserByVID(userVID)
		if err != nil {
			return
		}
		err = s.Service.Model.Log.AddLogDocument(user.ID.Hex())
		if err != nil {
			return
		}
		err = s.Service.Model.Log.AddLogRecord(user.ID.Hex(), "注册")
		if err != nil {
			return
		}
	} else {
		if user.Token != token {
			err = s.Model.SetUserToken(user.ID.Hex(), token)
		}
		err = s.Service.Model.Log.AddLogRecord(user.ID.Hex(), "登录")
	}
	return
}

/*func (s *userService) Login(name, password string) (valid bool, data string, err error) {
	resp, tErr := s.Violet.Login(name, password)
	if tErr != nil {
		err = tErr
		return
	}
	// 非正常的返回码
	if resp.StatusCode() != 200 {
		err = errors.New(resp.String())
		return
	}
	// 解析结果
	var loginRes LoginRes
	err = json.Unmarshal([]byte(resp.String()), &loginRes)
	if err != nil {
		return
	}
	valid = loginRes.Valid
	// 未激活邮箱
	if !valid {
		data = loginRes.Email
		return
	}
	// 登陆成功
	valid = true
	data = loginRes.Code
	return
}

type TokenRes struct {
	UserID string
	Token  string
}

func (s *userService) GetUser(code string) (ID, name string, err error) {
	// 获取用户Token
	return
}

type UserInfoRes struct {
	Email string
	Name  string
	Info  UserInfo
}

type UserInfo struct {
	Avatar string
	Gender int
}

func (s *userService) SaveUser(userVID, token string) (ID string, err error) {
	return
}

func (s *userService) Register(name, email, password string) error {
	resp, err := s.Violet.Register(name, email, password)
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return errors.New(resp.String())
	}
	return nil
}

func (s *userService) GetEmailCode(email string) error {
	resp, err := s.Violet.GetEmailCode(email)
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return errors.New(resp.String())
	}
	return nil
}

func (s *userService) ValidEmail(email, vCode string) error {
	resp, err := s.Violet.ValidEmail(email, vCode)
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return errors.New(resp.String())
	}
	return nil
}

func (s *userService) GetUserInfo(id string) (user models.Users, err error) {
	user, err = s.Model.GetUserByID(id)
	return
}

func (s *userService) SetUserName(id, name string) error {
	err := s.Model.SetUserName(id, name)
	return err
}*/

func (s *userService) InitViolet(c violetSdk.Config) {
	s.Violet = violetSdk.NewViolet(c)
}
