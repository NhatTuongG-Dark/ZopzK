package JsonParse

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)


var ConfigSyncs *ConfigSync



func LoadConfig() (bool, error) {
	File, errors := os.Open("build/config.json"); if errors != nil {
		log.Println("Failed To Parse \"config.json\"",errors)
  
		err := File.Close(); if err != nil {
		    log.Println("Failed Closing Of \"config.json\"")
		    return false,err
		}
		return false,errors
	}
	byteValue, errors := ioutil.ReadAll(File); if errors != nil {
		log.Println("Failed To Read \"config.json\"",errors)
  
		err := File.Close(); if err != nil {
		    log.Println("Failed Closing Of \"config.json\"")
		    return false,err
		}
		return false,errors
	}
	var CG ConfigSync
	errors = json.Unmarshal(byteValue, &CG)
	if errors != nil {
		log.Println("Failed To Parse \"config.json\"", errors)
  
		err := File.Close(); if err != nil {
		    log.Println("Failed Closing Of \"config.json\"")
		    return false,err
		}
		return false,errors
	}
  
	err := File.Close(); if err != nil {
		log.Println("Failed Closing Of \"config.json\"")
		return false,err
	}
  
  
	ConfigSyncs = &CG
	return true, nil
}
