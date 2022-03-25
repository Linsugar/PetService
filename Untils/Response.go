package Untils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseOkState(c *gin.Context, Data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "请求成功",
		"Result":  Data,
	})
}

func ResponseBadState(c *gin.Context, err error) {
	c.JSON(http.StatusGone, gin.H{
		"Message": "请求有误,请检查",
		"Error":   err.Error(),
	})
}
