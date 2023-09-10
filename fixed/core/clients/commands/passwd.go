package commands

import (
	Sessions "Yami/core/clients/session"
	YamiDB "Yami/core/db"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

// clear command
func init() {
	Register(&Command{
		Name: "passwd",
		Admin: false,
		Reseller: false,

		Descriptions: "Changes your current password!",
		Execute: func(channel ssh.Channel, sshConn *ssh.ServerConn, Session *Sessions.Session, cmd []string) error {


				Term := terminal.NewTerminal(channel, "New Password>\x1b[38;5;15m")
				Password, error := Term.ReadLine(); if error != nil {
					channel.Write([]byte("\r\nGoodbye,\r\n"))
					return nil
				}
	
				Terms := terminal.NewTerminal(channel, "Confirm New Password>\x1b[38;5;15m")
				ComfirmPassword, error := Terms.ReadLine(); if error != nil {
					channel.Write([]byte("\r\nGoodbye,\r\n"))
					return nil
				}
	
				if Password != ComfirmPassword {
					channel.Write([]byte("Passwords do not match!\r\n"))
					return nil
				}
	
				Password = YamiDB.HashPassword(Password)
	
	
				changed := YamiDB.EditFeild(sshConn.User(), "Password", Password); if !changed {
					channel.Write([]byte("Failed to update password correctly!\r\n"))
					return nil
				}
				Session.User.Password = Password
	
				channel.Write([]byte("Password has been changed correctly\r\n"))




			return nil
		},
	})
}