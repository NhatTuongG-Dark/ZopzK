package YamiView

import (
	YamiDB "Yami/core/db"
	"Yami/core/functions"
	"Yami/core/models/Json"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/xlzd/gotp"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

func NewMFA(channel ssh.Channel, conn *ssh.ServerConn) error {

	fmt.Fprintln(channel, "PRESS ENTER")

	Byteess := make([]byte, 4)
	_, error := channel.Read(Byteess)
	channel.Write([]byte("[8;55;100t\r\n"))

	secret := functions.GenTOTPSecret()

	totp := gotp.NewDefaultTOTP(secret)

	qrs := functions.New()


	qrcode := qrs.Get("otpauth://totp/" + url.QueryEscape(JsonParse.ConfigSyncs.Controls.MFA.AppName) + ":" + url.QueryEscape(conn.User()) + "?secret=" + secret + "&issuer=" + url.QueryEscape(JsonParse.ConfigSyncs.Controls.MFA.AppName) + "&digits=6&period=30").Sprint()
	fmt.Fprintln(channel, strings.ReplaceAll(qrcode, "\n", "\r\n"))
	fmt.Fprint(channel, "You may scan this code to register your account info a 2FA App, Google Auth, Twilio Authy\r\n")
	fmt.Fprint(channel, "or enter this code> "+secret+"\r\n")
	term := terminal.NewTerminal(channel, "Code>")

	Code, error := term.ReadLine(); if error != nil {
		fmt.Fprint(channel, "\r\n")
		return errors.New("KILLING")
	}

	if totp.Now() != Code {
		fmt.Fprintln(channel, "Invaild MFA Code")
		return errors.New("KILLING") 
	}

	errorss := YamiDB.EditFeild(conn.User(), "MFA", secret); if !errorss {
		fmt.Fprintln(channel, "Failed to enable MFA!", error)
		return errors.New("KILLING") 
	} else {
		channel.Write([]byte("[8;31;80t"))
	}

	return nil
}