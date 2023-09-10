package YamiSshServer

import (
	"Yami/core/models/Json"
	Options "Yami/core/models/config"
	"Yami/core/db"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

var SSHConfig *ssh.ServerConfig

func NewSSH() {
	config := &ssh.ServerConfig{

		BannerCallback: func(c ssh.ConnMetadata) string {
			return JsonParse.ConfigSyncs.Branding.AppName
		},
		PasswordCallback: func(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {

			error := YamiDB.Auth(string(c.User()), string(pass)); if !error {
				log.Printf("Invaild Login Details Provided From %s For %s", c.RemoteAddr().String(), c.User())
				return nil, fmt.Errorf("Invaild Login Details Inputed")
			} else {
				return nil, nil
			}
		},
		ServerVersion: "SSH-2.0-OpenSSH_8.6p1 "+JsonParse.ConfigSyncs.Branding.AppName+" "+Options.ClientVersion,
	}

	config.MaxAuthTries = JsonParse.ConfigSyncs.Masters.MaxAuthTries

	keyBytes, err := rsa.GenerateKey(rand.Reader, 2048); if err != nil {
		return
	}
	key, err := ssh.NewSignerFromSigner(keyBytes); if err != nil {
		return
	}
	config.AddHostKey(key)

	SSHConfig = config
	Listen()
}