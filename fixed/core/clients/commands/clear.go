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
		Name: "cls",
		Admin: false,
		Reseller: false,

		Descriptions: "fully clear your terminal screen!",
		Execute: func(channel ssh.Channel, sshConn *ssh.ServerConn, Session *Sessions.Session, cmd []string) error {

			errors := Branding.TermFXExecute("clear-splash", sshConn, channel); if errors != nil {
				log.Printf("Failed to print \"clear-splash.tfx\" for client")
				return nil
			}

			return nil
		},
	})
}