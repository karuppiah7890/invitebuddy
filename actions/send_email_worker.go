package actions

import (
	"fmt"
	"log"

	"github.com/gobuffalo/buffalo/worker"
	"github.com/karuppiah7890/invitebuddy/mailers"
)

var w worker.Worker

func init() {
	w = App().Worker
	w.Register("send_email", func(args worker.Args) error {
		event := Event{
			Name:        args["eventName"].(string),
			Description: args["eventDescription"].(string),
			Email:       args["email"].(string),
		}
		content := fmt.Sprintf("You have been invited to an event. %s - %s", event.Name, event.Description)

		err := mailers.SendEmail(event.Email, content)

		if err != nil {
			log.Println(err)
		}

		return err
	})
}
