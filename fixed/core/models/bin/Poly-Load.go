package BinLoad

import (
	"io/ioutil"
	"log"
	"strconv"
	"sync"

	"github.com/vaughan0/go-ini"
	"golang.org/x/crypto/ssh"
)

var CommandFile string = "build/commands/bin/"

var (
	BetaMapHandler = make(map[string]*CommandBin)
	Handle sync.Mutex
)

type CommandBin struct {
	CommandName		string
	CommandAdmin   	bool
	CommandReseller     bool
	CommandVip		bool
	Execute     		string
	CommandDescription  string
}


func Load(channel ssh.Channel) bool {
	for i, _ := range BetaMapHandler {
		delete(BetaMapHandler, i)
	}

	Files, err := ioutil.ReadDir(CommandFile); if err != nil {
		return false
	}
	for _, f := range Files {
		file, _ := ini.LoadFile(CommandFile+f.Name())

		CommandExecuteLoad, ok := file.Get("","execute"); if !ok {
			channel.Write([]byte("\x1b[0m[\x1b[38;5;1m FAILED \x1b[0m] Failed to get execute from \""+f.Name()+"\"\r\n"))
			continue
		}

		CommandNameLoad, ok := file.Get("","name"); if !ok {
			channel.Write([]byte("\x1b[0m[\x1b[38;5;1m FAILED \x1b[0m] Failed to get name from \""+f.Name()+"\"\r\n"))
			continue
		}

		CommandDescriptionLoad, ok := file.Get("","description"); if !ok {
			channel.Write([]byte("\x1b[0m[\x1b[38;5;1m FAILED \x1b[0m] Failed to get description from \""+f.Name()+"\"\r\n"))
			continue
		}

		CommandAdminLoad, ok := file.Get("","admin"); if !ok {
			channel.Write([]byte("\x1b[0m[\x1b[38;5;1m FAILED \x1b[0m] Failed to Admin boolen from \""+f.Name()+"\"\r\n"))
			continue
		}

		CommandResellerLoad, ok := file.Get("","reseller"); if !ok {
			channel.Write([]byte("\x1b[0m[\x1b[38;5;1m FAILED \x1b[0m] Failed to Reseller boolen from \""+f.Name()+"\"\r\n"))
			continue
		}

		CommandVipLoad, ok := file.Get("","vip"); if !ok {
			channel.Write([]byte("\x1b[0m[\x1b[38;5;1m FAILED \x1b[0m] Failed to vip boolen from \""+f.Name()+"\"\r\n"))
			continue
		}

		CommandAdminType, error := strconv.ParseBool(CommandAdminLoad); if error != nil {
			channel.Write([]byte("\x1b[0m[\x1b[38;5;1m FAILED \x1b[0m] Failed to Admin boolen from \""+f.Name()+"\"\r\n"))
			continue
		}

		CommandResellerType, error := strconv.ParseBool(CommandResellerLoad); if error != nil {
			channel.Write([]byte("\x1b[0m[\x1b[38;5;1m FAILED \x1b[0m] Failed to Reseller boolen from \""+f.Name()+"\"\r\n"))
			continue
		}

		CommandVipType, error := strconv.ParseBool(CommandVipLoad); if error != nil {
			channel.Write([]byte("\x1b[0m[\x1b[38;5;1m FAILED \x1b[0m] Failed to Vip boolen from \""+f.Name()+"\"\r\n"))
			continue
		}

		var MapCommand = CommandBin {
			CommandName: 		CommandNameLoad,
			CommandAdmin:       CommandAdminType,
			CommandReseller:    CommandResellerType,
			CommandVip:         CommandVipType,
			Execute:    		CommandExecuteLoad,
			CommandDescription: CommandDescriptionLoad,
		}

		Handle.Lock()
		BetaMapHandler[MapCommand.CommandName] = &MapCommand
		Handle.Unlock()

	}

	channel.Write([]byte("[ \x1b[38;5;2mOK\x1b[0m ] Reloaded "+strconv.Itoa(len(BetaMapHandler))+" Bin Commands Correctly\r\n"))

	log.Println(" [RELOADED] [Reloaded BIN Commands Correctly] ["+strconv.Itoa(len(BetaMapHandler))+"]")



	return true
}