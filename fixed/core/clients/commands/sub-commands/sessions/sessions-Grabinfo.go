package SessionsCommands

import (
	Sessions "Yami/core/clients/session"
	YamiDB "Yami/core/db"
	"fmt"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

func GrabINFO(channel ssh.Channel, sshConn *ssh.ServerConn, cmd []string) {

	if len(cmd) != 3 {
		channel.Write([]byte("Syntax Correction: sessions info <User>@<Session ID>\r\n"))
		return
	}

	args := strings.Split(cmd[2], "@")

	Test := YamiDB.Exists(cmd[0]); if Test {
		channel.Write([]byte("User doesnt exist!\r\n"))
		return
	}

	for i, s := range Sessions.Sessions {
		if strconv.Itoa(int(i)) == args[1] && s.User.Username == args[0] {
			channel.Write([]byte("Session Uptime: "+fmt.Sprintf("%.2f mins", time.Since(s.OpenAt).Minutes())+", Connected From: "+s.Conn.RemoteAddr().String()+"\r\n"))
			return
		}
	}

	channel.Write([]byte("Session doesnt exist!\r\n"))
	return
}