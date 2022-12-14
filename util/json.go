package util

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func Json2Map(c *gin.Context) (map[string]any, error) {
	var m map[string]any
	data, e := c.GetRawData()
	e = json.Unmarshal(data, &m)
	return m, e
}
