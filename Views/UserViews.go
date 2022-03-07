package Views

import (
	"PetService/Models"
	"PetService/MysqlDo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {

}

func (UserController)UserGet(c *gin.Context) {

	res:=MysqlDo.Db.Find(Models.User{})
	c.JSON(200,gin.H{
		"msg":"返回成功",
		"result":res,
	})
}

func (UserController)UserPut(c *gin.Context) {
	//更新当前表所有数据的特定字段
	//MysqlDo.Db.Model(&Models.User{}).Update("Password","789456")
	//更新当前表条件查询到的数据进行字段的更新-Update必须放后面
	//MysqlDo.Db.Model(&Models.User{}).Where("username=?","周zz").Update("Password","55555")
	//更新多列的字段
	value:=map[string]interface{}{"username":"tt","Password":"88888"}

	print(MysqlDo.Db.Debug().Model(&Models.User{}).Updates(value))
	c.JSON(200,gin.H{
		"msg":"ok",
	})

}

func (UserController)UserDelete(c *gin.Context) {

	MysqlDo.Db.Delete(&Models.User{},1)
	c.JSON(200,gin.H{
		"msg":"Delete",
	})

}

func (UserController)UserPost(c *gin.Context) {
	//user:= Models.User{
	//	Username: "周zz",
	//	Password: "123456",
	//	Phone: "18381894430",
	//	CreateAddress: "达州市宣汉县",
	//	CreateCity: "成都",
	//	InitIp: "192.168.3.199",
	//	NowIp: "192.168.3.199",
	//	Token: "xjakjsxjas43123322sda3d",
	//	UserId: sql.NullString{
	//		String: "262xjsak557765s",
	//		//valud如果是false 数据存入不进去
	//		Valid: true,
	//	},
	//
	//}
	//MysqlDo.Db.Debug().Create(&user)--该操作会显示执行的sql语句
	//MysqlDo.Db.Create(&user)
	var u Models.User
	//find适用于主键为id，并且是int-返回第一条
	//MysqlDo.Db.Find(&u)
	//Take适用于主键为id，并且是int-随机返回一条
	//MysqlDo.Db.Take(&u)
	//Last适用于主键为id，并且是int-返回最后一条
	MysqlDo.Db.Last(&u)
	// us []Models.User find 获取用户表所有的,MysqlDo.Db.Find(&us,3)加上3 类似返回多少条
	var us []Models.User
	MysqlDo.Db.Find(&us,3)
	c.JSON(http.StatusCreated,gin.H{
		"msg":u,
		"ms":us,
	})
	//c.String(200,"UserPost")
}
