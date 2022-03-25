package Views

import (
	"PetService/Middlewares"
	"PetService/Models"
	"PetService/Untils"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type UserController struct {
}

type RegisterController struct {
}

// PingExample godoc
// @Summary 用户登录
// @Schemes
// @Description 专门用户获取所用用户列表的
// @Tags 获取用户列表
// @Accept x-www-form-urlencoded
// @Produce x-www-form-urlencoded
// @Success 200 {string} test1
// @Router /user [get]
func UserGet(c *gin.Context) {
	//获取所有的用户列表
	NowIp := c.ClientIP()
	fmt.Printf("得到的访问ip：%v", NowIp)
	var us []Models.User
	Untils.Db.Model(&Models.User{}).Find(&us)
	c.JSON(200, gin.H{
		"msg":    "返回成功",
		"result": us,
	})
}

type Login struct {
	//必须大写
	//form:"username" json:"user" uri:"user" xml:"user" binding:"required"
	Phone    string `form:"phone" json:"phone" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	//Token    string `form:"token" json:"token" binding:"required"`
	Nowip string `form:"nowip" json:"nowip"`
}

// PingExample godoc
// @Summary 用户登录
// @Schemes
// @Description 对于用户进行登录
// @Tags 登录
// @Accept x-www-form-urlencoded
// @Produce x-www-form-urlencoded
// @Param phone formData  string true "用户手机号"
// @Param password formData   string true "用户密码"
// @Success 200 {string} test2
// @Router /user [post]
func UserPost(c *gin.Context) {
	//用户登录
	var formData Login
	var Data Models.User
	err := c.Bind(&formData)
	formData.Nowip = c.ClientIP()
	if err != nil {
		c.JSON(400, gin.H{
			"msg": formData,
			"err": err.Error(),
		})
		return
	}
	err3 := Untils.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Models.User{}).Where("phone =? and password=?", formData.Phone, formData.Password).First(&Data).Error; err != nil {
			return err
		}
		token, tokerr := Middlewares.GenToken(Data.UserId.String, Data.Username)
		if tokerr != nil {
			return tokerr
		}
		if err2 := tx.Model(&Models.User{}).Where("user_id=?", Data.UserId.String).Update("token", token).Error; err2 != nil {
			return err2
		}
		return nil

	})
	if err3 != nil {
		c.JSON(http.StatusConflict, gin.H{
			"ms":  "请确认是否有误",
			"res": err3.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"result": Data,
	})

}

// @Description 关于用户注册
// @Tags 注册
// @Accept x-www-form-urlencoded
// @Produce x-www-form-urlencoded
// @Param Phone formData  string true "用户手机号"
// @Param Password formData   string true "用户密码"
// @Param Username formData  string true "用户账户名"
// @Success 200 {string} hello
// @Router /register [post]
func Register(c *gin.Context) {
	//用户注册
	value := strconv.FormatInt(rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(10000000), 10)
	fmt.Printf("得到的随机数：%v", value)
	NowIp := c.ClientIP()
	fmt.Printf("得到的访问ip：%v", NowIp)
	var Register Models.User
	err := c.Bind(&Register)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}
	Register.NowIp = NowIp
	Register.InitIp = NowIp
	Register.UserId = sql.NullString{
		String: value,
		Valid:  true,
	}
	Err := Untils.Db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Model(&Models.User{}).Create(&Register).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		d, err1 := Register.UserId.Value()
		if err1 != nil {
			return err1
		}
		if err2 := tx.Model(&Models.User{}).Debug().Where("user_id =?", d).Update("Token", value).Error; err2 != nil {
			return err2
		}
		// 返回 nil 提交事务
		return nil
	})
	if Err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "请检查您的信息重新注册",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":   "注册成功",
			"token": value,
		})
	}
	//result := MysqlDo.Db.Model(Models.User{}).Create(&Register).GetErrors()

}
