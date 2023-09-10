package YamiView

import (
	Branding "Yami/core/clients/branding"
	"log"
	"time"

	"golang.org/x/crypto/ssh"
)

// plan expired menu for users plan menu
func PlanEnded(channel ssh.Channel, conn *ssh.ServerConn) {

	error := Branding.TermFXExecute("plan-expired", conn, channel); if error != nil {
		log.Println(error)
	}

	time.Sleep(10 * time.Second)

	channel.Close()
	return
}