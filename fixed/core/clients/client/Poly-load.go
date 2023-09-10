package ClientPoly

import (
	"bufio"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/vaughan0/go-ini"
	"golang.org/x/crypto/ssh"
)

var (
	BetaMapHandler = make(map[string]*CommandText)
	Handle sync.Mutex
)

type CommandText struct {
	CommandName		string
	CommandAdmin   	bool
	CommandReseller     bool
	CommandVip		bool
	CommandContains     string
	CommandDescription  string
}

var CommandFile string = "build/commands/assets/"

func PolyLoader(channel ssh.Channel) bool {

	for i, _ := range BetaMapHandler {
		delete(BetaMapHandler, i)
	}

	Files, err := ioutil.ReadDir(CommandFile); if err != nil {
		return false
	}
	loaded := 0

	for _, f := range Files {
		file, _ := ini.LoadFile(CommandFile+f.Name())

		CommandNameLoad, ok := file.Get("","name"); if !ok {
			channel.Write([]byte("\x1b[0m[\x1b[38;5;1m FAILED \x1b[0m] Failed to get name from \""+f.Name()+"\"\r\n"))
			continue
		}

		CommandDescriptionLoad, ok := file.Get("","description"); if !ok {
			channel.Write([]byte("\x1b[0m[\x1b[38;5;1m FAILED \x1b[0m] Failed to get Description from \""+f.Name()+"\"\r\n"))
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

		Files, err := ioutil.ReadFile(CommandFile+f.Name()); if err != nil {
			channel.Write([]byte("\x1b[0m[\x1b[38;5;1m FAILED \x1b[0m] Failed to load Texture from \""+f.Name()+"\"\r\n"))
			continue
		}

		var Banner string = ""
		var FoundEnd bool = false
		scan := bufio.NewScanner(strings.NewReader(string(Files)))
		for scan.Scan() {

			if FoundEnd != true {
				if strings.Contains(scan.Text(), "================== END ==================") {
					FoundEnd = true
					continue
				}
			} else {
				if Banner == "" {
					Banner = scan.Text() + "\n"
					continue
				}
				Banner = Banner + scan.Text() + "\n"
			}
		}

		var MapCommand = CommandText {
			CommandName: 		CommandNameLoad,
			CommandAdmin:       CommandAdminType,
			CommandReseller:    CommandResellerType,
			CommandVip:         CommandVipType,
			CommandContains:    Banner,
			CommandDescription: CommandDescriptionLoad,
		}

		Handle.Lock()
		BetaMapHandler[MapCommand.CommandName] = &MapCommand
		Handle.Unlock()


		loaded++
	}

	channel.Write([]byte("[ \x1b[38;5;2mOK\x1b[0m ] Reloaded "+strconv.Itoa(len(BetaMapHandler))+" Commands Correctly\r\n"))

	log.Println(" [RELOADED] [Reloaded Commands Correctly] ["+strconv.Itoa(len(BetaMapHandler))+"]")


	return true
}