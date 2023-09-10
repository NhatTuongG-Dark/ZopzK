package commands

import (
	BinLoad "Yami/core/models/bin"
	"io"
	"os/exec"

	"github.com/creack/pty"
	"golang.org/x/crypto/ssh"
)

var finish = true

func NewPlayer(channel ssh.Channel, cmd []string, Command *BinLoad.CommandBin) {

	c := exec.Command(Command.Execute, cmd[1:]...)
	f, err := pty.Start(c)
	if err != nil {
		panic(err)
	}

	finish = false

	go func() {
		for {
			if finish {
				return
			}
			buf := make([]byte, 1000)
			n, _ := channel.Read(buf)
			f.Write(buf[:n])
		}

	}()


	io.Copy(channel, f)

	finish = true

	return
}