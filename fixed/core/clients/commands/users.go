package commands

import (
	"Yami/core/clients/commands/sub-commands/users"
	Sessions "Yami/core/clients/session"

	"golang.org/x/crypto/ssh"
)

// this needs remaking
// branding reload command
func init() {
	Register(&Command{
		Name: "users",
		Admin: true,
		Reseller: true,

		Descriptions: "Users control with options and lists !",
		Execute: func(channel ssh.Channel, sshConn *ssh.ServerConn, Session *Sessions.Session, cmd []string) error {


			if len(cmd) < 2 {
				subcommands.UsersList(channel, sshConn)
				return nil
			}

			switch cmd[1] {


				case "admin=true":
					subcommands.MakeAdmin(cmd, sshConn, channel)
				case "admin=false":
					subcommands.RemoveAdmin(cmd, sshConn, channel)
				case "ban":
					subcommands.AddBan(cmd, sshConn, channel)
				case "unban":
					subcommands.Unban(cmd, sshConn, channel)
				case "reseller=true":
					subcommands.MakeReseller(cmd, sshConn, channel)
				case "reseller=false":
					subcommands.Removereseller(cmd, sshConn, channel)
				case "vip=true":
					subcommands.MakeVip(cmd, sshConn, channel)
				case "vip=false":
					subcommands.RemoveVip(cmd, sshConn, channel)
				case "create":
					subcommands.CreateUser(channel, sshConn)
				case "remove":
					subcommands.Remove(cmd, sshConn, channel)
				case "newuser=true":
					subcommands.NewUser(cmd, sshConn, channel)
				case "attacktime":
					subcommands.AddTime(cmd, sshConn, channel)
				case "concurrents":
					subcommands.AddConns(cmd, sshConn, channel)
				case "adddays":
					subcommands.AddDays(cmd, sshConn, channel)
				case "powersaving=true":
					subcommands.PowerSavingON(cmd, sshConn, channel)
				case "powersaving=false":
					subcommands.PowerSavingOff(cmd, sshConn, channel)
				case "view":
					subcommands.UsersView(cmd, sshConn, channel)
				case "bypassblacklist=true":
					subcommands.BypassBlacklistTrue(cmd, sshConn, channel)
				case "bypassblacklist=false":
					subcommands.BypassBlacklistFalse(cmd, sshConn, channel)
				case "mfa=false":
					subcommands.MFAOff(cmd, sshConn, channel)
				case "cooldown":
					subcommands.AddCooldown(cmd, sshConn, channel)
					
			}

			return nil
		},
	})
}
