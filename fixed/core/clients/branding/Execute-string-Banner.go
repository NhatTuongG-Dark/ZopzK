package Branding

import (
	"Yami/core/db"
	"Yami/core/functions"
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"

	"Yami/core/clients/session"
	"Yami/core/models/tfx"
)

func TermFXExecute(name string, conn *ssh.ServerConn, channel ssh.Channel) error {

	New := termfx.New()


	New.RegisterVariable("username", conn.User())
	var Users YamiDB.User

	for _, s := range Sessions.Sessions {
		if s.Conn.User() == conn.User() {
			Users = *s.User
		}
	}

	if Users.Username == "" {
		return fmt.Errorf("Failed to grab user details from session")
	}

	lol := time.Duration(time.Until(time.Unix(Users.PlanExpiry, 0))).Hours()/24

	running,_ := YamiDB.GetRunning()


	New.RegisterVariable("daysleft", strconv.FormatFloat(lol, 'f', 2, 64))

	New.RegisterVariable("motd", functions.Motd)
	New.RegisterVariable("clear", "\033c")
	New.RegisterVariable("running", strconv.Itoa(len(strconv.Itoa(running))))
	New.RegisterVariable("powersavingmode", ColourizeBoolen(Users.PowerSavingExempt))
	New.RegisterVariable("bypassblacklist", ColourizeBoolen(Users.BypassBlacklist))
	New.RegisterVariable("admin", ColourizeBoolen(Users.Admin))
	New.RegisterVariable("banned", ColourizeBoolen(Users.Banned))
	New.RegisterVariable("reseller", ColourizeBoolen(Users.Reseller))
	New.RegisterVariable("vip", ColourizeBoolen(Users.Vip))
	New.RegisterVariable("maxtime", strconv.Itoa(Users.MaxTime))
	New.RegisterVariable("cooldown", strconv.Itoa(Users.Cooldown))
	New.RegisterVariable("conrurrents", strconv.Itoa(Users.Concurrents))
	New.RegisterVariable("maxsessions", strconv.Itoa(Users.MaxSessions))
	New.RegisterVariable("sent", strconv.Itoa(YamiDB.AmmountSent()))
	New.RegisterVariable("mysent", strconv.Itoa(YamiDB.MySent(conn.User())))

	New.RegisterFunction("sleep", func(session io.Writer, args string) (int, error) {

		sleep, err := strconv.Atoi(args)
		if err != nil {
			return 0, err
		}

		time.Sleep(time.Millisecond * time.Duration(sleep))
		return 0, nil
	})

	for _, s := range BrandingWorker {
		if s.CommandName == name {
			scan := bufio.NewScanner(strings.NewReader(s.CommandContains))
			for scan.Scan() {
				File, _:= New.ExecuteString(scan.Text())
				File = strings.Replace(File, "\\x1b", "", -1)
				File = strings.Replace(File, "\\e", "", -1)
				channel.Write([]byte(File+"\r\n"))
			}
		}
	}
	return nil
}

