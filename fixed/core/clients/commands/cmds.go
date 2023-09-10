package commands

import (
	Sessions "Yami/core/clients/session"
	JsonParse "Yami/core/models/Json"
	"strings"

	"golang.org/x/crypto/ssh"
)

var (
	commands = make(map[string]*Command)
)

//Command can be executed by a opperator via the terminal
type Command struct {
	Name string

	Descriptions string
	Admin bool
	Reseller bool

	Execute func(channel ssh.Channel, sshConn *ssh.ServerConn, Session *Sessions.Session, cmd []string) error
}

//Register adds the command to the handler
func Register(cmd *Command) {
	if _, ok := commands[cmd.Name]; ok == true {
		panic("more than one command with the same name")
	}

	commands[cmd.Name] = cmd
}


//Get returns a command via it's name
func Get(name string) *Command {
	cmd := commands[name]
	return cmd
}

//Checks if the command is disabled
func Check(name string) bool {
	for i := 0; i < len(JsonParse.ConfigSyncs.DisabledCommands); i++ {
		if strings.ToLower(name) == strings.ToLower(JsonParse.ConfigSyncs.DisabledCommands[i]) {
			return true
		}
	}

	return false
}