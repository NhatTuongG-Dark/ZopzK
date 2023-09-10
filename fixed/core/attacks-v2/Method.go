package AttacksV2

import (
	JsonParse "Yami/core/models/Json"
	"strings"
)

func MethodCheck(Name string) bool {
	for i := 0; i < len(JsonParse.AttackSyncs.Attacks); i++ {
		if Name == JsonParse.AttackSyncs.Attacks[i].Name {
			return true
		}
	}


	if !strings.Contains(Name, JsonParse.SlaveSync.AttackPrefix) {
		return false
	}
	
	for l := 0; l < len(JsonParse.SlaveSync.Slaves); l++ {
		Inputed := strings.ToLower(strings.Replace(Name, JsonParse.SlaveSync.AttackPrefix, "", -1))
		if Inputed == JsonParse.SlaveSync.Slaves[l].Name {
			return true
		}
	}

	return false
}