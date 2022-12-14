package router

import (
	"ex/controller"
	"ex/handler"
	"ex/repo"
	"github.com/gin-gonic/gin"
)

func ApiInit(g *gin.Engine, r *repo.UserRepo) {
	apiRouter := g.Group("/api")
	apiCtrl := controller.ApiController{Repo: r}
	{
		apiRouter.POST("/sign_up", apiCtrl.SignUp)
		apiRouter.POST("/sign_in", apiCtrl.SignIn)
	}

	userApiInit(g, r)
}

func userApiInit(g *gin.Engine, r *repo.UserRepo) {
	userApiRouter := g.Group("/api/user", handler.TokenHandler)
	userApiCtrl := controller.UserController{Repo: r}
	{
		userApiRouter.GET("/user_info", userApiCtrl.UserInfo)
	}
}
