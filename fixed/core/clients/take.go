package Masters

import (
	JsonParse "Yami/core/models/Json"
	"Yami/core/clients/session"
	"Yami/core/clients/views"
	"Yami/core/db"
	"log"
	"strconv"
	"time"

	"golang.org/x/crypto/ssh"
)

func New(channel ssh.Channel, conn *ssh.ServerConn, oldrequests <-chan *ssh.Request) {


	User, err := YamiDB.GetUser(conn.User()); if err != nil {
		channel.Write([]byte("Failed to get your account info from the database, contact FB or DownMyPath."))
		channel.Close()
		log.Printf("Failed to fetch details for %s from database", err.Error())
		return
	}

	channel.Write([]byte{
		255, 251, 1,
		255, 251, 3,
		255, 252, 34,
	})




	var session = &Sessions.Session{
		ID: 			time.Now().Unix(),
		User: 		User,
		
		Channel: 		channel,
		Conn: 		*conn,	

		Theme:		"standard",
		OpenAt: 		time.Now(),
		SpamLimiter:   0,
		Chat: false,
	}

	Sessions.SessionMutex.Lock()
	Sessions.Sessions[session.ID] = session
	Sessions.SessionMutex.Unlock()

	go session.Check(conn)

	if User.PlanExpiry < time.Now().Unix() {
		YamiView.PlanEnded(channel, conn)
		channel.Close()
		return
	}



	if User.Banned {
		YamiView.Banned(channel, conn)
		return
	}

	open := session.Open(conn)


	if open > User.MaxSessions {
		channel.Write([]byte("You have "+strconv.Itoa(open)+" Sessions Open and your limit is "+strconv.Itoa(User.MaxSessions)+".\r\n"))
		time.Sleep(5 * time.Second)
		channel.Close()
		return
	}

	if User.NewUser {
		YamiView.NewUser(channel, conn)
	}

	if JsonParse.ConfigSyncs.Controls.Catpcha.Status {
		if JsonParse.ConfigSyncs.Controls.Catpcha.AdminBypass && User.Admin {
		} else {
			YamiView.Catpcha(channel, conn, oldrequests)
		}
	}

	if len(User.MFA) > 1 {
		error := YamiView.MFANeeded(channel, conn, oldrequests, User); if error != nil {
			return
		}
	}

	if User.MFA == "0" && JsonParse.ConfigSyncs.Controls.MFA.ForceMFA {
		if JsonParse.ConfigSyncs.Controls.MFA.AdminBypassForce && User.Admin {

		} else {
			log.Printf(" [Now Forcing MFA] [%s]", conn.User())
			error := YamiView.NewMFA(channel, conn); if error != nil {
				channel.Close()
				return
			}
		}
	}


	go Title(channel, conn)


	YamiView.HomeSplash(channel, conn, session, oldrequests)
}
