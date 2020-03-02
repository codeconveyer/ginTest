package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/storyicon/grbac"
	"github.com/wubin/gin-test/api"
	_ "github.com/wubin/gin-test/gormdb"
	"github.com/wubin/gin-test/model"
	"net/http"
)

var db *gorm.DB

//func LoadAuthorizationRules() (rules grbac.Rules, err error) {
//	// 在这里实现你的逻辑
//	// ...
//	// 你可以从数据库或文件加载授权规则
//	// 但是你需要以 grbac.Rules 的格式返回你的身份验证规则
//	// 提示：你还可以将此函数绑定到golang结构体
//	return
//}

func QueryRolesByHeaders(header http.Header) (roles []string, err error) {
	// 在这里实现你的逻辑
	// ...
	// 这个逻辑可能是从请求的Headers中获取token，并且根据token从数据库中查询用户的相应角色。
	rId := header.Get("X-Consumer-Custom-ID")
	var user model.User
	db.Where("id = ?", rId).First(&user)
	return roles, err
}

func Authorization() gin.HandlerFunc {
	// 在这里，我们通过“grbac.WithLoader”接口使用自定义Loader功能
	// 并指定应每分钟调用一次LoadAuthorizationRules函数以获取最新的身份验证规则。
	// Grbac还提供一些现成的Loader：
	// grbac.WithYAML
	// grbac.WithRules
	// grbac.WithJSON
	// ...
	rbac, err := grbac.New(grbac.WithYAML("../gin-test/role.yml", -1))
	if err != nil {
		panic(err)
	}
	return func(c *gin.Context) {
		roles, err := QueryRolesByHeaders(c.Request.Header)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		state, _ := rbac.IsRequestGranted(c.Request, roles)
		if !state.IsGranted() {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}

func main() {
	c := gin.New()
	api.Init(c)
	//c.Use(Authorization())
	// 在这里通过c.Get、c.Post等函数绑定你的API
	// ...

	c.Run(":9789")
}
