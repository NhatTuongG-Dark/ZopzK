package YamiSshServer

import (
	"net"

	"golang.org/x/crypto/ssh"
)

func Handler(conn net.Conn) {
	sshConn, chans, reqs, error := ssh.NewServerConn(conn, SSHConfig); if error != nil {
		return
	} 

	go ssh.DiscardRequests(reqs)
	handleChannels(chans, sshConn)
}