package commands

import (
	"Yami/core/clients/branding"
	Sessions "Yami/core/clients/session"
	"log"

	"golang.org/x/crypto/ssh"
)

// clear command
func init() {
	Register(&Command{
		Name: "home",
		Admin: false,
		Reseller: false,

		Descriptions: "restart your session progress!",
		Execute: func(channel ssh.Channel, sshConn *ssh.ServerConn, Session *Sessions.Session, cmd []string) error {

			errors := Branding.TermFXExecute("home-splash", sshConn, channel); if errors != nil {
				log.Printf("Failed to print \"home-splash.tfx\" for client")
				return nil
			}

			return nil
		},
	})
}