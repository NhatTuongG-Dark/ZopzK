package Branding

import (
	"io/ioutil"
	"log"
	"strings"
	"sync"
	"time"
)

var (
	BrandingWorker = make(map[int64]*Branding)
	Commands sync.Mutex
)

type Branding struct {
	ID				int64
	CommandName		string
	CommandFile		string
	CommandContains	string
}


var BrandingFile = "build/branding/"

func CompleteLoad() (int, error) {
	for _, f := range BrandingWorker {
		delete(BrandingWorker, f.ID)
	}


	files, err := ioutil.ReadDir(BrandingFile); if err != nil {
		log.Println(err)
		return 0, err
	}

	loaded := 0

	for _, f := range files {
		CommandVisual, err := ioutil.ReadFile(BrandingFile+f.Name()); if err != nil {
			log.Printf("Failed to grab branding from \"%s\", Reason: %s.", f.Name(), err.Error())
			continue
		}

		if !strings.Contains(f.Name(), ".tfx") {
			log.Printf("Branding file is a invaild file type, Is \"%s\", needed \".tfx\"", f.Name())
			continue
		}

		var Working = &Branding{
			ID:      time.Now().UnixNano(),
			CommandName: strings.Replace(f.Name(), ".tfx", "", -1),
			CommandFile: f.Name(),
			CommandContains: string(CommandVisual),
		}

		Commands.Lock()
		BrandingWorker[Working.ID] = Working
		Commands.Unlock()

		loaded++
	}

	return loaded, nil
}