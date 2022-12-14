package controller

import (
	"ex/repo"
	"ex/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	Repo *repo.UserRepo
}

func (ctrl UserController) UserInfo(c *gin.Context) {
	uid := c.GetUint("uid")

	user, err := ctrl.Repo.FindById(uid)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1006,
			"message": "Fake token, User not exists",
			"data":    nil,
		})
		return
	}

	userJson, _ := util.Obj2Json(user)
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "Success",
		"data":    userJson,
	})
	return
}
