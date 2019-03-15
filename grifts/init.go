package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/karuppiah7890/invitebuddy/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
