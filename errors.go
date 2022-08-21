package apperror

import (
	"fmt"
)

type WtcError int64

const (
	// WTC000001 System Error (Something went wrong...)
	WTC000001 WtcError = iota
	// WTC000002 Bad Request (something wrong with user data)
	WTC000002
	// WTC000003 Not Found (Not found)
	WTC000003
	// WTC000004 User with such email is already registered
	WTC000004
	// WTC000005 Invalid email or password
	WTC000005
	// WTC000006 Passwords do not match
	WTC000006
)

func (e WtcError) String() string {
	return "WTC-" + fmt.Sprintf("%06d", e)
}
