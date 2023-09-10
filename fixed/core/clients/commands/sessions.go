package commands

import (
	"Yami/core/clients/commands/sub-commands/sessions"
	Sessions "Yami/core/clients/session"

	"golang.org/x/crypto/ssh"
)

// clear command
func init() {
	Register(&Command{
		Name: "sessions",
		Admin: true,
		Reseller: false,

		Descriptions: "Session control with options and lists!",
		Execute: func(channel ssh.Channel, sshConn *ssh.ServerConn, Session *Sessions.Session, cmd []string) error {
			if len(cmd) < 2 {
				SessionsCommands.ListSessions(channel, sshConn)
				return nil
			}

			switch cmd[1] {

			case "listids":
				SessionsCommands.ListIds(channel, sshConn)

			case "kick":
				SessionsCommands.Kick(channel, sshConn, cmd)

			case "info":
				SessionsCommands.GrabINFO(channel, sshConn, cmd)
			case "message":
				SessionsCommands.Message(channel, sshConn, cmd)
			}
			return nil
		},
	})
}