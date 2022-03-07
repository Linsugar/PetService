package Views

import "github.com/gin-gonic/gin"

type PetController struct {

}

func (PetController)PetGet(c *gin.Context)  {
	c.JSON(200,gin.H{
		"":"sss",
	})
}

func (PetController)PetPost(c *gin.Context){
	c.String(200,"petpost")
}
