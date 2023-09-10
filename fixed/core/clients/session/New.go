package Sessions 

import (

)

func New(User string) *Session {
	for _, s := range Sessions {
		if s.User.Username == User {
			return s
		}
	}
	return nil
}