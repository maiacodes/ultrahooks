package main

import (
	"./webhooktypes"
)

func dispatchMessage(webhook webhooktypes.WebhookEvent, response webhooktypes.EventResponse) error {
	_, err := session.ChannelMessageSendEmbed(webhook.Subscription.ChannelID, response.Embed)
	return err
}
