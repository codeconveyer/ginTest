package ginengine

import "github.com/gin-gonic/gin"

var Engine *gin.Engine

func Init() {
	Engine = gin.New()
}

