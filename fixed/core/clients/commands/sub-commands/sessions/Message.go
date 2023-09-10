package SessionsCommands

import (
	Sessions "Yami/core/clients/session"
	YamiDB "Yami/core/db"
	"strconv"
	"strings"

	"golang.org/x/crypto/ssh"
)

func Message(channel ssh.Channel, sshConn *ssh.ServerConn, cmd []string) {


	if len(cmd) <= 3 {
		channel.Write([]byte("Syntax Correction: sessions message <User>@<Session Id> <Message>\r\n"))
		return
	}


	args := strings.Split(cmd[2], "@")

	Test := YamiDB.Exists(cmd[0]); if Test {
		channel.Write([]byte("User doesnt exist!\r\n"))
		return
	}

	for i, s := range Sessions.Sessions {
		if strconv.Itoa(int(i)) == args[1] && s.User.Username == args[0] {
			s.Channel.Write([]byte("\x1b[0m\x1b7\x1b[1A\r\x1b[2K\x1b[0m	"+sshConn.User()+"\x1b[38;5;105m>\x1b[0m "+string(strings.Join(cmd[3:], " "))+"\x1b[0m\x1b8"))
			return
		}
	}

	channel.Write([]byte("Session doesnt exist!\r\n"))
	return
}