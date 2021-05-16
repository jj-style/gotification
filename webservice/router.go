package webservice

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/discord/:channelName", discordSendMessage)
	return r
}
