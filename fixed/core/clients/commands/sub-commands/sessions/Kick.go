package SessionsCommands

import (
	Sessions "Yami/core/clients/session"
	YamiDB "Yami/core/db"
	"strconv"
	"strings"

	"golang.org/x/crypto/ssh"
)

func Kick(channel ssh.Channel, sshConn *ssh.ServerConn, cmd []string) {

	if len(cmd) != 3 {
		channel.Write([]byte("Syntax Correction: sessions kick <User>@<Session ID>\r\n"))
		return
	}

	args := strings.Split(cmd[2], "@")

	Test := YamiDB.Exists(cmd[0]); if Test {
		channel.Write([]byte("User doesnt exist!\r\n"))
		return
	}

	for i, s := range Sessions.Sessions {
		if strconv.Itoa(int(i)) == args[1] && s.User.Username == args[0] {
			s.Channel.Write([]byte("\r\nYou have been kicked from your session\r\n"))
			s.Channel.Close()
			return
		}
	}

	channel.Write([]byte("Session doesnt exist!\r\n"))
	return
}