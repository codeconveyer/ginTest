package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wubin/gin-test/gormdb"
	"github.com/wubin/gin-test/model"
)

func getUser(c *gin.Context) {
	rId := c.Query("id")
	var user model.User
	if rId == "" {
		//c.Abort()
		c.JSON(404, fail("找不到相关资源"))
		return
	}
	if err := gormdb.DB.Where("id = ?", rId).First(&user).Error; err != nil {
		c.JSON(400, fail("用户不存在"))
		return
	}
	c.JSON(200, user)
}

func createUser(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, fail(fmt.Sprintf("%s", err)))
	}

}

func Init(eng *gin.Engine) {
	userGroup := eng.Group("/users")
	userGroup.GET("/info", getUser)
	userGroup.POST("/register", createUser)
	//register("user", func(router *gin.Engine) {
	//	v1UserGroup := router.Group("api/v1/user")
	//	v1UserGroup.POST("/register", userRegister)
	//	v1UserGroup.POST("/login", userLogin)
	//	v1UserGroup.POST("/logout", middleware.SessionAuth, userLogout)
	//})
}