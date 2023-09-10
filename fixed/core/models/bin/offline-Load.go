package BinLoad

import (
	"io/ioutil"
	"log"
	"strconv"

	"github.com/vaughan0/go-ini"
)


func OfflineLoad() bool {
	for i, _ := range BetaMapHandler {
		delete(BetaMapHandler, i)
	}

	Files, err := ioutil.ReadDir(CommandFile); if err != nil {
		return false
	}
	for _, f := range Files {
		file, _ := ini.LoadFile(CommandFile+f.Name())

		CommandExecuteLoad, ok := file.Get("","execute"); if !ok {
			continue
		}

		CommandNameLoad, ok := file.Get("","name"); if !ok {
			continue
		}
		CommandDescriptionLoad, ok := file.Get("","description"); if !ok {
			continue
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

	log.Println(" [RELOADED] [Reloaded Bin Commands Correctly] ["+strconv.Itoa(len(BetaMapHandler))+"]")



	return true
}