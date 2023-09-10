package JsonParse

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)


var Option *Options



func LoadOptions() (bool, error) {
	File, errors := os.Open("build/options.json"); if errors != nil {
		log.Println("Failed To Parse \"options.json\"",errors)
  
		err := File.Close(); if err != nil {
		    log.Println("Failed Closing Of \"options.json\"")
		    return false,err
		}
		return false,errors
	}
	byteValue, errors := ioutil.ReadAll(File); if errors != nil {
		log.Println("Failed To Read \"options.json\"",errors)
  
		err := File.Close(); if err != nil {
		    log.Println("Failed Closing Of \"options.json\"")
		    return false,err
		}
		return false,errors
	}
	var CG Options
	errors = json.Unmarshal(byteValue, &CG)
	if errors != nil {
		log.Println("Failed To Parse \"options.json\"", errors)
  
		err := File.Close(); if err != nil {
		    log.Println("Failed Closing Of \"options.json\"")
		    return false,err
		}
		return false,errors
	}
  
	err := File.Close(); if err != nil {
		log.Println("Failed Closing Of \"options.json\"")
		return false,err
	}
  
  
	Option = &CG
	return true, nil
}
