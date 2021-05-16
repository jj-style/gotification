package notify

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"gotify/config"
	"gotify/util"
	"log"
	"os"
	"sync"
)

type DiscordMessageRequest struct {
	Message  string `json:"message"`
	Language string `json:"language"`
	Tts		 bool `json:"tts"`
}

type discordNotifierImpl struct {
	session  *discordgo.Session
	channels map[string]string
}

type DiscordNotifier interface {
	SendMessage(channelName string, message DiscordMessageRequest) error
}

var (
	discordNotifier     DiscordNotifier
	discordNotifierOnce sync.Once
)

func Discord() DiscordNotifier {
	discordNotifierOnce.Do(func() {
		token := config.Config().GetString(util.DISCORD_TOKEN)
		discord, err := discordgo.New("Bot " + token)
		if err != nil {
			log.Fatalf("ERROR|notify/discord.Discord()|Couldn't get discord bot|%s", err.Error())
		}

		channels := make(map[string]string)
		guild := config.Config().GetString(util.DISCORD_GUILD)
		chn, err := discord.GuildChannels(guild)
		if err != nil {
			log.Fatalf("ERROR|notify/discord.Discord()|Couldn't get channels from guild '%s'|%s", guild, err.Error())
		}

		for _, c := range chn {
			if c.Type == 0 {
				channels[c.Name] = c.ID
			}
		}

		// Open a websocket connection to Discord and begin listening.
		err = discord.Open()
		if err != nil {
			log.Fatalf("ERROR|notify/discord.Discord()|Couldn't open connection to discord|%s", err.Error())
		}

		discordNotifier = &discordNotifierImpl{
			session:  discord,
			channels: channels,
		}
	})
	return discordNotifier
}

func (d *discordNotifierImpl) SendMessage(channelName string, message DiscordMessageRequest) error {
	if channelId, exists := d.channels[channelName]; exists {
		if message.Language != "" {
			message.Message = prepareCodeBlock(message.Message, message.Language)
		}
		var err error
		if message.Tts {
			_, err = d.session.ChannelMessageSendTTS(channelId, message.Message)
		} else {
			_, err = d.session.ChannelMessageSend(channelId, message.Message)
		}
		if err != nil {
			return fmt.Errorf("ERROR|notify/discord.SendMessage()|Failed to send message to channel (%s:%s)|%s", channelName, channelId, err.Error())
		}
	} else {
		return fmt.Errorf("ERROR|notify/discord.SendMessage()|Channel '%s' not found", channelName)
	}
	return nil
}

func (d *discordNotifierImpl) SendFileContents(channelName string, file *os.File) error {
	content, err := extractFileContents(file)
	if err != nil {
		return fmt.Errorf("ERROR|notify/discord.SendFileContents()|Couldn't get file contents|%s", err.Error())
	}
	message := DiscordMessageRequest{
		Message:  content,
	}
	
	return d.SendMessage(channelName, message)
}