package Untils

import (
	"PetService/Conf"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func QiNiuToken() string {
	bucket := "tanghuadong"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200 //示例2小时有效期
	mac := qbox.NewMac(Conf.AccessKey, Conf.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	//SetRedisValue("qiniu",upToken,3600);
	return upToken
}
