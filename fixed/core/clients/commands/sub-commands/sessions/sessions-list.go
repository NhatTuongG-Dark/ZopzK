package SessionsCommands

import (
	JsonParse "Yami/core/models/Json"
	"Yami/core/clients/branding"
	Sessions "Yami/core/clients/session"
	livewire "Yami/core/models/dlcs/Live_Wire"
	License "Yami/core/models/license"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/alexeyco/simpletable"
	"golang.org/x/crypto/ssh"
)

func ListSessions(channel ssh.Channel, sshConn *ssh.ServerConn) {
		table := simpletable.New()
		table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Username"},
			{Align: simpletable.AlignCenter, Text: "Admin"},
			{Align: simpletable.AlignCenter, Text: "Reseller"},
			{Align: simpletable.AlignCenter, Text: "Vip"},
			{Align: simpletable.AlignCenter, Text: "IPv4"},
			{Align: simpletable.AlignCenter, Text: "Uptime"},
		},
	}

	if License.LiveWire {
		if !JsonParse.LiveWireDLCSync.TableGradient.Status {
			for _, s := range Sessions.Sessions {
				r := []*simpletable.Cell{
					{Align: simpletable.AlignCenter, Text: strconv.Itoa(s.User.ID)},
					{Align: simpletable.AlignCenter, Text: s.User.Username},
					{Align: simpletable.AlignCenter, Text: Branding.ColourizeBoolen(s.User.Admin)},
					{Align: simpletable.AlignCenter, Text: Branding.ColourizeBoolen(s.User.Reseller)},
					{Align: simpletable.AlignCenter, Text: Branding.ColourizeBoolen(s.User.Vip)},
					{Align: simpletable.AlignCenter, Text: s.Conn.RemoteAddr().String()},
					{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%.2f mins", time.Since(s.OpenAt).Minutes())},
				}
				
				table.Body.Cells = append(table.Body.Cells, r)
			}
		} else {
			for _, s := range Sessions.Sessions {
				r := []*simpletable.Cell{
					{Align: simpletable.AlignCenter, Text: strconv.Itoa(s.User.ID)},
					{Align: simpletable.AlignCenter, Text: s.User.Username},
					{Align: simpletable.AlignCenter, Text: strconv.FormatBool(s.User.Admin)},
					{Align: simpletable.AlignCenter, Text: strconv.FormatBool(s.User.Reseller)},
					{Align: simpletable.AlignCenter, Text: strconv.FormatBool(s.User.Vip)},
					{Align: simpletable.AlignCenter, Text: s.Conn.RemoteAddr().String()},
					{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%.2f mins", time.Since(s.OpenAt).Minutes())},
				}
			
				table.Body.Cells = append(table.Body.Cells, r)
			}
		}
	} else {
		for _, s := range Sessions.Sessions {
			r := []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: strconv.Itoa(s.User.ID)},
				{Align: simpletable.AlignCenter, Text: s.User.Username},
				{Align: simpletable.AlignCenter, Text: Branding.ColourizeBoolen(s.User.Admin)},
				{Align: simpletable.AlignCenter, Text: Branding.ColourizeBoolen(s.User.Reseller)},
				{Align: simpletable.AlignCenter, Text: Branding.ColourizeBoolen(s.User.Vip)},
				{Align: simpletable.AlignCenter, Text: s.Conn.RemoteAddr().String()},
				{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%.2f mins", time.Since(s.OpenAt).Minutes())},
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
				fmt.Fprint(channel, "\r\n")
				return
			}
		}
		table.SetStyle(simpletable.StyleCompact)
		fmt.Fprint(channel, "")
		fmt.Fprint(channel, strings.ReplaceAll(table.String(), "\n", "\r\n"))
		fmt.Fprint(channel, "\r\n")
		return

	}
	return
}