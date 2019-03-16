package actions

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/worker"
)

// Event represents an event
type Event struct {
	Name        string
	Description string
	Email       string
}

// SendInviteRequest represents a http request for sending invites
type SendInviteRequest struct {
	Event Event
}

// SendInviteHandler is a handler to send invites
func SendInviteHandler(c buffalo.Context) error {
	requestBodyReader := c.Request().Body
	log.Println(requestBodyReader)
	var requestBodyBuffer bytes.Buffer
	_, err := requestBodyBuffer.ReadFrom(requestBodyReader)
	if err != nil {
		return err
	}

	requestBody := requestBodyBuffer.Bytes()

	var sendInviteRequest SendInviteRequest
	json.Unmarshal(requestBody, &sendInviteRequest)

	w.Perform(worker.Job{
		Queue:   "default",
		Handler: "send_email",
		Args: worker.Args{
			"eventName":        sendInviteRequest.Event.Name,
			"eventDescription": sendInviteRequest.Event.Description,
			"email":            sendInviteRequest.Event.Email,
		},
	})

	return c.Render(200, r.JSON(map[string]string{"message": "mail will be sent"}))
}
