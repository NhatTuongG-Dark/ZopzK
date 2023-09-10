package YamiView

import (
	"Yami/core/clients/branding"
	"Yami/core/db"
	"log"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

// new screen splash 
func NewUser(channel ssh.Channel, conn *ssh.ServerConn) {

	errors := Branding.TermFXExecute("new-user", conn, channel); if errors != nil {
		log.Printf("Failed to print \"new-user.tfx\" for client")
		return
	}

	PasswordBefore := terminal.NewTerminal(channel, "Password>")
	PasswordBeforeCom,_ := PasswordBefore.ReadLine()

	PasswordAfter := terminal.NewTerminal(channel, "Confirm Password>")
	PasswordAfterCom,_ := PasswordAfter.ReadLine()

	if PasswordBeforeCom != PasswordAfterCom {
		channel.Write([]byte("Passwords do not match!"))
		time.Sleep(10 * time.Second)
		channel.Close()
		return
	}
	PasswordBeforeCom = YamiDB.HashPassword(PasswordBeforeCom)

	error := YamiDB.EditFeild(conn.User(), "password", PasswordBeforeCom); if !error {
		channel.Write([]byte("Failed to update password"))
		time.Sleep(10 * time.Second)
		channel.Close()
		return
	}

	YamiDB.EditFeild(conn.User(), "NewUser", "0")

	return
}