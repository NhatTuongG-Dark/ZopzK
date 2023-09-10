package transition

import (
	"Yami/core/models/Json"
	"log"
	"net"
)

// listens too the net port...
func Connection() {
	Nets, error := net.Dial("tcp", JsonParse.Option.SlaveTransition.LoopbackPort); if error != nil {
		log.Printf(" [Failed to attached] [%s]", JsonParse.Option.SlaveTransition.LoopbackPort)
		Connection()
		return
	} else {
		log.Printf(" [Connected too Mirai server] [%s]", JsonParse.Option.SlaveTransition.LoopbackPort)
		Sort(Nets)
		return
	}
}