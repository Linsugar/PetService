package Untils

import (
	"PetService/Conf"
	"PetService/Models"
	"bytes"
	"fmt"
	"gopkg.in/gomail.v2"
	"math/rand"
	"strconv"
)

func SendEmailQQ(user string, devices string, ip string) error {
	//发送邮件
	LenCode := make([]int, 6)
	var le = []rune("1234567890")
	for i := 0; i < 6; i++ {
		initValue := le[rand.Intn(len(le))]
		value := string(initValue)
		LenCode[i], _ = strconv.Atoi(value)
	}
	var bt bytes.Buffer
	for _, i2 := range LenCode {
		bt.WriteString(strconv.Itoa(i2))
	}
	go CodeDataSql(bt.String(), devices, ip)
	fmt.Println("进入SendEmailQQ")
	m := gomail.NewMessage()
	m.SetHeader("From", "1753215994@qq.com")
	m.SetHeader("To", user)
	m.SetAddressHeader("Cc", "1753215994@qq.com", "Dan")
	m.SetHeader("Subject", "宠爱，您的验证码，请注意查收!")
	str := fmt.Sprintf("这是您的验证码： <b>%v</b> 请不要告诉他人，不然拉黑你!", LenCode)
	m.SetBody("text/html", str)
	//m.Attach("/home/Alex/lolcat.jpg") //传入图片或文件
	d := gomail.NewDialer(Conf.QQSmtp, 587, Conf.QQUser, Conf.QQPwd)
	// Send the email to Bob, Cora and Dan.
	err := d.DialAndSend(m)
	return err
}

func CodeDataSql(code string, codeDevice string, codeIp string) {
	value := Models.RegisterCode{
		CodeIp:     codeIp,
		CodeDevice: codeDevice,
		Code:       code,
	}
	Db.Model(Models.RegisterCode{}).Create(&value)
}
