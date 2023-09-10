package Animations

import (
	"fmt"
	"time"

	"golang.org/x/crypto/ssh"
)

var Spinner bool

func NewSpinner(channel ssh.Channel) {

	Spinner = true


	Frames := []string{
		"|", "/", "-", "\\", "|", "/", "-", "\\",
	}

	for _, I := range Frames {
		fmt.Fprint(channel, "\r"+I)

		if !Spinner {
			fmt.Fprint(channel, "\r\n")
			return
		}

		time.Sleep(100 * time.Millisecond)
	}

}