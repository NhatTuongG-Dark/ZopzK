package BinLoad

import (

)

func Command(name string) *CommandBin {
	for _, s := range BetaMapHandler {
		if s.CommandName == name {
			return s
		}
	}
	return nil
}