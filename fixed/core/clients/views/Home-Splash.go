package YamiView

import (
	"Yami/core/clients/branding"
	"Yami/core/clients/session"
	"log"

	"golang.org/x/crypto/ssh"
)

// Home screen
func HomeSplash(channel ssh.Channel, conn *ssh.ServerConn, Session *Sessions.Session, oldrequests <-chan *ssh.Request) {

	channel.Write([]byte("[8;31;80t"))
	errors := Branding.TermFXExecute("home-splash", conn, channel); if errors != nil {
		log.Printf("Failed to print \"home-splash.tfx\" for client")
		return
	}

	Prompt(channel, conn, Session, oldrequests)


}