package JsonParse

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)


var SlaveSync *Slaves



func LoadSlaves() (bool, error) {
	File, errors := os.Open("build/slaves.json"); if errors != nil {
		log.Println("Failed To Parse \"slaves.json\"",errors)
  
		err := File.Close(); if err != nil {
		    log.Println("Failed Closing Of \"slaves.json\"")
		    return false,err
		}
		return false,errors
	}
	byteValue, errors := ioutil.ReadAll(File); if errors != nil {
		log.Println("Failed To Read \"slaves.json\"",errors)
  
		err := File.Close(); if err != nil {
		    log.Println("Failed Closing Of \"slaves.json\"")
		    return false,err
		}
		return false,errors
	}
	var CG Slaves
	errors = json.Unmarshal(byteValue, &CG)
	if errors != nil {
		log.Println("Failed To Parse \"slaves.json\"", errors)
  
		err := File.Close(); if err != nil {
		    log.Println("Failed Closing Of \"slaves.json\"")
		    return false,err
		}
		return false,errors
	}
  
	err := File.Close(); if err != nil {
		log.Println("Failed Closing Of \"slaves.json\"")
		return false,err
	}

	SlaveSync = &CG
	return true, nil
}
