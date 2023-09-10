package JsonParse

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)


var LiveWireDLCSync *LiveWireDLC



func LoadLiveWire() (bool, error) {
	File, errors := os.Open("build/livewire-DLC.json"); if errors != nil {
		log.Println("Failed To Parse \"livewire-DLC.json\"",errors)
  
		err := File.Close(); if err != nil {
		    log.Println("Failed Closing Of \"livewire-DLC.json\"")
		    return false,err
		}
		return false,errors
	}
	byteValue, errors := ioutil.ReadAll(File); if errors != nil {
		log.Println("Failed To Read \"livewire-DLC.json\"",errors)
  
		err := File.Close(); if err != nil {
		    log.Println("Failed Closing Of \"livewire-DLC.json\"")
		    return false,err
		}
		return false,errors
	}
	var CG LiveWireDLC
	errors = json.Unmarshal(byteValue, &CG)
	if errors != nil {
		log.Println("Failed To Parse \"livewire-DLC.json\"", errors)
  
		err := File.Close(); if err != nil {
		    log.Println("Failed Closing Of \"livewire-DLC.json\"")
		    return false,err
		}
		return false,errors
	}
  
	err := File.Close(); if err != nil {
		log.Println("Failed Closing Of \"livewire-DLC.json\"")
		return false,err
	}
  
  
	LiveWireDLCSync = &CG
	return true, nil
}
