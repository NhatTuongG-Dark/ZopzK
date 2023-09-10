package SessionsCommands

import (
	JsonParse "Yami/core/models/Json"
	Sessions "Yami/core/clients/session"
	livewire "Yami/core/models/dlcs/Live_Wire"
	License "Yami/core/models/license"
	"fmt"
	"strconv"
	"strings"

	"github.com/alexeyco/simpletable"
	"golang.org/x/crypto/ssh"
)

func ListIds(channel ssh.Channel, sshConn *ssh.ServerConn) {
		table := simpletable.New()
		table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignCenter, Text: "Username"},
		},
	}

	if License.LiveWire {
		if !JsonParse.LiveWireDLCSync.TableGradient.Status {
			for _, s := range Sessions.Sessions {
				r := []*simpletable.Cell{
					{Align: simpletable.AlignCenter, Text: strconv.Itoa(int(s.ID))},
					{Align: simpletable.AlignCenter, Text: s.User.Username},

				}
				
				table.Body.Cells = append(table.Body.Cells, r)
			}
		} else {
			for _, s := range Sessions.Sessions {
				r := []*simpletable.Cell{
					{Align: simpletable.AlignCenter, Text: strconv.Itoa(int(s.ID))},
					{Align: simpletable.AlignCenter, Text: s.User.Username},
				}
			
				table.Body.Cells = append(table.Body.Cells, r)
			}
		}
	} else {
		for _, s := range Sessions.Sessions {
			r := []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: strconv.Itoa(int(s.ID))},
				{Align: simpletable.AlignCenter, Text: s.User.Username},
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
		fmt.Fprintln(channel, strings.ReplaceAll(table.String(), "\n", "\r\n"))
		fmt.Fprint(channel, "\r")
		return

	}
	return
}