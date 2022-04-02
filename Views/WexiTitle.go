package Views

import (
	"PetService/Models"
	"PetService/Tasks"
	"PetService/Untils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
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
		Untils.ResponseOkState(c, &wc)
	} else {
		Tasks.GetArticle()
		Untils.ResponseOkState(c, <-Tasks.JsonArticle)
	}
	//defer close(c2)
}

func WeixinPost(c *gin.Context) {

}
