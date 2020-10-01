package eventhookhandler

import (
	"fmt"

	"../webhooktypes"
)

func HandleEvent(event webhooktypes.WebhookEvent) (response webhooktypes.EventResponse, err error) {
	fmt.Println(event.Subscription.Service)
	service := event.Subscription.Service

	switch service {
	case "github":
		return github(event)
	}

	err = fmt.Errorf("Service not avaliable")
	return
}
