package notify

import (
	"github.com/bwmarrin/discordgo"
)

type DiscordMessageRequest struct {
	discordgo.MessageSend
	Language string `json:"language"`
}

type DiscordImageRequest struct {
	Url string `json:"url"`
}
