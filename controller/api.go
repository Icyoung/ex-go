package controller

import (
	"errors"
	"ex/repo"
	"ex/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type ApiController struct {
	Repo *repo.UserRepo
}

func (ctrl ApiController) SignUp(c *gin.Context) {
	m, _ := util.Json2Map(c)

	name := m["user_name"].(string)
	pass := m["password"].(string)

	_, err := ctrl.Repo.FindByName(name)

	if err == nil {
		// 用户已存在
		c.JSON(http.StatusOK, gin.H{
			"code":    1001,
			"message": "User exists",
			"data":    nil,
		})
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 其他错误
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err,
			"data":    nil,
		})
		return
	}

	_, err1 := ctrl.Repo.NewUser(name, pass)

	if err1 != nil {
		// 其他错误
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "sign up success",
		"data":    nil,
	})
	return
}

func (ctrl ApiController) SignIn(c *gin.Context) {
	m, _ := util.Json2Map(c)

	name := m["user_name"].(string)
	pass := m["password"].(string)

	user, err := ctrl.Repo.FindByName(name)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 用户不存在
			c.JSON(http.StatusOK, gin.H{
				"code":    1002,
				"message": "User not exists",
				"data":    nil,
			})
		} else {
			// 其他错误
			c.JSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": err,
				"data":    nil,
			})
		}
		return
	}

	if user.Password != pass {
		// 密码错误
		c.JSON(http.StatusOK, gin.H{
			"code":    1003,
			"message": "Password not correct",
			"data":    nil,
		})
		return
	}

	token := util.GenToken(user.ID)
	ctrl.Repo.SaveToken(user, token)

	data := map[string]any{}
	data["token"] = token
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "Login success",
		"data":    data,
	})
	return
}
