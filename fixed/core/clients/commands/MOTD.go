package commands

import (
	Sessions "Yami/core/clients/session"
	"log"

	"Yami/core/functions"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

var MessageOfTheDay string

// clear command
func init() {
	Register(&Command{
		Name: "motd",
		Admin: false,
		Reseller: false,

		Descriptions: "Shows the current Message of the day!",
		Execute: func(channel ssh.Channel, sshConn *ssh.ServerConn, Session *Sessions.Session, cmd []string) error {
			if Session.User.Admin {
				channel.Write([]byte("Enter your new Message of the day!\r\n"))

				Term := terminal.NewTerminal(channel, "New Message of the day>")
	
				Motd, error := Term.ReadLine(); if error != nil {
					channel.Write([]byte("\r\n"))
					return nil
				}
	
				functions.Motd = Motd

				channel.Write([]byte("New MOTD: "+Motd+"\r\n"))

				log.Printf(" [MOTD Changed] [%s] [%s]", Motd, sshConn.User())

				return nil
			}

			channel.Write([]byte("Motd "+MessageOfTheDay+"..."))



			return nil
		},
	})
}
