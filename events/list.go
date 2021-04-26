package events

import (
	"github.com/ambelovsky/gosf"
)

func Routers() {
	gosf.Listen("balout:system:ping", ping)
	gosf.Listen("balout:system:error", echo)

	gosf.Listen("balout:dev", dev)

	gosf.Listen("event", echo)
	gosf.Listen("disconnect", echo)
	gosf.Listen("connect", echo)

	gosf.Listen("balout:player:invalid-token", playerIdentity)
	gosf.Listen("balout:player:valid-token", playerIdentity)
	gosf.Listen("balout:player:identity", playerIdentity)

	gosf.Listen("balout:match:start", matchStart)
	gosf.Listen("balout:match:waiting", matchStart)
	gosf.Listen("balout:match:progress", matchStart)
	gosf.Listen("balout:match:finish", matchStart)
	gosf.Listen("balout:match:error", matchStart)
	gosf.Listen("balout:match:alert:is-same", matchStart)
	gosf.Listen("balout:match:you-are-disconnected", matchStart)
	gosf.Listen("balout:match:cheat", matchStart)
	gosf.Listen("balout:match:disconnect-other-player", matchStart)
	gosf.Listen("balout:match:join-again-other-player", matchStart)
	gosf.Listen("balout:match:player:act", matchStart)
	gosf.Listen("balout:match:player:act:retry", matchStart)
	gosf.Listen("balout:match:player:leave", matchStart)

	gosf.Listen("balout:chat:send:ack", sendMessage)
	gosf.Listen("balout:chat:inbox:latest", latestMessage)
}
