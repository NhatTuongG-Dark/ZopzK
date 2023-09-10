package YamiView

import (
	"Yami/core/models/Json"
	"Yami/core/clients/branding"
	"log"
	"math/rand"
	"strconv"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

// catpcha screen for clients
func Catpcha(channel ssh.Channel, conn *ssh.ServerConn, oldrequests <-chan *ssh.Request) string {
	var Used int

	fromthetop:

	if JsonParse.ConfigSyncs.Controls.Catpcha.AllowedAttempts == Used {
		channel.Write([]byte("You have answered too many wrong questions!"))
		time.Sleep(10 * time.Second)
		channel.Close()
		return "F"
	}

	NumOne := rand.Intn(JsonParse.ConfigSyncs.Controls.Catpcha.Question.MaxGen - JsonParse.ConfigSyncs.Controls.Catpcha.Question.MinGen) + JsonParse.ConfigSyncs.Controls.Catpcha.Question.MinGen
	NumTwo := rand.Intn(JsonParse.ConfigSyncs.Controls.Catpcha.Question.MaxGen - JsonParse.ConfigSyncs.Controls.Catpcha.Question.MinGen) + JsonParse.ConfigSyncs.Controls.Catpcha.Question.MinGen
	
	
	error := Branding.TermFXExecuteBannerCustom("catpcha-banner", conn, channel, "question", strconv.Itoa(NumOne)+" + "+strconv.Itoa(NumTwo)); if error != nil {
		channel.Write([]byte("Answer This Question to gain access -> "+strconv.Itoa(NumOne) + " + " + strconv.Itoa(NumTwo)))
	}

	Prompt,error := Branding.TermFXExecutePrompt("catpcha-prompt", conn); if error != nil {
		log.Println("Failed to load prompt for catpcha!")
	}

	Answer := terminal.NewTerminal(channel, "")

	/*go func() { // oob requests lmao
		for req := range oldrequests {
			switch req.Type {
			case "pty-req":
				termLen := req.Payload[3]
				w, h := ParseDims(req.Payload[termLen+4:])
				Answer.SetSize(int(w), int(h))
				req.Reply(true, nil)
			case "window-change":
				w, h := ParseDims(req.Payload)
				Answer.SetSize(int(w), int(h))
			}
		}
	}()*/

	if Prompt == "" {
		channel.Write([]byte(">"))
	} else {
		channel.Write([]byte(Prompt))
	}

	RAnswer,err := Answer.ReadLine(); if err != nil {
		channel.Write([]byte("\r\nVoiding\r\n"))
		channel.Close()
		return "F"
	}

	AnswerQ := NumOne + NumTwo

	if RAnswer != strconv.Itoa(AnswerQ) {
		Used++
		goto fromthetop
	} else {
		return ""
	}


}