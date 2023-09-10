package Masters

import (
	"Yami/core/models/Json"
	Branding "Yami/core/clients/branding"
	"Yami/core/db"
	"Yami/core/models/license"
	"time"

	"golang.org/x/crypto/ssh"
)

var Testing bool

func Title(channel ssh.Channel, conn *ssh.ServerConn) {

	// live wire title worker!

	for {
		if License.LiveWire {
			if JsonParse.LiveWireDLCSync.TitleSpinner.Active {
				for i := 0; i < len(JsonParse.LiveWireDLCSync.TitleSpinner.Frames); i++ {
					title, err := Branding.TermFXExecuteTitle("title", conn)
					if err != nil {
						channel.Write([]byte(string("\033]0; EOF: title.tfx\007")))
						return
					}


				
					if JsonParse.LiveWireDLCSync.TitleSpinner.Position == "front" {
						channel.Write([]byte(string("\033]0; [" + JsonParse.LiveWireDLCSync.TitleSpinner.Frames [i] + "]" + title + "\007")))
					} else if JsonParse.LiveWireDLCSync.TitleSpinner.Position == "back" {
						channel.Write([]byte(string("\033]0; " + title + "[" + JsonParse.LiveWireDLCSync.TitleSpinner.Frames [i] + "]\007")))

					} else if JsonParse.LiveWireDLCSync.TitleSpinner.Position == "both" {
						channel.Write([]byte(string("\033]0; [" + JsonParse.LiveWireDLCSync.TitleSpinner.Frames [i] + "]" + title + "[" + JsonParse.LiveWireDLCSync.TitleSpinner.Frames [i] + "]\007")))
					} else {

						Users, _ := YamiDB.GetUser(conn.User())

						if Users.Admin {
							channel.Write([]byte("\033]0; Invaild Position : \"front\", \"back\", \"both\"\007"))
						} else {
							channel.Write([]byte(string("\033]0;" + title + "\007")))
						}
					}
	
					time.Sleep(1 * time.Second)
				}
			} else {
				title, err := Branding.TermFXExecuteTitle("title", conn)
				if err != nil {
					channel.Write([]byte(string("\033]0; EOF: title.tfx\007")))
					return
				}
		
				channel.Write([]byte(string("\033]0;" + title + "\007")))
		
				time.Sleep(1 * time.Second)
			}
		} else {
			title, err := Branding.TermFXExecuteTitle("title", conn)
			if err != nil {
				channel.Write([]byte(string("\033]0; EOF: title.tfx\007")))
				return
			}
	
			channel.Write([]byte(string("\033]0;" + title + "\007")))
	
			time.Sleep(1 * time.Second)
		}
	}

}
