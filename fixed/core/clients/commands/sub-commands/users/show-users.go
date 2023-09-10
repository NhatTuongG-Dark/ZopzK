package subcommands

import (
	"Yami/core/clients/branding"
	"Yami/core/db"
	"Yami/core/functions"
	JsonParse "Yami/core/models/Json"
	livewire "Yami/core/models/dlcs/Live_Wire"
	License "Yami/core/models/license"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/alexeyco/simpletable"
	"golang.org/x/crypto/ssh"
)

func UsersList(channel ssh.Channel, conn *ssh.ServerConn) {
	channel.Write([]byte("\x1b[0m"))
	Users, err := YamiDB.GetUsers(); if err != nil {
		log.Println("Failed to fetch any users from database\r\n")
		return
	}


	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "User"},
			{Align: simpletable.AlignCenter, Text: "Admin"},
			{Align: simpletable.AlignCenter, Text: "Reseller"},
			{Align: simpletable.AlignCenter, Text: "VIP"},
			{Align: simpletable.AlignCenter, Text: "Banned"},
			{Align: simpletable.AlignCenter, Text: "MFA"},
			{Align: simpletable.AlignCenter, Text: "Plan Active"},
			{Align: simpletable.AlignCenter, Text: "Conns"},
		
		},
	}

	if License.LiveWire {
		if !JsonParse.LiveWireDLCSync.TableGradient.Status {
			for _, user := range Users {
				r := []*simpletable.Cell{
					{Text: strconv.Itoa(user.ID)},
					{Text: user.Username},
					{Text: Branding.ColourizeBoolen(user.Admin)},
					{Text: Branding.ColourizeBoolen(user.Reseller)},
					{Text: Branding.ColourizeBoolen(user.Vip)},
					{Text: Branding.ColourizeBoolen(user.Banned)},
					{Text: Branding.ColourizeBoolen(functions.MFA(user.MFA))},
					{Text: Branding.ColourizeBoolen(IsActive(user.PlanExpiry))},
					{Text: strconv.Itoa(user.Concurrents)},
				}
				
				table.Body.Cells = append(table.Body.Cells, r)
			}
		} else {
			for _, user := range Users {
				r := []*simpletable.Cell{
					{Text: strconv.Itoa(user.ID)},
					{Text: user.Username},
					{Text: strconv.FormatBool(user.Admin)},
					{Text: strconv.FormatBool(user.Reseller)},
					{Text: strconv.FormatBool(user.Vip)},
					{Text: strconv.FormatBool(user.Banned)},
					{Text: strconv.FormatBool(functions.MFA(user.MFA))},
					{Text: strconv.FormatBool(IsActive(user.PlanExpiry))},
					{Text: strconv.Itoa(user.Concurrents)},
				}

				table.Body.Cells = append(table.Body.Cells, r)
			}
		}
	}

	if !License.LiveWire {
		for _, user := range Users {
			r := []*simpletable.Cell{
				{Text: strconv.Itoa(user.ID)},
				{Text: user.Username},
				{Text: strconv.FormatBool(user.Admin)},
				{Text: strconv.FormatBool(user.Reseller)},
				{Text: strconv.FormatBool(user.Vip)},
				{Text: strconv.FormatBool(user.Banned)},
				{Text: strconv.FormatBool(IsActive(user.PlanExpiry))},
				{Text: strconv.Itoa(user.Concurrents)},
			}

			table.Body.Cells = append(table.Body.Cells, r)
		}
	}


	if len(table.Body.Cells) > 0 {
		table.SetStyle(simpletable.StyleCompact)

		if License.LiveWire {
			if JsonParse.LiveWireDLCSync.TableGradient.Status {
				fmt.Fprint(channel, "")
				livewire.Fade(strings.ReplaceAll(table.String(), "\n", "\r\n"), channel)
				fmt.Fprint(channel, "\r\n")
				return
			}
		}
		fmt.Fprint(channel, "")
		fmt.Fprintln(channel, strings.ReplaceAll(table.String(), "\n", "\r\n"))
		fmt.Fprint(channel, "\r")

	}

}

func IsActive(PlanEnd int64) bool {
	if PlanEnd < time.Now().Unix() {
		return false
	}

	return true
}