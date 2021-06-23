package webservice

import (
	"github.com/gin-gonic/gin"
	"gotification/src/util"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	defaultRouter := r.Group("/")
	authorisedRouter := defaultRouter
	if util.Config.Auth.Type == "basic" {
		authorisedRouter = r.Group("/", gin.BasicAuth(util.Config.GinAccounts()))
	}
	router := defaultRouter

	router.GET("/health", health)

	if !util.Config.Discord.Disable {
		discordRouter := defaultRouter
		if !util.Config.Discord.NoAuth {
			discordRouter = authorisedRouter
		}
		discordRouter.POST("/discord/:channelName", discordSendMessage)
		discordRouter.POST("/discord/:channelName/image", discordSendImage)
		discordRouter.POST("/discord/:channelName/file", discordSendFileContents)
	}
	return r
}
