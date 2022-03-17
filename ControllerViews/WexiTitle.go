package Views

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type WeiXinArticle struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
type T2 struct {
	TotalCount int `json:"total_count"`
	ItemCount  int `json:"item_count"`
	Item       []struct {
		MediaId string `json:"media_id"`
		Content struct {
			NewsItem []struct {
				Title            string `json:"title"`
				ThumbMediaId     string `json:"thumb_media_id"`
				ShowCoverPic     int    `json:"show_cover_pic"`
				Author           string `json:"author"`
				Digest           string `json:"digest"`
				Content          string `json:"content"`
				Url              string `json:"url"`
				ContentSourceUrl string `json:"content_source_url"`
				Thumb_url        string `json:"thumb_url"`
			} `json:"news_item"`
		} `json:"content"`
		UpdateTime int `json:"update_time"`
	} `json:"item"`
}

var we WeiXinArticle
var wc T2
var c1 = make(chan string, 10)
var c2 = make(chan string, 10)

func getToken() {
	appid := "wx50f04c5bde8f1938"
	secret := "784069c669fd121a564a836dae2f1d8b"
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%v&secret=%v", appid, secret)

	get, err := http.Get(url)
	if err != nil {
		return
	}
	body, _ := ioutil.ReadAll(get.Body)
	value := []byte(body)
	errs := json.Unmarshal(value, &we)
	c1 <- we.AccessToken
	if errs != nil {
		return
	}
}

func getArticle() {
	arr := map[string]interface{}{
		"type": "news", "offset": 0, "count": 1,
	}
	url2 := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token=%v", <-c1)
	jsonStr, _ := json.Marshal(arr)
	post, err2 := http.Post(url2, "application/json", bytes.NewBuffer(jsonStr))
	if err2 != nil {
		fmt.Printf("当前错误：%v", err2)
		return
	}

	body2, _ := ioutil.ReadAll(post.Body)
	fmt.Printf("pOst 返回的数据：%v\n", string(body2))

	value2 := []byte(body2)
	err4 := json.Unmarshal(value2, &wc)
	if err4 != nil {
		fmt.Printf("进入了这里5：%v", err4)
		return
	}

	fmt.Printf("当前解析：%v", wc.Item[0].Content.NewsItem[0].Thumb_url)
	c2 <- wc.Item[0].Content.NewsItem[0].Thumb_url
	defer close(c1)
}

func (WeiXinArticle) WeixinGet(c *gin.Context) {
	go getToken()
	go getArticle()
	c.JSON(200, gin.H{
		"res": "sss",
		"re":  <-c2,
	})
	defer close(c2)
}

func WeixinPost(c *gin.Context) {

}
