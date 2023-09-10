package Sessions

import (
	YamiDB "Yami/core/db"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"
)

var (
	Sessions     = make(map[int64]*Session)
	SessionMutex sync.Mutex
)

type Session struct {
	ID   	int64
	User 	*YamiDB.User

	Channel   ssh.Channel
	Conn 	ssh.ServerConn

	Theme	string
	OpenAt	time.Time
	Chat    bool

	SpamLimiter int
}