package subcommands

import (
	"Yami/core/clients/session"
	"Yami/core/db"
	"strings"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

//makes the user account admin
func Remove(cmd []string, conn *ssh.ServerConn, channel ssh.Channel) {

	if len(cmd) != 3 {
		channel.Write([]byte("Syntax Correction: users remove <User>"))
		return
	}

	if cmd[2] == conn.User() {
		channel.Write([]byte("You cant remove yourself\r\n"))
		return
	}


	There := YamiDB.Exists(cmd[2]); if !There {
		channel.Write([]byte("User doesnt exist!\r\n"))
		return
	}


	channel.Write([]byte("Are you sure you want to remove this user! y/n\r\n"))
	SureQ := terminal.NewTerminal(channel, ">")

	Choice, err := SureQ.ReadLine(); if err != nil {
		channel.Write([]byte("\r\nGoodbye,\r\n"))
		return
	}

	if strings.ToLower(Choice) != "y" {
		return
	}

	Done := YamiDB.RemoveUser(cmd[2]); if !Done {
		channel.Write([]byte("Failed to remove user from database\r\n"))
		return
	}

	channel.Write([]byte("Removed user from the database\r\n"))

	for _, s := range Sessions.Sessions {
		if s.Conn.User() == cmd[2] {
			s.Channel.Close()
			return
		}
	}
	return
}