package YamiSshServer

import (
	"Yami/core/models/Json"
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"

	"Yami/core/clients"
)

func handleChannels(chans <-chan ssh.NewChannel, sshConn *ssh.ServerConn) {
	for newChannel := range chans {
		go handleChannel(newChannel, sshConn)
	}
}

func handleChannel(newChannel ssh.NewChannel, sshConn *ssh.ServerConn) {

	if t := newChannel.ChannelType(); t != "session" {
		log.Printf("Connection Rejected As %s, Unknown Reason", t)
		errors := newChannel.Reject(ssh.UnknownChannelType, fmt.Sprintf("")); if errors != nil {
			log.Printf("Failed To Rejct Connection")
			return
		}
		return
	}

	// At this point, we have the opportunity to reject the client's
	// request for another logical connection
	channel, requests, err := newChannel.Accept()
	if err != nil {
		log.Printf("Failed To Accept Channel Connection From %s, Reason: %s.", sshConn.RemoteAddr(), err.Error())
		return
	}
	


	requestssh := make(chan *ssh.Request, 2) // request
	for req := range requests {
		switch req.Type {
		case "shell":
			if len(req.Payload) == 0 {
				err := req.Reply(true, nil); if err != nil {
					return
				}
				go func() {

					channel.Write([]byte(string("\033]0; "+JsonParse.ConfigSyncs.Branding.AppName+"\007")))
					
					Masters.New(channel, sshConn, requestssh)
				}()
			}
		case "pty-req":
			requestssh <- req

		case "window-change":
			requestssh <- req
		}
	}
}