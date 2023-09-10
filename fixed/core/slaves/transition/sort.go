package transition

import (
	"log"
	"net"
)

var Count string

func Sort(conn net.Conn) {


	for {
		Bytess := make([]byte, 1000)
		_, error := conn.Read(Bytess); if error != nil {
			log.Println(" [Failed to load your transition Server]")
		}

		Count = string(Bytess)
		continue
	}

	log.Println(" [Connection with transition server Finished]")
	return
}