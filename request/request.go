package request

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ParseJson(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		logrus.WithError(err).Error("解析失败")
		return err
	}
	return nil
}
