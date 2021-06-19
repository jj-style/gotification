package webservice

import (
	"github.com/gin-gonic/gin"
	"gotification/src/util"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health", health)

	if !util.DISABLE_DISCORD {
		r.POST("/discord/:channelName", discordSendMessage)
		r.POST("/discord/:channelName/image", discordSendImage)
	}
	return r
}
