package commands

import (
	JsonParse "Yami/core/models/Json"
	Sessions "Yami/core/clients/session"
	YamiDB "Yami/core/db"
	"Yami/core/models/dlcs/Live_Wire"
	License "Yami/core/models/license"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/alexeyco/simpletable"
	"golang.org/x/crypto/ssh"
)

// branding reload command
func init() {
	Register(&Command{
		Name: "ongoing",
		Admin: false,
		Reseller: false,

		Descriptions: "Lists all current running attacks!",
		Execute: func(channel ssh.Channel, sshConn *ssh.ServerConn, Session *Sessions.Session, cmd []string) error {

			if Session.User.Admin {
				table := simpletable.New()
				table.Header = &simpletable.Header{
					Cells: []*simpletable.Cell{
						{Align: simpletable.AlignCenter, Text: "#"},
						{Align: simpletable.AlignCenter, Text: "Target"},
						{Align: simpletable.AlignCenter, Text: "Method"},
						{Align: simpletable.AlignCenter, Text: "Type"},
						{Align: simpletable.AlignCenter, Text: "Port"},
						{Align: simpletable.AlignCenter, Text: "Length"},
						{Align: simpletable.AlignCenter, Text: "Finish"},
						{Align: simpletable.AlignCenter, Text: "User"},
					},
				}
				
				Attacks, _ := YamiDB.Ongoing(sshConn.User())
	
				count := 0
				for _, s := range Attacks {
					lol,_ := strconv.ParseInt(strconv.Itoa(int(s.End)), 10, 64)
					TimeToWait := time.Unix(lol, 0)
					count++
					r := []*simpletable.Cell{
						{Align: simpletable.AlignCenter, Text: strconv.Itoa(count)},
						{Align: simpletable.AlignCenter, Text: s.Target},
						{Align: simpletable.AlignCenter, Text: s.Method},
						{Align: simpletable.AlignCenter, Text: s.Type},
						{Align: simpletable.AlignCenter, Text: strconv.Itoa(s.Port)},
						{Align: simpletable.AlignCenter, Text: strconv.Itoa(s.Duration)},
						{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%.0f secs", time.Until(TimeToWait).Seconds())},
						{Align: simpletable.AlignCenter, Text: s.Username},
					}
			
					table.Body.Cells = append(table.Body.Cells, r)
				}
	
				if len(table.Body.Cells) == 0 {
	
					if License.LiveWire {
						if JsonParse.LiveWireDLCSync.TableGradient.Status {
							livewire.Fade("There are no running attacks currently\r\n", channel)
							return nil
						}
					}
					channel.Write([]byte("There are no running attacks currently\r\n"))
					return nil
				}
	
				if len(table.Body.Cells) > 0 {
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
			
				}
			} else {
				table := simpletable.New()
				table.Header = &simpletable.Header{
					Cells: []*simpletable.Cell{
						{Align: simpletable.AlignCenter, Text: "#"},
						{Align: simpletable.AlignCenter, Text: "Target"},
						{Align: simpletable.AlignCenter, Text: "Method"},
						{Align: simpletable.AlignCenter, Text: "Type"},
						{Align: simpletable.AlignCenter, Text: "Port"},
						{Align: simpletable.AlignCenter, Text: "Length"},
						{Align: simpletable.AlignCenter, Text: "Finish"},
					},
				}
	
				Attacks, _ := YamiDB.Ongoing(sshConn.User())
	
				count := 0
				for _, s := range Attacks {
					lol,_ := strconv.ParseInt(strconv.Itoa(int(s.End)), 10, 64)
					TimeToWait := time.Unix(lol, 0)
					count++
					r := []*simpletable.Cell{
						{Align: simpletable.AlignCenter, Text: strconv.Itoa(count)},
						{Align: simpletable.AlignCenter, Text: s.Target},
						{Align: simpletable.AlignCenter, Text: s.Method},
						{Align: simpletable.AlignCenter, Text: s.Type},
						{Align: simpletable.AlignCenter, Text: strconv.Itoa(s.Port)},
						{Align: simpletable.AlignCenter, Text: strconv.Itoa(s.Duration)},
						{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%.0f secs", time.Until(TimeToWait).Seconds())},
					}
			
					table.Body.Cells = append(table.Body.Cells, r)
				}
	
				if len(table.Body.Cells) == 0 {
	
					if License.LiveWire {
						if JsonParse.LiveWireDLCSync.TableGradient.Status {
							livewire.Fade("There are no running attacks currently\r\n", channel)
							return nil
						}
					}
					channel.Write([]byte("There are no running attacks currently\r\n"))
					return nil
				}
	
				if len(table.Body.Cells) > 0 {
					table.SetStyle(simpletable.StyleCompact)
	
					if License.LiveWire {
						if JsonParse.LiveWireDLCSync.TableGradient.Status {
							fmt.Fprint(channel, "")
							livewire.Fade(strings.ReplaceAll(table.String(), "\n", "\r\n"), channel)
							fmt.Fprintln(channel, "\r\n")
							return nil
						}
					}
					fmt.Fprint(channel, "")
					fmt.Fprintln(channel, strings.ReplaceAll(table.String(), "\n", "\r\n"))
			
				}
			}
		
			return nil
		},
	})
}