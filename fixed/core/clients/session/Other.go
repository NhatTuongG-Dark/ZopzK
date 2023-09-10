package Sessions

import (

	"golang.org/x/crypto/ssh"
)


func (s *Session) Check(conn *ssh.ServerConn) {
	error := conn.Wait(); if error != nil {
		s.Remove()
	}
	return
}

//checks how many sessions a user has open (session locker)
func (s *Session) Open(conn *ssh.ServerConn) int {
	Ammount := 0
	// loops throughout the session list checking the ammount of sessions open by that person
	for _, s := range Sessions {
		if s.User.Username == conn.User() {
			Ammount++
		}
	}
	return Ammount
}

func Online() int {
	return len(Sessions)
}


//direct message all sessions open by a user
func DirectMessage(user string, payload string) error {
	for _, s := range Sessions {
		if s.User.Username == user {
			_, err := s.Channel.Write([]byte(payload)); if err != nil {
				return err
			}
		}
	}
	return nil
}

//sends a message to all users online
func Broadcast(payload []byte) error {
	for _, s := range Sessions {
		_, err := s.Channel.Write(payload); if err != nil {
			return err
		}
	}
	return nil
}
