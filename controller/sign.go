package controller

import (
	"errors"
	"ex/base"
	"ex/constant"
	"ex/repo"
	"ex/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SignController struct {
	base.Controller
	Repo *repo.UserRepo
}

type SignReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (req *SignReq) Validate() (bool, error) {
	if req.UserName == "" {
		return false, errors.New("userName miss")
	}
	if req.Password == "" {
		return false, errors.New("password miss")
	}
	return true, nil
}

func (ctrl *SignController) SignUp(c *gin.Context) (*base.Resp, error) {
	req := SignReq{}
	err := ctrl.BindJSON(c, &req)
	if err != nil {
		return base.UnknownResp(), err
	}

	name := req.UserName
	pass := req.Password

	_, err1 := ctrl.Repo.FindByName(name)
	if err1 == nil {
		// 用户已存在
		return base.ErrorResp(constant.UserExists), errors.New("user exists")
	} else if !errors.Is(err1, gorm.ErrRecordNotFound) {
		// 其他错误
		return base.UnknownResp(), err1
	}

	_, err2 := ctrl.Repo.NewUser(name, pass)
	if err2 != nil {
		// 其他错误
		return base.UnknownResp(), err2
	}

	return base.SucResp(nil), nil
}

func (ctrl SignController) SignIn(c *gin.Context) (*base.Resp, error) {
	req := SignReq{}
	err := ctrl.BindJSON(c, &req)
	if err != nil {
		return base.UnknownResp(), err
	}

	name := req.UserName
	pass := req.Password

	if name == "" || pass == "" {

	}

	user, err1 := ctrl.Repo.FindByName(name)
	if err1 != nil {
		if errors.Is(err1, gorm.ErrRecordNotFound) {
			// 用户不存在
			return base.ErrorResp(constant.UserNotExists), errors.New("user not exists")
		}
		// 其他错误
		return base.UnknownResp(), nil
	}
	if user.Password != pass {
		// 密码错误
		return base.ErrorResp(constant.PassNotTrue), errors.New("password not correct")
	}

	token := util.GenToken(user.ID)
	ctrl.Repo.SaveToken(user, token)
	return base.SucResp(gin.H{
		"token": token,
	}), nil
}
