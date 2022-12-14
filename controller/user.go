package controller

import (
	"ex/base"
	"ex/constant"
	"ex/repo"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	base.Controller
	Repo *repo.UserRepo
}

func (ctrl UserController) UserInfo(c *gin.Context) (*base.Resp, error) {
	uid := c.GetUint("uid")
	user, err := ctrl.Repo.FindById(uid)

	if err != nil {
		return &base.Resp{Code: constant.TokenErr, Message: "Fake token, User not exists"}, err
	}
	return &base.Resp{Code: 0, Message: "Success", Data: user}, nil
}
