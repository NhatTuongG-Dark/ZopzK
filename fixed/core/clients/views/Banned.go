package YamiView

import (
	"Yami/core/clients/branding"
	"log"
	"time"

	"golang.org/x/crypto/ssh"
)



// Banned Screen for when a client is offline
func Banned(channel ssh.Channel, conn *ssh.ServerConn) {
	errors := Branding.TermFXExecute("banned-splash", conn, channel); if errors != nil {
		log.Printf("Failed to print \"banned-splash.tfx\" for client")
		time.Sleep(5 * time.Second)
		channel.Close()
		return
	}
	time.Sleep(5 * time.Second)
	channel.Close()
	return
}