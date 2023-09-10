package JsonParse

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)


var AttackSyncs *AttackSync



func LoadAttacks() (bool, error) {
	File, errors := os.Open("build/attack.json"); if errors != nil {
		log.Println("Failed To Parse \"attack.json\"",errors)
  
		err := File.Close(); if err != nil {
		    log.Println("Failed Closing Of \"attack.json\"")
		    return false,err
		}
		return false,errors
	}
	byteValue, errors := ioutil.ReadAll(File); if errors != nil {
		log.Println("Failed To Read \"attack.json\"",errors)
  
		err := File.Close(); if err != nil {
		    log.Println("Failed Closing Of \"attack.json\"")
		    return false,err
		}
		return false,errors
	}
	var CG AttackSync
	errors = json.Unmarshal(byteValue, &CG)
	if errors != nil {
		log.Println("Failed To Parse \"attack.json\"", errors)
  
		err := File.Close(); if err != nil {
		    log.Println("Failed Closing Of \"attack.json\"")
		    return false,err
		}
		return false,errors
	}
  
	err := File.Close(); if err != nil {
		log.Println("Failed Closing Of \"attack.json\"")
		return false,err
	}
  
  
	AttackSyncs = &CG
	return true, nil
}
