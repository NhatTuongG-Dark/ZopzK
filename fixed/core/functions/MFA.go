package functions 

import (

)

func MFA(Current string) bool {
	if len(Current) > 1 {
		return true
	} else {
		return false
	}
}