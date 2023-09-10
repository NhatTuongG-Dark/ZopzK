package YamiSshServer

import (
	"Yami/core/models/Json"
	"log"
	"net"
)

func Listen() {
	listener, error := net.Listen("tcp", JsonParse.ConfigSyncs.Masters.MasterPort); if error != nil {
		log.Println("Failed To Listen To Port, Reason:", error.Error()+".")
		return
	} else {
		log.Printf(" [SSH Connection Watcher Started] ["+JsonParse.ConfigSyncs.Masters.MasterPort+"]")
	}


	for {
		conn, error := listener.Accept(); if error != nil {
			return
		}



		
		go Handler(conn)
	}
}