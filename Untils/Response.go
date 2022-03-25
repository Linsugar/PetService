package Untils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseOkState() func() {
	return func() {
		fmt.Printf("")
	}
}

func ResponseBadState(c *gin.Context) {
	c.JSON(http.StatusGone, gin.H{})
}
