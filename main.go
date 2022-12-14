package main

import (
	"ex/repo"
	"ex/router"
	"github.com/gin-gonic/gin"
)

func main() {
	db := repo.ConnectMysql()

	ur := repo.User(db)

	g := gin.Default()

	router.ApiInit(g, ur)

	err := g.Run()
	if err != nil {
		return
	}
}
