package Fake

import (
	"Yami/core/models/Json"
	"math/rand"
)

func MakeLoop() int {

	Num := rand.Intn(JsonParse.Option.FakeSlaves.MaxGen-JsonParse.Option.FakeSlaves.MinGen) + JsonParse.Option.FakeSlaves.MinGen

	return Num
}