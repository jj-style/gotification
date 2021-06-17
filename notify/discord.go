package notify

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	. "gotify/util"
	"log"
	"sync"
)

type MessageRequest struct {
	discordgo.MessageSend
	Language string `json:"language"`
}

type discordNotifierImpl struct {
	session  *discordgo.Session
	channels map[string]string
}

type DiscordNotifier interface {
	SendMessage(channelName string, message MessageRequest) error
}

var (
	discordNotifier     DiscordNotifier
	discordNotifierOnce sync.Once
)

func Discord() DiscordNotifier {
	discordNotifierOnce.Do(func() {
		token := Config.Discord.Token
		discord, err := discordgo.New("Bot " + token)
		if err != nil {
			log.Fatalf("ERROR|notify/discord.Discord()|Couldn't get discord bot|%s", err.Error())
		}

		channels := make(map[string]string)
		guild := Config.Discord.Guild
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

func (d *discordNotifierImpl) SendMessage(channelName string, message MessageRequest) error {
	if channelId, exists := d.channels[channelName]; exists {
		if message.Language != "" {
			message.Content = prepareCodeBlock(message.Content, message.Language)
		}
		_, err := d.session.ChannelMessageSendComplex(channelId, &message.MessageSend)
		if err != nil {
			return fmt.Errorf("ERROR|notify/discord.SendMessage()|Failed to send message to channel (%s:%s)|%s", channelName, channelId, err.Error())
		}
	} else {
		return fmt.Errorf("ERROR|notify/discord.SendMessage()|Channel '%s' not found", channelName)
	}
	return nil
}