package subcommands

import (
	"Yami/core/clients/session"
	"Yami/core/db"
	"strconv"
	"strings"

	"golang.org/x/crypto/ssh"
)

func AddTime(cmd []string, conn *ssh.ServerConn, channel ssh.Channel) {

	if len(cmd) != 4 {
		channel.Write([]byte("Your request Syntax Must Be > users attacktime <username> <Time>\r\n"))
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

	if strconv.Itoa(Users.MaxTime) == cmd[3] {
		channel.Write([]byte("Users maxTime is already set to thisr\n"))
		return
	}

	duration, err := strconv.Atoi(cmd[3])
	if err != nil {
		channel.Write([]byte("Attack Time change request must be an int\r\n"))
		return
	}

	Issues := YamiDB.EditFeild(cmd[2], "MaxTime", strconv.Itoa(duration)); if !Issues {
		channel.Write([]byte("Failed to change users maxtime in database\r\n"))
		return
	} else {
		channel.Write([]byte("\x1b[38;5;11mUsers Attack time has been updated correctly\x1b[0m\r\n"))
	}

	for _, s := range Sessions.Sessions {
		if strings.ToLower(s.User.Username) == cmd[2] {
			s.User.Banned = true
			s.Channel.Write([]byte("\x1b[0m\x1b7\x1b[1A\r\x1b[2K\x1b[38;5;11mYour attack Time has been changed too "+strconv.Itoa(duration)+"\x1b[0m\x1b8"))
			s.Channel.Close()
			return
		}
		
	}
}