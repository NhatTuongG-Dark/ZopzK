package Slaves

import (
	JsonParse "Yami/core/models/Json"
	slavesFake "Yami/core/slaves/Fake"
	slavesMirai "Yami/core/slaves/mirai"
	"Yami/core/slaves/transition"
	"strconv"
)

func Get() string {
	if JsonParse.ConfigSyncs.Slaves.Status {
		return strconv.Itoa(slavesMirai.Count())
	} else if JsonParse.Option.SlaveTransition.Status {
		return transition.Count
	} else if JsonParse.Option.FakeSlaves.Status {
		return strconv.Itoa(slavesFake.MakeLoop())
	} else {
		return "0"
	}
}