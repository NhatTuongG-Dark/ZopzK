package commands

import (
	Branding "Yami/core/clients/branding"
	ClientPoly "Yami/core/clients/client"
	Sessions "Yami/core/clients/session"
	JsonParse "Yami/core/models/Json"
	BinLoad "Yami/core/models/bin"
	livewire "Yami/core/models/dlcs/Live_Wire"
	License "Yami/core/models/license"
	"fmt"
	"strconv"
	"strings"

	"github.com/alexeyco/simpletable"
	"golang.org/x/crypto/ssh"
)

// clear command
func init() {
	Register(&Command{
		Name: "commands",
		Admin: false,
		Reseller: false,

		Descriptions: "Lists all registered commands!",
		Execute: func(channel ssh.Channel, sshConn *ssh.ServerConn, Session *Sessions.Session, cmd []string) error {
			
			table := simpletable.New()
			table.Header = &simpletable.Header{
				Cells: []*simpletable.Cell{
					{Align: simpletable.AlignCenter, Text: "Name"},
					{Align: simpletable.AlignCenter, Text: "Description"},
					{Align: simpletable.AlignCenter, Text: "Admin"},
					{Align: simpletable.AlignCenter, Text: "Reseller"},
				},
			}

			if License.LiveWire {
				if JsonParse.LiveWireDLCSync.TableGradient.Status {
					for _, C := range commands {

						Blacklisted := Check(C.Name); if !Blacklisted {
							r := []*simpletable.Cell{
								{Align: simpletable.AlignCenter, Text: C.Name},
								{Align: simpletable.AlignCenter, Text: C.Descriptions},
								{Align: simpletable.AlignCenter, Text: strconv.FormatBool(C.Admin)},
								{Align: simpletable.AlignCenter, Text: strconv.FormatBool(C.Reseller)},
							}
					
							table.Body.Cells = append(table.Body.Cells, r)
						}
					}

					for _, C := range ClientPoly.BetaMapHandler {


						Blacklisted := Check(C.CommandName); if !Blacklisted {
							r := []*simpletable.Cell{
								{Align: simpletable.AlignCenter, Text: C.CommandName},
								{Align: simpletable.AlignCenter, Text: C.CommandDescription},
								{Align: simpletable.AlignCenter, Text: strconv.FormatBool(C.CommandAdmin)},
								{Align: simpletable.AlignCenter, Text: strconv.FormatBool(C.CommandReseller)},
							}
					
							table.Body.Cells = append(table.Body.Cells, r)
						}
					}

					for _, C := range BinLoad.BetaMapHandler {


						Blacklisted := Check(C.CommandName); if !Blacklisted {
							r := []*simpletable.Cell{
								{Align: simpletable.AlignCenter, Text: C.CommandName},
								{Align: simpletable.AlignCenter, Text: C.CommandDescription},
								{Align: simpletable.AlignCenter, Text: strconv.FormatBool(C.CommandAdmin)},
								{Align: simpletable.AlignCenter, Text: strconv.FormatBool(C.CommandReseller)},
							}
					
							table.Body.Cells = append(table.Body.Cells, r)
						}
					}
				} else {
					for _, C := range commands {


						Blacklisted := Check(C.Name); if !Blacklisted {
							r := []*simpletable.Cell{
								{Align: simpletable.AlignCenter, Text: C.Name},
								{Align: simpletable.AlignCenter, Text: C.Descriptions},
								{Align: simpletable.AlignCenter, Text: Branding.ColourizeBoolen(C.Admin)},
								{Align: simpletable.AlignCenter, Text: Branding.ColourizeBoolen(C.Reseller)},
							}
					
							table.Body.Cells = append(table.Body.Cells, r)
						}
					}

					for _, C := range ClientPoly.BetaMapHandler {

						Blacklisted := Check(C.CommandName); if !Blacklisted {
							r := []*simpletable.Cell{
								{Align: simpletable.AlignCenter, Text: C.CommandName},
								{Align: simpletable.AlignCenter, Text: C.CommandDescription},
								{Align: simpletable.AlignCenter, Text: Branding.ColourizeBoolen(C.CommandAdmin)},
								{Align: simpletable.AlignCenter, Text: Branding.ColourizeBoolen(C.CommandReseller)},
							}
					
							table.Body.Cells = append(table.Body.Cells, r)
						}
					}

					for _, C := range BinLoad.BetaMapHandler {

						Blacklisted := Check(C.CommandName); if !Blacklisted {
							r := []*simpletable.Cell{
								{Align: simpletable.AlignCenter, Text: C.CommandName},
								{Align: simpletable.AlignCenter, Text: C.CommandDescription},
								{Align: simpletable.AlignCenter, Text: Branding.ColourizeBoolen(C.CommandAdmin)},
								{Align: simpletable.AlignCenter, Text: Branding.ColourizeBoolen(C.CommandReseller)},
							}
					
							table.Body.Cells = append(table.Body.Cells, r)
						}
					}
				}
			} else {
				for _, C := range commands {


					Blacklisted := Check(C.Name); if !Blacklisted {
						r := []*simpletable.Cell{
							{Align: simpletable.AlignCenter, Text: C.Name},
							{Align: simpletable.AlignCenter, Text: C.Descriptions},
							{Align: simpletable.AlignCenter, Text: Branding.ColourizeBoolen(C.Admin)},
							{Align: simpletable.AlignCenter, Text: Branding.ColourizeBoolen(C.Reseller)},
						}
				
						table.Body.Cells = append(table.Body.Cells, r)
					}
				}

				for _, C := range ClientPoly.BetaMapHandler {


					Blacklisted := Check(C.CommandName); if !Blacklisted {
						r := []*simpletable.Cell{
							{Align: simpletable.AlignCenter, Text: C.CommandName},
							{Align: simpletable.AlignCenter, Text: C.CommandDescription},
							{Align: simpletable.AlignCenter, Text: Branding.ColourizeBoolen(C.CommandAdmin)},
							{Align: simpletable.AlignCenter, Text: Branding.ColourizeBoolen(C.CommandReseller)},
						}
				
						table.Body.Cells = append(table.Body.Cells, r)
					}
				}

				for _, C := range BinLoad.BetaMapHandler {


					Blacklisted := Check(C.CommandName); if !Blacklisted {
						r := []*simpletable.Cell{
							{Align: simpletable.AlignCenter, Text: C.CommandName},
							{Align: simpletable.AlignCenter, Text: C.CommandDescription},
							{Align: simpletable.AlignCenter, Text: Branding.ColourizeBoolen(C.CommandAdmin)},
							{Align: simpletable.AlignCenter, Text: Branding.ColourizeBoolen(C.CommandReseller)},
						}
				
						table.Body.Cells = append(table.Body.Cells, r)
					}
				}
			}

			table.SetStyle(simpletable.StyleCompact)

			if License.LiveWire {
				if JsonParse.LiveWireDLCSync.TableGradient.Status {
					fmt.Fprint(channel, "")
					livewire.Fade(strings.ReplaceAll(table.String(), "\n", "\r\n"), channel)
					fmt.Fprintln(channel, "\r")
					return nil
				}
			}
			fmt.Fprint(channel, "")
			fmt.Fprintln(channel, strings.ReplaceAll(table.String(), "\n", "\r\n"))
			fmt.Fprint(channel, "\r")


			return nil
		},
	})
}