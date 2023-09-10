package subcommands

import (
	"Yami/core/models/Json"
	YamiDB "Yami/core/db"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)


func CreateUser(channel ssh.Channel, conn *ssh.ServerConn) {


	channel.Write([]byte("tip: Make the username easy to remember\r\n"))

	UserBefore := terminal.NewTerminal(channel, "Username>")
	Username,err := UserBefore.ReadLine(); if err != nil {
		channel.Write([]byte("\r\nGoodbye,\r\n"))
		return
	}

	error := YamiDB.Exists(Username); if error {
		channel.Write([]byte("User already exists in the database\r\n"))
		return
	}

	channel.Write([]byte("tip: Make the password secure, Users can change this later, Must be over 5\r\n"))

	PassBefore := terminal.NewTerminal(channel, "Password>")
	Password,err := PassBefore.ReadLine(); if err != nil {
		channel.Write([]byte("\r\nGoodbye,\r\n"))
		return
	}

	if len(Password) <= 5 {
		channel.Write([]byte("Password must be longer then 5 Chars\r\n"))
		return
	}

	channel.Write([]byte("tip: Plan Preset, Leave this blank if you dont want it\r\n"))

	PlanBefore := terminal.NewTerminal(channel, "Plan>")
	Plan,err := PlanBefore.ReadLine(); if err != nil {
		channel.Write([]byte("\r\nGoodbye,\r\n"))
		return
	}

	if Plan != "" {
		for i := 0; i < len(JsonParse.ConfigSyncs.Plans.Presets); i++ {

			if JsonParse.ConfigSyncs.Plans.Presets[i].Name == strings.ToLower(Plan) {
				channel.Write([]byte("tip: Are you sure you want to add this user!\r\n"))
				
				ChoiceBefore := terminal.NewTerminal(channel, "y/n>")
				Choice,err := ChoiceBefore.ReadLine(); if err != nil {
					channel.Write([]byte("\r\nGoodbye,\r\n"))
					return
				}

				if strings.ToLower(Choice) != "y" {
					channel.Write([]byte("Voided, Goodbye."))
					return
				}


				var Users = YamiDB.User {
					Username: 		string(Username),
					Password:           string(Password),


					NewUser:            true,
					Admin:              false,
					Banned:             false,
					Reseller:           JsonParse.ConfigSyncs.Plans.Presets[i].Reseller,
					Vip: 			JsonParse.ConfigSyncs.Plans.Presets[i].Vip,

					MaxTime:            JsonParse.ConfigSyncs.Plans.Presets[i].MaxTime,
					Cooldown:           JsonParse.ConfigSyncs.Plans.Presets[i].Cooldown,
					Concurrents:        JsonParse.ConfigSyncs.Plans.Presets[i].Concurrents,

					MaxSessions:        JsonParse.ConfigSyncs.Plans.Presets[i].MaxSessions,
					PowerSavingExempt:  JsonParse.ConfigSyncs.Plans.Presets[i].PowerSavingExempt,
					PlanExpiry:         time.Now().Add((time.Hour*24)*time.Duration(JsonParse.ConfigSyncs.Plans.Presets[i].PlanLenDays)).Unix(),
					BypassBlacklist:    JsonParse.ConfigSyncs.Plans.Presets[i].BypassBlacklist,
				} 

				error := YamiDB.NewUser(&Users); if !error {
					channel.Write([]byte("Failed to add user into the database\r\n"))
					return
				} else {
					channel.Write([]byte("User has been added to the database\r\n"))
					return
				}



				return
			}
		}

		channel.Write([]byte("Plan preset was NOT Found!"))
		return
	} else {

		channel.Write([]byte("tip: How long the user can attack for!\r\n"))
		
		MaxTimeBefore := terminal.NewTerminal(channel, "Max Attack Time>")
		MaxTime,err := MaxTimeBefore.ReadLine(); if err != nil {
			channel.Write([]byte("\r\nGoodbye,\r\n"))
			return
		}

		MaxTimeInt, err := strconv.Atoi(MaxTime); if err != nil {
			channel.Write([]byte("This must be an Integer"))
		}

		if MaxTimeInt > 86400 {
			channel.Write([]byte("Attack Time can not be greater than a day\r\n"))
			return
		}

		channel.Write([]byte("tip: Time too wait inbetween attacks once concurrents reached!\r\n"))
		
		CooldownBefore := terminal.NewTerminal(channel, "Cooldown>")
		Cooldown,err := CooldownBefore.ReadLine(); if err != nil {
			channel.Write([]byte("\r\nGoodbye,\r\n"))
			return
		}
		
		CooldownInt, err := strconv.Atoi(Cooldown); if err != nil {
			channel.Write([]byte("This must be an Integer"))
		}


		channel.Write([]byte("tip: Max running attacks happening at once!\r\n"))

		ConcurrentsBefore := terminal.NewTerminal(channel, "Concurrents>")
		Concurrents,err := ConcurrentsBefore.ReadLine(); if err != nil {
			channel.Write([]byte("\r\nGoodbye,\r\n"))
			return
		}

		ConcurrentsInt, err := strconv.Atoi(Concurrents); if err != nil {
			channel.Write([]byte("This must be an Integer"))
		}

		channel.Write([]byte("tip: Max open sessions of the cnc!\r\n"))

		MaxSessionsBefore := terminal.NewTerminal(channel, "max Sessions>")
		MaxSessions,err := MaxSessionsBefore.ReadLine(); if err != nil {
			channel.Write([]byte("\r\nGoodbye,\r\n"))
			return
		}

		MaxSessionsInt, err := strconv.Atoi(MaxSessions); if err != nil {
			channel.Write([]byte("This must be an Integer"))
		}

		channel.Write([]byte("tip: How long this account has on the cnc in DAYS\r\n"))

		PlanBefore := terminal.NewTerminal(channel, "Days Left>")
		PlanFinish,err := PlanBefore.ReadLine(); if err != nil {
			channel.Write([]byte("\r\nGoodbye,\r\n"))
			return
		}

		PlanFinishlol, err := strconv.Atoi(PlanFinish); if err != nil {
			channel.Write([]byte("This must be an Integer"))
		}




	

		var Users = YamiDB.User {
			Username: 		string(Username),
			Password:           string(Password),


			NewUser:            true,
			Admin:              false,
			Banned:             false,
			Reseller:           false,
			Vip: 			false,

			MaxTime:            MaxTimeInt,
			Cooldown:           CooldownInt,
			Concurrents:        ConcurrentsInt,

			MaxSessions:        MaxSessionsInt,
			PowerSavingExempt:  true,
			PlanExpiry:         time.Now().Add((time.Hour*24)*time.Duration(PlanFinishlol)).Unix(),
			BypassBlacklist:    false,
		}

		error := YamiDB.NewUser(&Users); if !error {
			channel.Write([]byte("Failed to add user into the database\r\n"))
			return
		} else {
			channel.Write([]byte("User has been added to the database\r\n"))
			return
		}
	}
}