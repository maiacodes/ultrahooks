package webhooktypes

import "github.com/bwmarrin/discordgo"

type WebhookEvent struct {
	Body         map[string]interface{}
	Subscription Subscription
}

type EventResponse struct {
	Embed *discordgo.MessageEmbed
}

type Subscription struct {
	Service   string
	ChannelID string
	GuildID   int
}
