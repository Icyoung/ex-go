package base

import (
	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (ctrl *Controller) BindJSON(c *gin.Context, req Req) error {
	if ok, err := req.Validate(); !ok {
		return err
	}
	return c.BindJSON(req)
}
