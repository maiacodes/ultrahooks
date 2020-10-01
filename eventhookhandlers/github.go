package eventhookhandler

import (
	"fmt"

	"../webhooktypes"
	embed "github.com/clinet/discordgo-embed"
)

func github(event webhooktypes.WebhookEvent) (response webhooktypes.EventResponse, err error) {
	action := event.Body["action"]
	eventEmbed := embed.NewEmbed().SetTitle(fmt.Sprintf("%v", action))

	var eventResponse webhooktypes.EventResponse
	eventResponse.Embed = eventEmbed.MessageEmbed

	return eventResponse, nil
}
