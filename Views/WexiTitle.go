package Views

import (
	"PetService/Models"
	"PetService/Tasks"
	"PetService/Untils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

var wc Models.T2

func WeixinGet(c *gin.Context) {
	redisResult := Untils.GetRedisValue("weixin")
	if redisResult != nil {
		str := fmt.Sprintf("%v", redisResult)
		err := json.Unmarshal([]byte(str), &wc)
		if err != nil {
			return
		}
		c.JSON(200, gin.H{
			"res": "从redis获取到数据",
			"re":  &wc,
		})
	} else {
		Tasks.GetArticle()
		//fmt.Printf("当前返回的值1：%v\n", <-Tasks.JsonArticle)
		Untils.SetRedisValue("weixin", <-Tasks.StringArticle, time.Second*100)
		//fmt.Printf("当前返回的值2：%v\n", <-Tasks.JsonArticle)
		c.JSON(200, gin.H{
			"res": "sss",
			"re":  <-Tasks.JsonArticle,
		})
	}
	//defer close(c2)
}

func WeixinPost(c *gin.Context) {

}
