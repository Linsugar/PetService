package Tasks

import (
	"PetService/Conf"
	"PetService/Models"
	"PetService/Untils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/robfig/cron"
	"io/ioutil"
	"net/http"
	"time"
)

//添加定时任务
var Cr *cron.Cron
var JsonArticle = make(chan interface{}, 2)

func init() {
	Cr = cron.New()
	//这里会出现循环掉包的问题-已经
	Cr.Start()
	TaskInitAll()
}

func TaskInitAll() {
	//全部定时任务
	Cr.AddFunc("0 0 13 * * ?", GetArticle)

}

func getToken() chan string {
	ChanToken := make(chan string, 2)
	type WeiXinArticle struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	we := WeiXinArticle{}
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%v&secret=%v", Conf.Appid, Conf.Secret)
	get, err := http.Get(url)
	if err != nil {
		return ChanToken
	}
	body, _ := ioutil.ReadAll(get.Body)
	value := body
	errs := json.Unmarshal(value, &we)

	if errs != nil {
		return ChanToken
	}
	ChanToken <- we.AccessToken
	return ChanToken
}

func GetArticle() {
	StringArticle := make(chan string, 2)
	var wc Models.T2
	token := getToken()
	arr := map[string]interface{}{
		"type": "news", "offset": 0, "count": 10,
	}
	url2 := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token=%v", <-token)
	jsonStr, _ := json.Marshal(arr)
	post, err2 := http.Post(url2, "application/json", bytes.NewBuffer(jsonStr))
	if err2 != nil {
		fmt.Printf("当前错误：%v\n", err2)
		return
	}
	body2, _ := ioutil.ReadAll(post.Body)
	err4 := json.Unmarshal(body2, &wc)
	marshal, err := json.Marshal(&wc)
	if err4 != nil {
		fmt.Printf("进入了这里5：%v\n", err4)
		return
	}
	if err != nil {
		return
	}
	JsonArticle <- wc
	js := string(marshal)
	StringArticle <- js
	Untils.SetRedisValue("weixin", <-StringArticle, time.Second*1000)

}
