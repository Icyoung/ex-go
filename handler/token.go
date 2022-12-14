package handler

import (
	"ex/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TokenHandler(c *gin.Context) {
	tokenStr := c.GetHeader("token")
	if tokenStr == "" {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code":    1004,
			"message": "The request needs token in headers",
			"data":    nil,
		})
		return
	}

	token, err := util.ParseToken(tokenStr)

	if err != nil {
		fmt.Printf("Token parse error: %s\n", err)
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code":    1005,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.Set("uid", token.Uid)

	c.Next()
}
