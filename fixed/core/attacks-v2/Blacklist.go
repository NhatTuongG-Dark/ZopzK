package AttacksV2

import (
	JsonParse "Yami/core/models/Json"
	"strings"
)

func CheckBlacklist(Addr string) bool {
	for i := 0; i < len(JsonParse.ConfigSyncs.Attacks.Blacklists); i++ {
		if strings.ToLower(Addr) == JsonParse.ConfigSyncs.Attacks.Blacklists[i] {
			return true
		}
	}

	return false
}