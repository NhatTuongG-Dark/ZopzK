package subcommands

import (
	YamiDB "Yami/core/db"

	"golang.org/x/crypto/ssh"
)

func MFAOff(cmd []string, conn *ssh.ServerConn, channel ssh.Channel) {
	if len(cmd) != 3 {
		channel.Write([]byte("Syntax Correction: users MFA=false <Username>\r\n"))
		return
	}

	User, error := YamiDB.GetUser(cmd[2]); if error != nil {
		channel.Write([]byte("Failed to gather user's info from the database\r\n"))
		return
	}

	if len(User.MFA) <= 1 {
		channel.Write([]byte("User already hasn't got MFA Enabled\r\n"))
		return
	}

	errors := YamiDB.EditFeild(cmd[2], "MFA", "0"); if !errors {
		channel.Write([]byte("Failed to disable MFA for user!\r\n"))
		return
	} else {
		channel.Write([]byte("MFA has been disabled for "+cmd[2]+" Correctly\r\n"))
		return
	}
}