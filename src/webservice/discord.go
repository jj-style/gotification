package webservice

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"gotification/src/notify"
	"gotification/src/util"
	"net/http"
)

func discordSendMessage(c *gin.Context) {
	channelName := c.Param("channelName")
	if channelName == "" {
		c.String(http.StatusBadRequest, "No channel")
		return
	}

	var msg notify.DiscordMessageRequest
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

	c.String(http.StatusOK, "message sent")
}

func discordSendImage(c *gin.Context) {
	channelName := c.Param("channelName")
	if channelName == "" {
		c.String(http.StatusBadRequest, "No channel")
		return
	}
	var req notify.DiscordImageRequest
	err := c.BindJSON(&req)
	if err != nil {
		util.DoHttpError(c, http.StatusBadRequest, fmt.Errorf("ERROR|webservice/discord.discordSendImage()|couldn't get message from request|%s", err.Error()))
		return
	}

	discordEmbedImgMessage := discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Image: &discordgo.MessageEmbedImage{
				URL: req.Url,
			},
		},
	}
	message := notify.DiscordMessageRequest{MessageSend: discordEmbedImgMessage}

	err = notify.Discord().SendMessage(channelName, message)
	if err != nil {
		util.DoHttpError(c, http.StatusBadRequest, fmt.Errorf("ERROR|webservice/discord.discordSendMessage()|error sending message|%s", err.Error()))
		return
	}

	c.String(http.StatusOK, "message sent")
}
