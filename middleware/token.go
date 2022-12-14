package middleware

import (
	"ex/base"
	"ex/constant"
	"ex/errorx"
	"ex/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

func TokenHandler(c *gin.Context) (*base.Resp, error) {
	tokenStr := c.GetHeader("token")
	if tokenStr == "" {
		return base.ErrorResp(constant.TokenMiss), errorx.ErrTokenMiss
	}

	token, err := util.ParseToken(tokenStr)

	if err != nil {
		fmt.Printf("Token parse error: %s\n", err)
		return base.ErrorResp(constant.TokenExpire), errorx.ErrTokenExpire
	}

	c.Set("uid", token.Uid)

	return nil, nil
}
