package events

import (
	"blackoak.cloud/balout/v2/helper/gosf"
)

func Routers() {
	gosf.Listen("balout:system:ping", ping)
	gosf.Listen("balout:system:error", echo)

	gosf.Listen("balout:dev", dev)

	gosf.OnConnect(OnConnectHandler)
	gosf.OnDisconnect(OnDisconnectHandler)

	gosf.OnBeforeRequest(BeforeRequestHandler)

	player := new(Player)
	gosf.Listen("balout:player:identity", player.playerIdentity)

	match := new(Match)
	gosf.Listen("balout:match:player:ready", match.matchStart)
	gosf.Listen("balout:match:player:act", match.act)
	gosf.Listen("balout:match:player:act:retry", match.act)
	gosf.Listen("balout:match:cheat", match.cheat)
	gosf.Listen("balout:match:player:leave", match.leave)

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
