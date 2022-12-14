package router

import (
	"ex/base"
	"ex/controller"
	"ex/middleware"
	"ex/repo"
	"github.com/gin-gonic/gin"
)

func ApiInit(g *gin.Engine, r *repo.UserRepo) {
	apiRouter := g.Group("/api")
	apiCtrl := controller.SignController{Repo: r}
	{
		apiRouter.POST("/sign_up", base.Wrapper(apiCtrl.SignUp))
		apiRouter.POST("/sign_in", base.Wrapper(apiCtrl.SignIn))
	}

	userApiInit(g, r)
}

func userApiInit(g *gin.Engine, r *repo.UserRepo) {
	userApiRouter := g.Group("/api/user", base.Wrapper(middleware.TokenHandler))
	userApiCtrl := controller.UserController{Repo: r}
	{
		userApiRouter.GET("/user_info", base.Wrapper(userApiCtrl.UserInfo))
	}
}
