package base

import (
	"errors"
	"ex/errorx"
	"github.com/gin-gonic/gin"
	"net/http"
)

type H func(c *gin.Context) (*Resp, error)

func Wrapper(h H) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := h(c)
		// 中间件error
		if err != nil && (errors.Is(err, errorx.ErrTokenMiss) || errors.Is(err, errorx.ErrTokenExpire)) {
			c.Abort()
		}
		// 中间件pass
		if resp == nil && err == nil {
			c.Next()
			return
		}
		// 错误补全
		if err != nil && resp.Message != "" {
			resp.Message = err.Error()
		}
		c.JSON(http.StatusOK, resp)
	}
}
