package main

import (
	"fmt"

	"./webhooktypes"

	handler "./eventhookhandlers"
	"github.com/labstack/echo"
)

func hookEvent(id string, c echo.Context) error {
	var body map[string]interface{} = make(map[string]interface{})
	err := c.Bind(&body)
	if err != nil {
		return err
	}
	var webhookEvent webhooktypes.WebhookEvent
	webhookEvent.Subscription.Service = "github"
	webhookEvent.Subscription.ChannelID = "733250810707705877"
	response, err := handler.HandleEvent(webhookEvent)
	if err != nil {
		return err
	}
	dispatchMessage(webhookEvent, response)

	return fmt.Errorf("Webhook not found")
}
