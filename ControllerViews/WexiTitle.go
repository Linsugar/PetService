package Views

import (
	"PetService/Models"
	"PetService/Untils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	c2 = make(chan interface{}, 2)
	c3 = make(chan string, 2)
)

type WeiXinArticle struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

var wc Models.T2

func getToken(c1 chan string) {
	appid := "wx50f04c5bde8f1938"
	secret := "784069c669fd121a564a836dae2f1d8b"
	we := WeiXinArticle{}
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%v&secret=%v", appid, secret)
	get, err := http.Get(url)
	if err != nil {
		return
	}
	body, _ := ioutil.ReadAll(get.Body)
	value := body
	errs := json.Unmarshal(value, &we)
	c1 <- we.AccessToken
	if errs != nil {
		return
	}
}

func GetArticle() {
	var c1 = make(chan string, 2)
	getToken(c1)
	arr := map[string]interface{}{
		"type": "news", "offset": 0, "count": 2,
	}
	url2 := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token=%v", <-c1)
	jsonStr, _ := json.Marshal(arr)
	post, err2 := http.Post(url2, "application/json", bytes.NewBuffer(jsonStr))
	if err2 != nil {
		fmt.Printf("当前错误：%v\n", err2)
		return
	}
	body2, _ := ioutil.ReadAll(post.Body)
	value2 := body2
	err4 := json.Unmarshal(value2, &wc)
	marshal, err := json.Marshal(&wc)
	if err != nil {
		return
	}
	js := string(marshal)
	c3 <- js
	fmt.Printf("\n当前json格式解析内容：%v\n,解析类型：%T\n", js, js)
	if err4 != nil {
		fmt.Printf("进入了这里5：%v\n", err4)
		return
	}
	fmt.Printf("\n当前解析：%v\n", wc)
	c2 <- wc
	defer close(c1)
}

func (WeiXinArticle) WeixinGet(c *gin.Context) {

	fmt.Println("V1===")
	redisResult := Untils.RedisDo{}.GetRedisValue("weixin")
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
		go GetArticle()
		Untils.RedisDo{}.SetRedisValue("weixin", <-c3, time.Minute*60*24)
		c.JSON(200, gin.H{
			"res": "sss",
			"re":  <-c2,
		})
	}
	//defer close(c2)
}

func WeixinPost(c *gin.Context) {

}
