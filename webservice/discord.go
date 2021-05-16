package webservice

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gotify/notify"
	"gotify/util"
	"net/http"
)

func discordSendMessage(c *gin.Context) {
	channelName := c.Param("channelName")
	if channelName == "" {
		c.String(http.StatusBadRequest, "No channel")
		return
	}

	var msg notify.MessageRequest
	err := c.BindJSON(&msg)
	if err != nil {
		util.DoHttpError(c, http.StatusBadRequest, fmt.Errorf("ERROR|webservice/discord.discordSendMessage()|couldn't get message from request|%s", err.Error()))
		return
	}

	err = notify.Discord().SendMessage(channelName, msg)
	if err != nil {
		util.DoHttpError(c, http.StatusBadRequest, fmt.Errorf("ERROR|webservice/discord.discordSendMessage()|error sending message|%s", err.Error()))
		return
	}

	c.String(http.StatusOK, "")
}