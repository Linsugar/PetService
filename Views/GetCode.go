package Views

import (
	"PetService/Conf"
	"PetService/Untils"
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//获取验证码

func CodeController(c *gin.Context) {
	if c.Request.Method == "GET" {
		QQCode(c)
	} else if c.Request.Method == "POST" {
		PostCode(c)
	}
}

func QQCode(c *gin.Context) {
	user := c.Query("phone")
	strlist := strings.Split(user, "@")
	if strlist[1] == "qq.com" {
		fmt.Println("获取到当前的参数：", user)
		devices := c.Query("devices")
		cip := c.ClientIP()
		err := Untils.SendEmailQQ(user, devices, cip)
		if err != nil {
			Untils.ResponseBadState(c, err)
			return
		}
		Untils.ResponseOkState(c, "发送成功")
	} else {
		Untils.ResponseBadState(c, errors.New("您的邮箱不合法，只支持QQ邮箱"))
	}
}

func PostCode(c *gin.Context) {

	currentTime := time.Now().Unix()
	str := fmt.Sprintf("%v%v%v", Conf.AppSecret, Conf.Nonce, currentTime)
	url := "https://api.netease.im/sms/sendcode.action"
	s1 := sha1.New()
	n, _ := s1.Write([]byte(str))
	result := s1.Sum([]byte(str))
	vale := string(result)
	arr := map[string]interface{}{
		"AppKey":   Conf.AppKey,
		"Nonce":    Conf.Nonce,
		"CurTime":  currentTime,
		"CheckSum": vale,
	}
	jsonStr, _ := json.Marshal(arr)
	post, err2 := http.Post(url, "application/x-www-form-urlencoded;charset=utf-8", bytes.NewBuffer(jsonStr))
	if err2 != nil {
		fmt.Println("当前错误=======>>>>>", err2)
	}
	body2, _ := ioutil.ReadAll(post.Body)

	fmt.Println("当前转换：", n)
	fmt.Println("当前转换：", string(body2))
	Untils.ResponseOkState(c, result)

}
