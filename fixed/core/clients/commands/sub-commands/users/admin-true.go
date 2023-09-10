package subcommands

import (
	"Yami/core/clients/session"
	"Yami/core/db"
	"strings"

	"golang.org/x/crypto/ssh"
)

//makes the user account admin
func MakeAdmin(cmd []string, conn *ssh.ServerConn, channel ssh.Channel) {


	if len(cmd) != 3 {
		channel.Write([]byte("Your request must contain a username!\r\n"))
		return
	}

	error := YamiDB.Exists(cmd[2]); if !error {
		channel.Write([]byte("User doesnt exist in the database\r\n"))
		return
	}


	Users, err := YamiDB.GetUser(cmd[2]); if err != nil {
		channel.Write([]byte("Failed to fetch user details from the database"))
		return
	}

	if Users.Admin {
		channel.Write([]byte("User is already registered as admin!"))
		return
	}


	Issues := YamiDB.EditFeild(cmd[2], "admin", "1"); if !Issues {
		channel.Write([]byte("Failed to promote user to Admin"))
		return
	} else {
		channel.Write([]byte("\x1b[38;5;11mUser has been promoted to admin correctly\x1b[0m\r\n"))
	}


	for _, s := range Sessions.Sessions {
		if strings.ToLower(s.User.Username) == cmd[2] {
			s.User.Admin = true
			s.Channel.Write([]byte("\x1b[0m\x1b7\x1b[1A\r\x1b[2K\x1b[38;5;2mYou have been promoted to admin\x1b[0m\x1b8"))
		}
		
	}

	return
}