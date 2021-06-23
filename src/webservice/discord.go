package webservice

import (
	"fmt"
	"github.com/alecthomas/chroma/lexers"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gotification/src/extract"
	"gotification/src/notify"
	"gotification/src/util"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
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

func discordSendFileContents(c *gin.Context) {
	channelName := c.Param("channelName")
	if channelName == "" {
		c.String(http.StatusBadRequest, "No channel")
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "could not get file: %s", err.Error())
		return
	}
	filename := uuid.New().String() + filepath.Ext(file.Filename)

	dir, err := ioutil.TempDir("", "discord-send-file-contents")
	if err != nil {
		c.String(http.StatusInternalServerError, "could not get directory to save file in : %s", err.Error())
		return
	}
	defer os.RemoveAll(dir)

	path := filepath.Join(dir, filename)

	if err = c.SaveUploadedFile(file, path); err != nil {
		c.String(http.StatusInternalServerError, "could not get save file: %s", err.Error())
		return
	}

	theFile, err := os.Open(path)
	if err != nil {
		c.String(http.StatusInternalServerError, "could not read file back in: %s", err.Error())
		return
	}
	defer theFile.Close()

	contents, err := extract.Extract().ExtractFile(theFile)
	if err != nil {
		c.String(http.StatusInternalServerError, "could not read file contents: %s", err.Error())
		return
	}

	language := ""
	languageMatch := lexers.Match(file.Filename)
	if languageMatch.Config() != nil {
		aliases := languageMatch.Config().Aliases
		if len(aliases) > 0 {
			language = aliases[0]
		}
	}

	_, err = theFile.Seek(0, 0)
	if err != nil {
		c.String(http.StatusInternalServerError, "could not rewind file", err.Error())
		return
	}

	discordEmbedFileMessage := discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title:       filepath.Base(file.Filename),
			Description: contents,
		},
		Files: []*discordgo.File{
			{
				Name:        file.Filename,
				ContentType: http.DetectContentType([]byte(contents)),
				Reader:      theFile,
			},
		},
	}

	message := notify.DiscordMessageRequest{MessageSend: discordEmbedFileMessage, Language: language}

	err = notify.Discord().SendMessage(channelName, message)
	if err != nil {
		util.DoHttpError(c, http.StatusBadRequest, fmt.Errorf("ERROR|webservice/discord.discordSendFileContents()|error sending message|%s", err.Error()))
		return
	}

	c.String(http.StatusOK, "message sent")
}
