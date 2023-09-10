package commands

import (
	Sessions "Yami/core/clients/session"
	YamiDB "Yami/core/db"
	"Yami/core/functions"
	JsonParse "Yami/core/models/Json"
	"fmt"
	"net/url"
	"strings"

	"github.com/xlzd/gotp"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

// clear command
func init() {
	Register(&Command{
		Name: "mfa",
		Admin: false,
		Reseller: false,

		Descriptions: "enabled & disabled MFA on this account!",
		Execute: func(channel ssh.Channel, sshConn *ssh.ServerConn, Session *Sessions.Session, cmd []string) error {
			if len(cmd) < 2 {
				return nil
			}

			switch cmd[1] {

			case "on":
				fmt.Fprintln(channel, "PRESS ENTER")

				Byteess := make([]byte, 4)
				_, error := channel.Read(Byteess)
				channel.Write([]byte("[8;55;100t\r\n"))
			
				secret := functions.GenTOTPSecret()
			
				totp := gotp.NewDefaultTOTP(secret)
			
				qrs := functions.New()
			
			
				qrcode := qrs.Get("otpauth://totp/" + url.QueryEscape(JsonParse.ConfigSyncs.Controls.MFA.AppName) + ":" + url.QueryEscape(sshConn.User()) + "?secret=" + secret + "&issuer=" + url.QueryEscape(JsonParse.ConfigSyncs.Controls.MFA.AppName) + "&digits=6&period=30").Sprint()
				fmt.Fprintln(channel, strings.ReplaceAll(qrcode, "\n", "\r\n"))
				fmt.Fprint(channel, "You may scan this code to register your account info a 2FA App, Google Auth, Twilio Authy\r\n")
				fmt.Fprint(channel, "or enter this code> "+secret+"\r\n")
				term := terminal.NewTerminal(channel, "Code>")
			
				Code, error := term.ReadLine(); if error != nil {
					channel.Write([]byte("\r\n"))
					return nil
				}
			
				if totp.Now() != Code {
					fmt.Fprintln(channel, "Invaild MFA Code")
					return nil
				}
			
				errorss := YamiDB.EditFeild(sshConn.User(), "MFA", secret); if !errorss {
					fmt.Fprintln(channel, "Failed to enable MFA!", error)
					return nil
				} else {
					fmt.Fprintln(channel, "MFA Has been enabled!\r\n")
					channel.Write([]byte("[8;31;80t"))
					return nil
				}
				return nil

			case "off":

				User, _ := YamiDB.GetUser(sshConn.User())
				Term := terminal.NewTerminal(channel, "MFA Code>")

				Code, error := Term.ReadLine(); if error != nil {
					channel.Close()
					return error
				}
			
			
				TOTP := gotp.NewDefaultTOTP(User.MFA)
				if TOTP.Now() != Code {
					channel.Write([]byte("Invaild MFA Code!\r\n"))
					return nil
				} else {
					error := YamiDB.EditFeild(sshConn.User(), "MFA", "0"); if !error {
						channel.Write([]byte("Failed to remove MFA from your account\r\n"))
						return nil
					}

					channel.Write([]byte("MFA has been removed from your account\r\n"))
					return nil
				}
			}
			return nil
		},
	})
}