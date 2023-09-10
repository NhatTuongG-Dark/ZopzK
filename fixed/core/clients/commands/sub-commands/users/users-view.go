package subcommands

import (
	Branding "Yami/core/clients/branding"
	Sessions "Yami/core/clients/session"
	"Yami/core/db"
	"strconv"
	"time"

	"golang.org/x/crypto/ssh"
)

//makes the user account admin
func UsersView(cmd []string, conn *ssh.ServerConn, channel ssh.Channel) {
	if len(cmd) != 3 {
		channel.Write([]byte("You must include a username to view\r\n"))
		return
	}

	There := YamiDB.Exists(cmd[2]); if !There {
		channel.Write([]byte("User doesnt Exist in the database\r\n"))
		return
	}

	User, error := YamiDB.GetUser(cmd[2]); if error != nil {
		channel.Write([]byte("Failed to get users details\r\n"))
	}
	DaysLeft := time.Duration(time.Until(time.Unix(User.PlanExpiry, 0))).Hours()/24
	SecondsLeft := time.Duration(time.Until(time.Unix(User.PlanExpiry, 0))).Seconds()
	HoursLeft := time.Duration(time.Until(time.Unix(User.PlanExpiry, 0))).Hours()



	channel.Write([]byte("ID: \x1b[38;5;11m"+strconv.Itoa(User.ID)+"\x1b[0m\r\n"))
	channel.Write([]byte("Username: "+User.Username+"\x1b[0m\r\n"))
	channel.Write([]byte("Admin: "+Branding.ColourizeBoolen(User.Admin)+"\r\n"))
	channel.Write([]byte("Reseller: "+Branding.ColourizeBoolen(User.Reseller)+"\r\n"))
	channel.Write([]byte("VIP Plan: "+Branding.ColourizeBoolen(User.Vip)+"\r\n"))
	channel.Write([]byte("Banned: "+Branding.ColourizeBoolen(User.Banned)+"\r\n"))
	channel.Write([]byte("Power Saving Mode: "+Branding.ColourizeBoolen(User.PowerSavingExempt)+"\r\n"))
	channel.Write([]byte("Bypass Blacklist: "+Branding.ColourizeBoolen(User.BypassBlacklist)+"\r\n\r\n"))


	channel.Write([]byte("Days Left: \x1b[38;5;11m"+strconv.Itoa(int(DaysLeft))+"\x1b[0m\r\n"))
	channel.Write([]byte("Hours Left: \x1b[38;5;11m"+strconv.Itoa(int(HoursLeft))+"\x1b[0m\r\n"))
	channel.Write([]byte("Seconds Left: \x1b[38;5;11m"+strconv.Itoa(int(SecondsLeft))+"\x1b[0m\r\n\r\n"))

	channel.Write([]byte("MaxTime: \x1b[38;5;11m"+strconv.Itoa(User.MaxTime)+"\x1b[0m\r\n"))
	channel.Write([]byte("Cooldown: \x1b[38;5;11m"+strconv.Itoa(User.Cooldown)+"\x1b[0m\r\n"))
	channel.Write([]byte("Concurrents: \x1b[38;5;11m"+strconv.Itoa(User.Concurrents)+"\x1b[0m\r\n"))
	channel.Write([]byte("Max Sessions: \x1b[38;5;11m"+strconv.Itoa(User.MaxSessions)+"\x1b[0m\r\n"))
	fully,_ := YamiDB.GetRunning()
	channel.Write([]byte("Current Running Attacks globally: "+strconv.Itoa(fully)+"\r\n"))
	channel.Write([]byte("Global Attacks Launched: "+strconv.Itoa(YamiDB.AmmountSent())+"\r\n"))
	minerunning,_ := YamiDB.GetRunningUser(cmd[2])
	channel.Write([]byte("Current Attacks Running By "+cmd[2]+": "+strconv.Itoa(minerunning)+"\r\n"))
	channel.Write([]byte("Attacks Sent By "+cmd[2]+": "+strconv.Itoa(YamiDB.MySent(conn.User()))+"\r\n"))

	// GetRunningUser



	for _, s := range Sessions.Sessions {
		if s.User.Username == User.Username {
			channel.Write([]byte("Currently Logged In: "+Branding.ColourizeBoolen(true)+"\r\n"))
			channel.Write([]byte("Current IPv4: [38;5;26m"+s.Conn.RemoteAddr().String()+"[0m\r\n"))
			return
		}
	}

	channel.Write([]byte("Currently Logged In: "+Branding.ColourizeBoolen(false)+"\r\n"))

}