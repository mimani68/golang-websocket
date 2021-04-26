package events

import (
	"github.com/ambelovsky/gosf"
)

func EventList() {
	gosf.Listen("ping", Ping)
}
