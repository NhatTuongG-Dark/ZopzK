package Branding

import (
	"Yami/core/models/tfx"
	"fmt"

	"golang.org/x/crypto/ssh"
)

func TermFXExecutePrompt(name string, conn *ssh.ServerConn) (string, error) {

	New := termfx.New()


	New.RegisterVariable("username", conn.User())

	for _, s := range BrandingWorker {
		if s.CommandName == name {
			File, err := New.ExecuteString(s.CommandContains)
			return File, err
		}
	}

	return "", fmt.Errorf("Failed to get your prompt file")
}