package YamiView

import (
	Branding "Yami/core/clients/branding"
	YamiDB "Yami/core/db"
	"errors"

	"github.com/xlzd/gotp"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

func MFANeeded(channel ssh.Channel, conn *ssh.ServerConn, oldrequests <-chan *ssh.Request, User *YamiDB.User) error {
	
	error := Branding.TermFXExecute("mfa-code-required", conn, channel); if error != nil {
		channel.Close()
		return error
	}

	Prompt,error := Branding.TermFXExecutePrompt("mfa-prompt", conn); if error != nil {
		channel.Close()
		return error
	}

	Term := terminal.NewTerminal(channel, Prompt)

	Code, error := Term.ReadLine(); if error != nil {
		channel.Close()
		return error
	}


	TOTP := gotp.NewDefaultTOTP(User.MFA)
	if TOTP.Now() != Code {
		channel.Write([]byte("Invaild MFA Code!\r\n"))
		channel.Close()
		return errors.New("Invaild MFA Code")
	}

	return nil
}