package events

import (
	"blackoak.cloud/balout/v2/components/log"
	"blackoak.cloud/balout/v2/helper/gosf"
)

func EventsList() {
	log.Log("Server is running")

	gosf.Listen("balout:system:ping", ping)
	gosf.Listen("balout:system:error", echo)

	gosf.Listen("balout:dev", dev)

	// gosf.OnConnect(OnConnectHandler)
	// gosf.OnDisconnect(OnDisconnectHandler)
	// gosf.OnBeforeRequest(BeforeRequestHandler)

	gosf.Listen("balout:player:authenticate", authenticate)
	gosf.Listen("balout:player:identity", playerIdentity)

	gosf.Listen("balout:match:player:new-match", creatNewMatch)
	gosf.Listen("balout:match:player:ready", matchStart)
	gosf.Listen("balout:match:player:act", act)
	gosf.Listen("balout:match:cheat", cheat)
	gosf.Listen("balout:match:player:leave", leaveSingleRoom)
	gosf.Listen("balout:match:player:leave:all", leaveAllRoom)

	gosf.Listen("balout:chat:send:ack", sendMessage)
	gosf.Listen("balout:chat:inbox:latest", latestMessage)
}
