package Tasks

import (
	"PetService/Conf"
	"PetService/Models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/robfig/cron"
	"io/ioutil"
	"net/http"
)

var (
	JsonArticle   = make(chan interface{}, 2)
	StringArticle = make(chan string, 2)
	ChanToken     = make(chan string, 2)
)

//添加定时任务
var Cr *cron.Cron

func init() {
	Cr = cron.New()
	Cr.Start()
	TaskInitAll()
}

func TaskInitAll() {
	//全部定时任务
	Cr.AddFunc("0 0 13 * * ?", GetArticle)

}

func getToken() {
	type WeiXinArticle struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	we := WeiXinArticle{}
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%v&secret=%v", Conf.Appid, Conf.Secret)
	get, err := http.Get(url)
	if err != nil {
		return
	}
	body, _ := ioutil.ReadAll(get.Body)
	value := body
	errs := json.Unmarshal(value, &we)
	ChanToken <- we.AccessToken
	if errs != nil {
		return
	}
}

func GetArticle() {
	var wc Models.T2
	getToken()
	arr := map[string]interface{}{
		"type": "news", "offset": 0, "count": 2,
	}
	url2 := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token=%v", <-ChanToken)
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
	StringArticle <- js
	//fmt.Printf("\n当前json格式解析内容：%v\n,解析类型：%T\n", js, js)
	if err4 != nil {
		fmt.Printf("进入了这里5：%v\n", err4)
		return
	}
	//fmt.Printf("\n当前解析：%v\n", wc)
	JsonArticle <- wc
}
