package events

import (
	"fmt"

	"blackoak.cloud/balout/v2/helper/gosf"
)

func Routers() {
	fmt.Println("Server is running")

	gosf.Listen("balout:system:ping", ping)
	gosf.Listen("balout:system:error", echo)

	gosf.Listen("balout:dev", dev)

	gosf.OnConnect(OnConnectHandler)
	gosf.OnDisconnect(OnDisconnectHandler)
	// gosf.OnBeforeRequest(BeforeRequestHandler)

	gosf.Listen("balout:player:authenticate", authenticate)
	gosf.Listen("balout:player:identity", playerIdentity)

	gosf.Listen("balout:match:player:ready", matchStart)
	gosf.Listen("balout:match:player:act", act)
	gosf.Listen("balout:match:player:act:retry", act)
	gosf.Listen("balout:match:cheat", cheat)
	gosf.Listen("balout:match:player:leave", leave)

	// gosf.Listen("balout:match:start", matchStart)
	// gosf.Listen("balout:match:waiting", matchStart)
	// gosf.Listen("balout:match:progress", matchStart)
	// gosf.Listen("balout:match:finish", matchStart)
	// gosf.Listen("balout:match:error", matchStart)
	// gosf.Listen("balout:match:alert:is-same", matchStart)
	// gosf.Listen("balout:match:you-are-disconnected", matchStart)
	// gosf.Listen("balout:match:cheat", matchStart)
	// gosf.Listen("balout:match:disconnect-other-player", matchStart)
	// gosf.Listen("balout:match:join-again-other-player", matchStart)
	// gosf.Listen("balout:match:player:act", matchStart)
	// gosf.Listen("balout:match:player:act:retry", matchStart)
	// gosf.Listen("balout:match:player:leave", matchStart)

	gosf.Listen("balout:chat:send:ack", sendMessage)
	gosf.Listen("balout:chat:inbox:latest", latestMessage)
}
