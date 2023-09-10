package ClientPoly

import (
	"bufio"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/vaughan0/go-ini"
)




func OfflineLoader() bool {

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
			continue
		}

		CommandDescriptionLoad, ok := file.Get("","description"); if !ok {

		}

		CommandAdminLoad, ok := file.Get("","admin"); if !ok {
			continue
		}

		CommandResellerLoad, ok := file.Get("","reseller"); if !ok {
			continue
		}

		CommandVipLoad, ok := file.Get("","vip"); if !ok {
			continue
		}

		CommandAdminType, error := strconv.ParseBool(CommandAdminLoad); if error != nil {
			continue
		}

		CommandResellerType, error := strconv.ParseBool(CommandResellerLoad); if error != nil {
			continue
		}

		CommandVipType, error := strconv.ParseBool(CommandVipLoad); if error != nil {
			continue
		}

		Files, err := ioutil.ReadFile(CommandFile+f.Name()); if err != nil {
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

		var CommandInfo = &CommandText {
			CommandName: 		CommandNameLoad,
			CommandAdmin:       CommandAdminType,
			CommandReseller:    CommandResellerType,
			CommandVip:         CommandVipType,
			CommandContains:    Banner,
			CommandDescription: CommandDescriptionLoad,
		}

		Handle.Lock()
		BetaMapHandler[CommandInfo.CommandName] = CommandInfo
		Handle.Unlock()

		loaded++

	}

	log.Printf(" [Loaded %d Custom Commands Correctly]", loaded)
	return true
}