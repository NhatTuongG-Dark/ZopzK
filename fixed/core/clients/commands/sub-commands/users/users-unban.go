package subcommands

import (
	"Yami/core/db"

	"golang.org/x/crypto/ssh"
)

//makes the user account admin
func Unban(cmd []string, conn *ssh.ServerConn, channel ssh.Channel) {

	if len(cmd) != 3 {
		channel.Write([]byte("Your request must contain a username!\r\n"))
		return
	}

	error := YamiDB.Exists(cmd[2]); if !error {
		channel.Write([]byte("User doesnt exist in the database\r\n"))
		return
	}

	Users, err := YamiDB.GetUser(cmd[2]); if err != nil {
		channel.Write([]byte("Failed to fetch user details from the database\r\n"))
		return
	}

	if !Users.Banned {
		channel.Write([]byte("User is already unbanned!\r\n"))
		return
	}

	Issues := YamiDB.EditFeild(cmd[2], "banned", "0"); if !Issues {
		channel.Write([]byte("Failed to unban user from the database\r\n"))
		return
	} else {
		channel.Write([]byte("\x1b[38;5;11mUser has been Unbanned correctly\x1b[0m\r\n"))
	}

}