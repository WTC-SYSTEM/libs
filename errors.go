package apperror

type ErrorsEnum = string

const (
	// WTC000001 System Error (Something went wrong...)
	WTC000001 ErrorsEnum = "WTC-000001"
	// WTC000002 Bad Request (some thing wrong with user data)
	WTC000002 ErrorsEnum = "WTC-000002"
	// WTC000003 Not Found (Not found)
	WTC000003 ErrorsEnum = "WTC-000003"
	// WTC000004 User with such email is already registered
	WTC000004 ErrorsEnum = "WTC-000004"
	// WTC000005 Invalid email or password
	WTC000005 ErrorsEnum = "WTC-000005"
	WTC000006 ErrorsEnum = "WTC-000006"
)
